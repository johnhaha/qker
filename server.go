package qker

import (
	"context"

	"github.com/lucas-clemente/quic-go"
)

type Server struct {
	Addr    string
	Handler Handler
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) SetHandler(handler Handler) {
	s.Handler = handler
}

func (s *Server) StartServer(ctx context.Context) error {
	listener, err := quic.ListenAddr(s.Addr, generateTLSConfig(), &quic.Config{
		EnableDatagrams: true,
	})
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept(ctx)
		if err != nil {
			return err
		}
		go handleMsg(conn, s.Handler)
	}
}
