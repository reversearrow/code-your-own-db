package db0204

type DB struct {
	KV KV
}

func (db *DB) Open() error  { return db.KV.Open() }
func (db *DB) Close() error { return db.KV.Close() }

func (db *DB) Select(schema *Schema, row Row) (ok bool, err error) {
	k := row.EncodeKey(schema)
	val, ok, err := db.KV.Get(k)
	if !ok || err != nil {
		return ok, err
	}
	if err := row.DecodeVal(schema, val); err != nil {
		return false, err
	}

	return true, nil
}

func (db *DB) Insert(schema *Schema, row Row) (updated bool, err error) {
	k := row.EncodeKey(schema)
	v := row.EncodeVal(schema)
	return db.KV.SetEx(k, v, ModeInsert)
}

func (db *DB) Upsert(schema *Schema, row Row) (updated bool, err error) {
	k := row.EncodeKey(schema)
	v := row.EncodeVal(schema)
	return db.KV.SetEx(k, v, ModeUpsert)
}

func (db *DB) Update(schema *Schema, row Row) (updated bool, err error) {
	k := row.EncodeKey(schema)
	v := row.EncodeVal(schema)
	return db.KV.SetEx(k, v, ModeUpdate)
}

func (db *DB) Delete(schema *Schema, row Row) (deleted bool, err error) {
	k := row.EncodeKey(schema)
	return db.KV.Del(k)
}

// QzBQWVJJOUhU https://trialofcode.org/
