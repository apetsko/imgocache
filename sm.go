package imgocache

import (
	"sync"
	"time"
)

// SM type is the same as sync.Map type, which is safe for concurrent read-write operations
type SM struct {
	sync.Map
}

// NewSM creates a new instance of SM
func NewSM() *SM {
	return &SM{}
}

// Set sets the value v for the key k in the cache with TTL consideration
func (s *SM) Set(k string, v interface{}, ttl time.Duration) {
	s.Store(k, item{
		value:      v,
		expiration: time.Now().Add(ttl),
	})
}

// Get retrieves the value v for the key k from the cache
// It returns v and a boolean indicating whether the value exists or not
func (s *SM) Get(k string) (v interface{}, exist bool) {
	if val, ok := s.Load(k); ok && (val != nil) {
		item := val.(item)
		if time.Now().Before(item.expiration) {
			return item.value, true
		}
	}
	return nil, false
}

// Delete deletes the value from the cache for the given key
func (s *SM) Delete(k string) {
	s.Store(k, nil)
}
