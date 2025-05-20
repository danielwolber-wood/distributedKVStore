package main

import (
	"fmt"
	"sync"
)

type KVStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]string),
	}
}

func (s *KVStore) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key does not exist")
	}
	return value, nil
}

func (s *KVStore) Set(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

func (s *KVStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}
