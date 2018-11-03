package peer

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"p2p/message"
	"time"
)

// 连接
// 发消息
// 收消息

type Peer struct {
	conn net.Conn
}

func (p *Peer) readMessage() (*message.Message, error) {
	buf := make([]byte, message.HeaderSize)
	_, err := io.ReadFull(p.conn, buf)
	if err != nil {
		return nil, err
	}

	header := message.Header{}
	err = header.Deserialize(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	msg := message.Message{
		Header:header,
	}
	msg.Payload = make([]byte, header.Length)
	_, err = io.ReadFull(p.conn, msg.Payload)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (p *Peer) SendMessage(msg *message.Message) error {
	buf := new(bytes.Buffer)
	if err := msg.Serialize(buf); err != nil {
		return err
	}

	_, err := p.conn.Write(buf.Bytes())
	return err
}

func NewPeer(conn net.Conn) *Peer {
	p := Peer{
		conn: conn,
	}

	go func() {
		for {
			msg, err := p.readMessage()
			fmt.Println(msg, err)

			err = p.SendMessage(msg)
			if err != nil {
				fmt.Println(err)
			}

			time.Sleep(time.Second)
		}
	}()

	return &p
}