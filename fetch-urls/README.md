### **Problem Description**
You are given a list of URLs. Implement a function that performs an HTTP GET request for each URL:
- Limit the number of concurrent requests using a parameter `concurrency`.
- Return the results in the same order as the input list.
- Ignore errors (return empty string or `nil` if a request fails).
### Requirements
```go
func FetchURLs(urls []string, concurrency int) []string
```
1. Preserve order: results must match input URL order.
2. Concurrency limit: at most `concurrency` requests in flight simultaneously.
3. Error handling: ignore HTTP errors; optionally return empty string.
4. Efficiency: use goroutines, channels, or `sync.WaitGroup` without unnecessary blocking.
5. Context support: cancel all requests on timeout via context.Context.
6. Retry mechanism: automatically retry failed requests a configurable number of times.
### Hints
- Consider a worker pool or semaphore pattern to limit concurrency.
- Use indexed results to maintain order of URLs.
- Use http.NewRequestWithContext(ctx, ...) for timeout/cancellation support.
- Combine sync.WaitGroup with channels to collect results efficiently.
### Skills Tested
- Concurrency with goroutines and channels.
- Correct synchronization of results with order preservation.
- Implementing worker pools or semaphore patterns.
- Optional: context cancellation and retry logic.

