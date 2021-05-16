package store

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type RecordStore interface {
	Set(key, value string) (string, bool)
	Get(key string) (string, error)
	Sync()
	Map() map[string]string
	Init(map[string]string)
}

type InMemoryRecordStorage struct {
	store map[string]string
	m     sync.Mutex
}

func NewInMemoryRecordStorage() RecordStore {
	return &InMemoryRecordStorage{
		store: make(map[string]string),
		m:     sync.Mutex{},
	}
}

func (s *InMemoryRecordStorage) Set(key, value string) (string, bool) {
	s.m.Lock()
	defer s.m.Unlock()
	// check if key already present
	val, ok := s.store[key]
	if ok {
		return val, true
	}

	// if key is not present
	s.store[key] = value

	return value, false
}

func (s *InMemoryRecordStorage) Get(value string) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()
	for key, val := range s.store {
		if val == value {
			return key, nil
		}
	}
	return "", fmt.Errorf("unable to get: original link for %s not found", value)
}

func (s *InMemoryRecordStorage) Map() map[string]string {
	return s.store
}

func (s *InMemoryRecordStorage) Init(store map[string]string) {
	s.store = store
}

func (s *InMemoryRecordStorage) Sync() {
}

type FileRecordStorage struct {
	filename     string
	cache        RecordStore
	modified     bool
	syncDuration time.Duration
	m            sync.Mutex
}

func NewFileRecordStorage(storePath string, syncDuration time.Duration) RecordStore {
	return &FileRecordStorage{
		cache:        NewInMemoryRecordStorage(),
		filename:     storePath,
		modified:     false,
		syncDuration: syncDuration,
	}
}

func (s *FileRecordStorage) Set(key, value string) (string, bool) {
	s.m.Lock()
	defer s.m.Unlock()
	val, ok := s.cache.Set(key, value)
	if ok {
		return val, true
	}
	s.modified = true
	return val, false
}

func (s *FileRecordStorage) Get(value string) (string, error) {
	return s.cache.Get(value)
}

func (s *FileRecordStorage) Map() map[string]string {
	return s.cache.Map()
}

func (s *FileRecordStorage) Init(st map[string]string) {
	s.cache.Init(st)
}

func (s *FileRecordStorage) Sync() {
	// check file exist
	r, err := os.Open(s.filename)
	if err == nil {
		d := gob.NewDecoder(r)
		var cache map[string]string
		d.Decode(&cache)
		s.Init(cache)
	}
	r.Close()

	go func() {
		ticker := time.NewTicker(s.syncDuration * time.Second)
		for t := range ticker.C {
			fmt.Printf("Sync: cache synced & backedup at %s\n", t.String())
			if !s.modified {
				continue
			}
			s.m.Lock()

			b := new(bytes.Buffer)
			e := gob.NewEncoder(b)

			err := e.Encode(s.cache.Map())
			if err != nil {
				fmt.Println("Unable to encode")
				return
			}
			err = ioutil.WriteFile(s.filename, b.Bytes(), 0644)
			if err != nil {
				fmt.Println("unable to write to file")
				return
			}

			s.m.Unlock()
		}
	}()
}
