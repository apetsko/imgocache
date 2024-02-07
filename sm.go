package imgocache

import "sync"

// SM type is the same sync.Map type, which is safe to conrurrent read-write operations
type SM sync.Map

// NewSM creates new SM
func NewSM() *SM {
	return &SM{}
}

// Set method sets the value v for the key k in the cache
func (s *SM) Set(k string, v interface{}) {
	s.Set(k, v)
}

// Get method retrieves the value v for the key k from the cache
// It returns v and a boolean indicating whether the value exists or not
func (s *SM) Get(k string) (v interface{}, exist bool) {
	v, exist = s.Get(k)
	return
}

// Delete method deletes the value from the cache for the given key
func (s *SM) Delete(k string) {
	s.Delete(k)
}
