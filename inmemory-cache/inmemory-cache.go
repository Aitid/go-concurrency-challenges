package main

import (
	"container/list"
	"sync"
	"time"
)

type Entry struct {
	key       string
	value     any
	expiredAt time.Time
}

type APICacher struct {
	mu       sync.Mutex
	data     *list.List
	mapping  map[string]*list.Element
	capacity int
}

func NewAPICacher(capacity int) *APICacher {
	return &APICacher{
		data:     list.New(),
		mapping:  make(map[string]*list.Element),
		capacity: capacity,
	}
}

func (c *APICacher) Get(key string) (Entry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.mapping[key]
	entry := el.Value.(Entry)

	if !ok || time.Now().After(entry.expiredAt) {
		return Entry{}, false
	}

	c.data.MoveToFront(el)

	return entry, ok
}

func (c *APICacher) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := Entry{
		key:       key,
		value:     value,
		expiredAt: time.Now().Add(ttl),
	}

	if c.data.Len() == c.capacity {
		element := c.data.Back()
		if element != nil {
			ent := element.Value.(Entry)
			delete(c.mapping, ent.key)
			c.data.Remove(element)
		}
	}
	newEl := c.data.PushFront(entry)
	c.mapping[key] = newEl
}

func main() {
}
