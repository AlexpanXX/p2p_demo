package message

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
)

const (
	CommandSize  = 12
	ChecksumSize = 4
	HeaderSize   = 24
)

// P2P 消息
type Message struct {
	// Header
	Header
	// Payload
	Payload []byte
}

func (m Message) String() string {
	return string(m.Payload)
}

func (m *Message) Deserialize(r io.Reader) error {
	err := m.Header.Deserialize(r)
	if err != nil {
		return err
	}

	m.Payload, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) Serialize(w io.Writer) error {
	err := m.Header.Serialize(w)
	if err != nil {
		return err
	}

	_, err = w.Write(m.Payload)
	if err != nil {
		return err
	}

	return nil
}

func NewMessage(magic uint32, cmd string,
	payload []byte) *Message {
	command := [CommandSize]byte{}
	copy(command[:], cmd)
	h := Header{
		Magic: magic,
		Command: command,
	}
	m := Message{
		Header : h,
		Payload: payload,
	}

	m.Header.Length = uint32(len(payload))

	hash := sha256.Sum256(m.Payload)

	copy(m.Header.Checksum[:], hash[:4])

	return &m
}
