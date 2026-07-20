package main

import (
	"container/list"
	"context"
	"time"
)

func ExternalAPIRequest(ctx context.Context, key string) (string, error) {
	time.Sleep(100 * time.Millisecond)
	return "value for " + key, nil
}

type Item struct {
	v       string
	expired time.Time
}

type APICache struct {
	m         map[string]Item
	data      list.List
	sizeLimit int
	ttl       time.Duration
}

func NewAPICache(ttl time.Duration, sizeLimit int) *APICache {
	return &APICache{ttl: ttl, sizeLimit: sizeLimit}
}

func (c *APICache) add() {}
func (c *APICache) get() {}
func (c *APICache) Get(ctx context.Context, key string) (string, error) {
	if v, ok := c.m[key]; ok {
		return v.v, nil
	}
	res, err := ExternalAPIRequest(ctx, key)
	if err != nil {
	}
	c.add()

	return res, nil
}

// LRU?
func main() {
}
