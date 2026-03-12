package db0202

import (
	"bytes"
	"slices"
)

type Schema struct {
	Table string
	Cols  []Column
	PKey  []int // indexes of primary key columns
}

type Column struct {
	Name string
	Type CellType
}

type Row []Cell

func (schema *Schema) NewRow() Row {
	return make(Row, len(schema.Cols))
}

func (row Row) EncodeKey(schema *Schema) (key []byte) {
	key = append([]byte(schema.Table), 0x00)
	for _, v := range schema.PKey {
		key = row[v].Encode(key)
	}
	return key
}

func (row Row) EncodeVal(schema *Schema) (val []byte) {
	for i := range schema.Cols {
		if slices.Contains(schema.PKey, i) {
			continue
		}
		encodedVal := row[i].Encode(val)
		val = append(val, encodedVal...)
	}

	return val
}

func (row Row) DecodeKey(schema *Schema, key []byte) (err error) {
	zeroSeperator := bytes.IndexByte(key, 0x00)
	rest := key[zeroSeperator+1:]
	for _, pKey := range schema.PKey {
		cellType := schema.Cols[pKey].Type
		c := Cell{
			Type: cellType,
		}
		rest, err = c.Decode(rest)
		if err != nil {
			return err
		}
		row[pKey] = c
	}
	return nil
}

func (row Row) DecodeVal(schema *Schema, val []byte) (err error) {
	for i := range schema.PKey {
		if slices.Contains(schema.PKey, i) {
			continue
		}
		cellType := schema.Cols[i].Type
		c := Cell{
			Type: cellType,
		}
		val, err = c.Decode(val)
		if err != nil {
			return err
		}
		row[i] = c
	}

	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
