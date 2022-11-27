package memcache

import "sync"

type SessionMemCache struct {
	cache sync.Map
}

func NewSessionMemCache() *SessionMemCache {
	return &SessionMemCache{
		cache: sync.Map{},
	}
}


// Create new session
func (s *SessionMemCache) Create(uuidA string){}

// GetByID session.
func (s *SessionMemCache) GetByID(uuidS string){}

// GetAll account sessions using provided account id.
func (s *SessionMemCache) GetAll(uuidA string){}

// Finish session by id excluding current session with id.
func (s *SessionMemCache) Finish(uuidS string){}

// FinishAll account sessions excluding current session with id.
func (s *SessionMemCache) FinishAll(uuidA string){}