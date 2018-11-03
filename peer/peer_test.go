package peer

import (
	"net"
	"p2p/message"
	"testing"
	"time"
)

func TestNewPeer(t *testing.T) {
	listener, err := net.Listen(
		"tcp", ":6000")
	if err != nil {
		t.Error(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Error(err)
				continue
			}

			NewPeer(conn)
		}
	}()
}

func TestPeer_SendMessage(t *testing.T) {

	conn, err := net.Dial(
		"tcp", "localhost:6000")
	if err != nil {
		t.Error(err)
	}

	peer := NewPeer(conn)

	payload := "im the message payload"

	for {
		err = peer.SendMessage(message.NewMessage(
			12345, "command", []byte(payload)))
		if err != nil {
			t.Error(err)
		}

		time.Sleep(time.Second)
	}
}
