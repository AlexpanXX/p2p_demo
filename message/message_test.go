package message

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestMessage_Serialize(t *testing.T) {
	command := [CommandSize]byte{}
	copy(command[:], "getblocks")
	payload := "i am the message payload"

	h := Header{
		Magic: 12345,
		Command: command,
	}
	m := Message{
		Header : h,
		Payload: []byte(payload),
	}

	m.Header.Length = uint32(len(payload))

	hash := sha256.Sum256(m.Payload)

	copy(m.Header.Checksum[:], hash[:4])

	buf := new(bytes.Buffer)
	err := m.Serialize(buf)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(m)
	fmt.Println(buf.Bytes())


	msg := Message{}
	err = msg.Deserialize(buf)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(msg)
}
