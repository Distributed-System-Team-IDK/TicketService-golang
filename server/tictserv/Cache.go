package tictserv

import (
	"sync"
	"time"
)

const (
	cacheSize = 10
)

type entry struct {
	value any
	id    string
	time  time.Time
	valid bool

	mu sync.Mutex
}

type Cache struct {
	entries [cacheSize]*entry
}

func NewCache() *Cache {
	entries := [cacheSize]*entry{}
	for i := range entries {
		entries[i] = &entry{
			valid: false,
		}
	}

	return &Cache{
		entries: entries,
	}
}

func (c *Cache) Cache(id string, value any) {
	firstInvalid := -1
	lastOneCalled := -1
	oldestTime := time.Time{}
	timeNotSet := true
	for i := 0; i < cacheSize; i++ {
		entry := c.entries[i]
		if !entry.valid {
			if firstInvalid == -1 {
				firstInvalid = i
			}
			continue
		}
		if timeNotSet || entry.time.Before(oldestTime) {
			lastOneCalled = i
			oldestTime = entry.time
			timeNotSet = false
		}

		if entry.id == id {
			entry.mu.Lock()
			entry.value = value
			entry.time = time.Now()
			entry.valid = true
			entry.mu.Unlock()
			return
		}
	}
	var entry *entry
	if firstInvalid != -1 {
		entry = c.entries[firstInvalid]
	} else {
		entry = c.entries[lastOneCalled]
	}
	entry.mu.Lock()
	entry.value = value
	entry.time = time.Now()
	entry.valid = true
	entry.mu.Unlock()
}

func (c *Cache) Load(id string) (any, bool) {
	for i := 0; i < cacheSize; i++ {
		entry := c.entries[i]
		if !entry.valid {
			continue
		}
		if entry.id == id {
			entry.mu.Lock()
			entry.time = time.Now()
			entry.mu.Unlock()
			return entry.value, true
		}
	}
	return nil, false
}
