package imgocache

import (
	"sync"
)

// Imgocache struct contains simple map[string]interface{} as db field and sync.Mutex for synchronization
type Imgocache struct {
	sync.RWMutex
	db map[string]interface{}
}

// New creates new Imgocache
func New() *Imgocache {
	return &Imgocache{
		db: make(map[string]interface{}),
	}
}

// Set method sets the value v for the key k in the cache
func (s *Imgocache) Set(k string, v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.db[k] = v
}

// Get method retrieves the value v for the key k from the cache
// It returns v and a boolean indicating whether the value exists or not
func (s *Imgocache) Get(k string) (v interface{}, exist bool) {
	s.RLock()
	defer s.RUnlock()
	v, exist = s.db[k]
	return
}

// Delete method deletes the value from the cache for the given key
func (s *Imgocache) Delete(k string) {
	s.Lock()
	defer s.Unlock()
	delete(s.db, k)
}
