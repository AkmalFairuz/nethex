package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/akmalfairuz/nethex/packet"
	"net"
	"sync"
)

type Client struct {
	s    *Server
	conn *net.TCPConn

	sendMu     sync.Mutex
	sendPacket []packet.Packet

	closed bool

	flushMu sync.Mutex
}

func (c *Client) Address() string {
	return c.conn.RemoteAddr().String()
}

func (c *Client) ReadPacket() (packet.Packet, error) {
	var readLen uint32
	if err := binary.Read(c.conn, binary.BigEndian, &readLen); err != nil {
		return nil, err
	}

	packetBuf := make([]byte, readLen)
	if _, err := c.conn.Read(packetBuf); err != nil {
		return nil, err
	}
	pkt, ok := c.s.packetPool.Get(packetBuf[0])
	if !ok {
		return nil, fmt.Errorf("there are no packet with id %d", packetBuf[0])
	}
	pkt.Unmarshal(packet.NewReader(bytes.NewBuffer(packetBuf[1:])))
	return pkt, nil
}

func (c *Client) WritePacket(pk packet.Packet) {
	c.sendMu.Lock()
	c.sendPacket = append(c.sendPacket, pk)
	c.sendMu.Unlock()
}

func (c *Client) Flush() {
	var pktToSend []packet.Packet
	c.sendMu.Lock()
	pktToSend = c.sendPacket
	c.sendPacket = make([]packet.Packet, 0)
	c.sendMu.Unlock()

	c.flushMu.Lock()
	for _, pkt := range pktToSend {
		buf := bytes.NewBuffer(nil)
		buf.WriteByte(pkt.Id())
		pkt.Marshal(packet.NewWriter(buf))
		if err := binary.Write(c.conn, binary.BigEndian, uint32(buf.Len())); err != nil {
			panic(err)
		}
		if _, err := c.conn.Write(buf.Bytes()); err != nil {
			panic(err)
		}
	}
	c.flushMu.Unlock()
}

func (c *Client) Close() {
	c.flushMu.Lock()
	c.sendMu.Lock()

	if err := c.conn.Close(); err != nil {
		panic(err)
	}
	c.s.RemoveClient(c)
	c.closed = true

	c.flushMu.Unlock()
	c.sendMu.Unlock()
}

func (c *Client) Closed() bool {
	return c.closed
}

func (c *Client) WriteRaw(b []byte) {
	if _, err := c.conn.Write(b); err != nil {
		panic(err)
	}
}

func newClient(s *Server, conn *net.TCPConn) *Client {
	return &Client{s: s, conn: conn, closed: false}
}
