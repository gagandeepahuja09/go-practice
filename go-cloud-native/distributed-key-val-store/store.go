package main

import (
	"errors"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrNoSuchKey = errors.New("no such key")

func Put(key, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	val, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return "", ErrNoSuchKey
	}
	return val, nil
}

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}
