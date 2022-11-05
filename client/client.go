package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/akmalfairuz/nethex/packet"
	"net"
	"sync"
)

type Client struct {
	conn *net.TCPConn

	sendMu     sync.Mutex
	sendPacket []packet.Packet

	flushMu sync.Mutex

	closed bool

	packetPool *packet.Pool
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
	pkt, ok := c.packetPool.Get(packetBuf[0])
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

func (c *Client) Flush() error {
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
			return err
		}
		if _, err := c.conn.Write(buf.Bytes()); err != nil {
			return err
		}
	}
	c.flushMu.Unlock()
	return nil
}

func (c *Client) Close() {
	c.flushMu.Lock()
	c.sendMu.Lock()

	if err := c.conn.Close(); err != nil {
		panic(err)
	}
	c.closed = true

	c.flushMu.Unlock()
	c.sendMu.Unlock()
}

func (c *Client) PacketPool() *packet.Pool {
	return c.packetPool
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn.(*net.TCPConn), packetPool: packet.NewPool(), closed: false}, nil
}
