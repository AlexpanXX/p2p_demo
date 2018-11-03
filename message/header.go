package message

import (
	"encoding/binary"
	"io"
)

// 消息头 24byte
type Header struct {
	// Magic 4 byte
	Magic uint32
	// Command 12 byte
	Command [CommandSize]byte
	// Length 4 byte
	Length uint32
	// Checksum 4 byte
	Checksum [ChecksumSize]byte
}

func (h *Header) Deserialize(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &h.Magic)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, h.Command[:])
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &h.Length)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, h.Checksum[:])
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) Serialize(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, h.Magic)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, h.Command)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, h.Length)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, h.Checksum)
	if err != nil {
		return err
	}

	return nil
}
