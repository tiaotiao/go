package sync

import (
	"sync"
)

type Map struct {
	l sync.RWMutex
	m map[interface{}]interface{}
}

func NewMap() *Map {
	m := new(Map)
	m.m = make(map[interface{}]interface{})
	return m
}

func (m *Map) Get(key interface{}) (interface{}, bool) {
	m.l.RLock()
	v, ok := m.m[key]
	m.l.RUnlock()
	return v, ok
}

func (m *Map) Set(key interface{}, val interface{}) {
	m.l.Lock()
	m.m[key] = val
	m.l.Unlock()
}

func (m *Map) Del(key interface{}) {
	m.l.Lock()
	delete(m.m, key)
	m.l.Unlock()
}

func (m *Map) Len() int {
	m.l.RLock()
	n := len(m.m)
	m.l.RUnlock()
	return n
}

func (m *Map) For(f func(k interface{}, v interface{}) bool) {
	m.l.RLock()
	var ok bool
	for k, v := range m.m {
		ok = f(k, v)
		if !ok {
			break
		}
	}
	m.l.RUnlock()
}

func (m *Map) Keys() []interface{} {
	m.l.RLock()
	keys := make([]interface{}, 0, len(m.m))
	for k, _ := range m.m {
		keys = append(keys, k)
	}
	m.l.RUnlock()
	return keys
}

func (m *Map) Vals() []interface{} {
	m.l.RLock()
	vals := make([]interface{}, 0, len(m.m))
	for _, v := range m.m {
		vals = append(vals, v)
	}
	m.l.RUnlock()
	return vals
}

func (m *Map) KeyVals() ([]interface{}, []interface{}) {
	m.l.RLock()
	keys := make([]interface{}, 0, len(m.m))
	vals := make([]interface{}, 0, len(m.m))
	for k, v := range m.m {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	m.l.RUnlock()
	return keys, vals
}
