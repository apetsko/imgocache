package imgocache

import (
	"sync"
	"time"
)

// Imgocache struct contains simple map[string]interface{} as db field and sync.Mutex for synchronization
type Imgocache struct {
	sync.RWMutex
	db map[string]item
}

// item holds the value and the expiration time of the ttl
type item struct {
	value      interface{}
	expiration time.Time
}

// New creates new Imgocache
func New() *Imgocache {
	return &Imgocache{
		db: make(map[string]item),
	}
}

// Set method sets the value v for the key k in the cache
func (s *Imgocache) Set(k string, v interface{}, ttl time.Duration) {
	s.Lock()
	defer s.Unlock()
	s.db[k] = item{
		value:      v,
		expiration: time.Now().Add(ttl),
	}
}

// Get method retrieves the value v for the key k from the cache
// It returns v and a boolean indicating whether the value exists or not
func (s *Imgocache) Get(k string) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()
	item, exist := s.db[k]
	if !exist {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		delete(s.db, k)
		return nil, false
	}
	return item.value, true
}

// Delete method deletes the value from the cache for the given key
func (s *Imgocache) Delete(k string) {
	s.Lock()
	defer s.Unlock()
	delete(s.db, k)
}
