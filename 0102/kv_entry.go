package db0102

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	data := make([]byte, 4+4+len(ent.key)+len(ent.val))
	binary.LittleEndian.PutUint32(data[0:4], uint32(len(ent.key)))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.val)))
	copy(data[8:], ent.key)
	copy(data[8+len(ent.key):], ent.val)
	return data
}

func (ent *Entry) Decode(r io.Reader) error {
	data := make([]byte, 4096)
	_, err := r.Read(data)
	if err != nil {
		return err
	}
	keyLen := binary.LittleEndian.Uint32(data[0:4])
	valLen := binary.LittleEndian.Uint32(data[4:8])
	ent.key = data[8 : 8+keyLen]
	ent.val = data[8+keyLen : 8+keyLen+valLen]

	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
