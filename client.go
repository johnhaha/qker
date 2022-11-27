package qker

import (
	"crypto/tls"
	"encoding/json"

	"github.com/lucas-clemente/quic-go"
)

type Client struct {
	Addr    string
	Handler Handler
	quic.Connection
}

func NewClient(addr string) *Client {
	return &Client{Addr: addr}
}

func (s *Client) SetHandler(handler Handler) {
	s.Handler = handler
}

func (s *Client) Dial() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"qker-cnn"},
	}
	conn, err := quic.DialAddr(s.Addr, tlsConf, &quic.Config{
		EnableDatagrams: true,
	})
	if err != nil {
		return err
	}
	s.Connection = conn
	if s.Handler != nil {
		go handleMsg(conn, s.Handler)
	}
	return nil
}

func (s *Client) Send(data string) error {
	err := s.SendMessage([]byte(data))
	return err
}

func (s *Client) SendJson(data any) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.SendMessage(res)
}
