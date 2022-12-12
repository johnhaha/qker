package qker

import (
	"crypto/tls"
	"encoding/json"

	"github.com/lucas-clemente/quic-go"
)

type Client struct {
	Addr string
	quic.Connection
}

func NewClient(addr string) *Client {
	return &Client{Addr: addr}
}

func (s *Client) InitConn() error {
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
	return nil
}

func (s *Client) Fetch(data []byte) ([]byte, error) {
	err := s.SendMessage(data)
	if err != nil {
		return nil, err
	}
	res, err := s.ReceiveMessage()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Client) Dial(data []byte, handler Handler) error {
	err := s.SendMessage(data)
	if err != nil {
		return err
	}
	return handleMsg(s.Connection, handler)
}

func (s *Client) Send(data []byte) error {
	err := s.SendMessage(data)
	return err
}

func (s *Client) SendJson(data any) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.SendMessage(res)
}
