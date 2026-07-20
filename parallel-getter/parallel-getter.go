package parallelgetter

import (
	"context"
	"sync"
)

type Getter interface {
	Get(ctx context.Context, address, key string) (string, error)
}

type Result struct {
	response string
	err      error
}

func DistributiveGet(ctx context.Context, getter Getter, addresses []string, key string) (string, []error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobCh := make(chan string, 10)
	resCh := make(chan Result, 10)

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				response, err := getter.Get(ctx, job, "any")
				select {
				case <-ctx.Done():
					return
				case resCh <- Result{response: response, err: err}:
				}
			}
		}()
	}
	go func() {
		defer close(jobCh)

		for _, addr := range addresses {
			select {
			case jobCh <- addr:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resCh)
	}()

	errors := []error{}
	for answer := range resCh {
		if answer.err == nil {
			return answer.response, nil
		}
		errors = append(errors, answer.err)
	}
	return "", errors
}
