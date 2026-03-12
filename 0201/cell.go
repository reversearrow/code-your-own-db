package db0201

import (
	"encoding/binary"
	"errors"
)

type CellType uint8

const (
	TypeI64 CellType = 1
	TypeStr CellType = 2
)

type Cell struct {
	Type CellType
	I64  int64
	Str  []byte
}

func (cell *Cell) Encode(toAppend []byte) []byte {
	switch cell.Type {
	case TypeStr:
		toAppend = binary.LittleEndian.AppendUint32(toAppend, uint32(len(cell.Str)))
		toAppend = append(toAppend, cell.Str...)
	case TypeI64:
		toAppend = binary.LittleEndian.AppendUint64(toAppend, uint64(cell.I64))
	}
	return toAppend
}

func (cell *Cell) Decode(data []byte) (rest []byte, err error) {
	switch cell.Type {
	case TypeStr:
		size := binary.LittleEndian.Uint32(data[0:4])
		cell.Str = data[4 : size+4]
		return data[size+4:], nil
	case TypeI64:
		i64 := binary.LittleEndian.Uint64(data[0:8])
		cell.I64 = int64(i64)
		return data[8:], nil
	}
	return nil, errors.New("unsupported cell type")
}

// QzBQWVJJOUhU https://trialofcode.org/
