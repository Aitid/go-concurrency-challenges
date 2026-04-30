//go:build solution
// +build solution

package parallelgetter

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type GetFunc func(ctx context.Context, address, key string) (string, error)

func Get(ctx context.Context, address, key string) (string, error) {
	delay := time.Duration(rand.Intn(5)+1) * time.Second
	select {
	case <-time.After(delay):
	case <-ctx.Done():
		return "", ctx.Err()
	}

	if rand.Intn(3) == 0 {
		return "", fmt.Errorf("%s isn't available", address)
	}

	return fmt.Sprintf("response from %s for %s after %s", address, key, delay), nil
}

func DistributiveGet(ctx context.Context, getter GetFunc, addresses []string, key string) (string, error) {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan string)
	done := make(chan struct{})
	var wg sync.WaitGroup

	for _, addr := range addresses {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			response, err := getter(ctxCancel, addr, key)
			if err != nil {
				return
			}
			select {
			case <-ctxCancel.Done():
				return
			case ch <- response:
			}
		}(addr)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-ctxCancel.Done():
			return "", ctxCancel.Err()
		case <-done:
			return "", fmt.Errorf("all addresses are unavailable")
		case res, ok := <-ch:
			if ok {
				return res, nil
			}
			return "", fmt.Errorf("something wrong")
		}
	}
}
