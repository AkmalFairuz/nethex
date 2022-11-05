package server

import (
	"github.com/akmalfairuz/nethex/packet"
	"net"
	"sync"
)

type Server struct {
	listener net.Listener

	clients   map[net.Addr]*Client
	clientsMu sync.Mutex

	packetPool *packet.Pool
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:   listener,
		clients:    map[net.Addr]*Client{},
		packetPool: packet.NewPool(),
	}, nil
}

func (s *Server) Accept() (*Client, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}
	s.clientsMu.Lock()
	c := newClient(s, conn.(*net.TCPConn))
	s.clients[conn.RemoteAddr()] = c
	s.clientsMu.Unlock()
	return c, nil
}

func (s *Server) PacketPool() *packet.Pool {
	return s.packetPool
}

func (s *Server) RemoveClient(c *Client) {
	s.clientsMu.Lock()
	delete(s.clients, c.conn.RemoteAddr())
	s.clientsMu.Unlock()
}
