package main

import (
	"context"
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]string, error)
}

/*

 scheduler([]result) --> []jobs
 // scheduler --cyclic dependency--> worker
 for range N{
 	go worker(job) --> []Result
 }

 Page {
	url string
	depth int
 }
*/

type Page struct {
	url   string
	depth int
}

type APICrawler struct {
	numWorkers int
	maxDepth   int
	client     Fetcher
	jobCh      chan Page
	resCh      chan []Page
	wg         sync.WaitGroup
}

func NewAPICrawler(numWorkers int, maxDepth int, client Fetcher) *APICrawler {
	return &APICrawler{
		numWorkers: numWorkers,
		maxDepth:   maxDepth,
		jobCh:      make(chan Page, 10),
		resCh:      make(chan []Page, 10),
	}
}

func (c *APICrawler) Crawle(ctx context.Context, url string) {
	// scheduler
	go func() {
		rootPage := Page{url: url, depth: 0}
		visited := make(map[string]bool)
		visited[rootPage.url] = true
		queue := []Page{rootPage}
		count := 1

		for {
			var rCh chan []Page
			var nextPage Page

			if len(queue) > 0 {
				rCh = c.resCh
				nextPage = queue[0]
			}
			select {
			case pages := <-rCh:
				count--
				for _, res := range pages {
					if !visited[res.url] && c.maxDepth < res.depth {
						count++
						queue = append(queue, Page{url: res.url, depth: res.depth})
					}
				}
			case c.jobCh <- nextPage:
				queue = queue[1:]
			case <-ctx.Done():
				return
			}
		}
	}()

	for range c.numWorkers {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			for {
				select {
				case job, ok := <-c.jobCh:
					if !ok {
						return
					}

					pages, err := c.client.Fetch(ctx, job.url)
					if err != nil {
						fmt.Println(job.url, " isn't available")
					} else {
						fmt.Println(job.url, " is handled")
						resPages := []Page{}
						for _, page := range pages {
							resPages = append(resPages, Page{url: page, depth: job.depth + 1})
						}

						select {
						case c.resCh <- resPages:
						case <-ctx.Done():
							return
						}
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	c.wg.Wait()
}

func main() {
}
