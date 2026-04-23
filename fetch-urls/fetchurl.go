package main

import (
	"context"
	"fmt"
	"time"
)

func Get(ctx context.Context, url string) (string, error) {
	return "", nil
}

func FetchURLs(ctx context.Context, urls []string, concurrency int) []string {
	return []string{}
}

func main() {
	urls := []string{"url1", "url2", "url3", "url4", "url5", "url6"}
	start := time.Now()

	ctx := context.Background()
	result := FetchURLs(ctx, urls, 3)

	elapsed := time.Since(start)
	fmt.Println(result, elapsed)
}
