package db0103

type KV struct {
	log Log
	mem map[string][]byte
}

func (kv *KV) Open() error {
	if err := kv.log.Open(); err != nil {
		return err
	}
	kv.mem = make(map[string][]byte)
	for {
		e := &Entry{}
		eof, err := kv.log.Read(e)
		if err != nil {
			return err
		}
		if eof {
			break
		}
		if e.deleted {
			delete(kv.mem, string(e.key))
			continue
		}
		kv.mem[string(e.key)] = e.val
	}

	return nil
}

func (kv *KV) Close() error { return kv.log.Close() }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return val, ok, nil
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	kv.mem[string(key)] = val
	if err := kv.log.Write(&Entry{
		key: key,
		val: val,
	}); err != nil {
		return false, err
	}
	return true, nil
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	if _, ok := kv.mem[string(key)]; !ok {
		return false, nil
	}
	delete(kv.mem, string(key))
	if err := kv.log.Write(&Entry{
		key:     key,
		deleted: true,
	}); err != nil {
		return false, err
	}
	return true, nil
}

// QzBQWVJJOUhU https://trialofcode.org/
