package syslog

import (
	"fmt"
	"net"
)

type Server struct {
	l *net.UDPConn
}

func NewServer(address string) (*Server, error) {

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	v := new(Server)

	v.l, err = net.ListenUDP("udp", addr)

	return v, err
}

func (s *Server) Read() (*Message, error) {

	buf := make([]byte, 2048)

	n, err := s.l.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	buf = buf[:n]

	return ParseMessage(buf)
}

func (s *Server) Close() {
	s.l.Close()
}
