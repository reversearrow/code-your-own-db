package db0101

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	v, ok := kv.mem[string(key)]
	return v, ok, nil
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	kv.mem[string(key)] = val
	return true, nil
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	_, ok := kv.mem[string(key)]
	if ok {
		delete(kv.mem, string(key))
	}
	return ok, nil
}

// QzBQWVJJOUhU https://trialofcode.org/
