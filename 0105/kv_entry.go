package db0105

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
)

type Entry struct {
	key     []byte
	val     []byte
	deleted bool
}

func (ent *Entry) Encode() []byte {
	const headerSize = 4 + 4 + 4 + 1
	data := make([]byte, headerSize+len(ent.key)+len(ent.val))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.key)))
	copy(data[headerSize:], ent.key)
	if ent.deleted {
		data[12] = 1
	} else {
		binary.LittleEndian.PutUint32(data[8:12], uint32(len(ent.val)))
		copy(data[headerSize+len(ent.key):], ent.val)
	}
	binary.LittleEndian.PutUint32(data[0:4], crc32.ChecksumIEEE(data[4:]))
	return data
}

var ErrBadSum = errors.New("bad checksum")

func (ent *Entry) Decode(r io.Reader) error {
	var header [13]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return err
	}
	entryChecksum := binary.LittleEndian.Uint32(header[0:4])
	klen := int(binary.LittleEndian.Uint32(header[4:8]))
	vlen := int(binary.LittleEndian.Uint32(header[8:12]))
	deleted := header[12]
	data := make([]byte, klen+vlen)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}
	ent.key = data[:klen]
	if deleted != 0 {
		ent.deleted = true
	} else {
		ent.deleted = false
		ent.val = data[klen:]
	}

	h := crc32.NewIEEE()
	h.Write(header[4:])
	h.Write(data)

	if h.Sum32() != entryChecksum {
		return ErrBadSum
	}

	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
