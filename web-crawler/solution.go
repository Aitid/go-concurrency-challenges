//go:build solution
// +build solution

package webcrawler

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FetchFunc func(ctx context.Context, url string) ([]string, error)

func Fetch(ctx context.Context, url string) ([]string, error) {
	pages := map[string][]string{
		"0":    {"01", "02", "03"},
		"01":   {"010", "011", "012"},
		"02":   {"020", "021", "022"},
		"03":   {"030", "031", "032"},
		"010":  {"0101", "0102", "0103"},
		"0101": {"01"},
	}
	wait := rand.Int63n(5)
	select {
	case <-time.After(time.Duration(wait) * time.Second):
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if v, ok := pages[url]; ok {
		return v, nil
	}
	return []string{}, nil
}

func Crawl(ctx context.Context, wg *sync.WaitGroup, mu *sync.Mutex, store map[string]struct{}, url string, depth int) {
	if depth <= 0 {
		return
	}

	select {
	case <-ctx.Done():
		return
	default:
	}

	mu.Lock()
	if _, ok := store[url]; ok {
		mu.Unlock()
		return
	} else {
		store[url] = struct{}{}
	}
	mu.Unlock()

	fmt.Println(url, depth)

	urls, err := Fetch(ctx, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Crawl(ctx, wg, mu, store, url, depth-1)
		}()
	}
}
