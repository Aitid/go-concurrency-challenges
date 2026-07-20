### Problem Description

Implement a concurrent web crawler that starts from a root URL and explores all reachable URLs up to a specified depth.
You are given:
```go
// Fetcher retrieves URLs from a page.
type Fetcher interface {
    Fetch(ctx context.Context, url string) ([]string, error)
}
```
- Fetch returns all URLs found on the page or an error.
    
Your crawler should:
1. Visit each URL at most once.
2. Print each URL as it is visited along with its depth.
3. Handle HTTP errors gracefully without stopping the crawl.
4. Optionally control concurrency (worker pool) for large URL sets.
### Requirements
- Depth control: do not visit URLs beyond the max depth.
- Duplicate suppression: ensure each URL is fetched only once.
- Concurrent fetching: URLs at the same depth can be fetched in parallel.
- Error handling: failed fetches are logged or skipped.
- Output: print URL and `depth` in real time.
### Hints
- Maintain a visited map with sync.Mutex or sync.Map.
- Use a queue or channel of (URL, depth) pairs.
- Launch goroutines for each URL to fetch concurrently.
- Use select/context for cancellation or timeout (optional).
- Recursive approaches work, but ensure thread-safe access to shared state.
- Consider a worker pool if controlling maximum concurrency.
### Optional Enhancements (Senior Level)
- Limit maximum concurrent fetches to avoid overloading servers.
- Support context-based cancellation or global timeout.
- Track metrics: number of URLs visited, errors, time per depth level.
- Log crawl path or statistics per depth.
### Skills Tested
- Concurrency with goroutines, channels, and mutexes.
- Safe shared state management (duplicate suppression).
- Error handling in networked operations.
- Traversal algorithms: breadth-first or depth-first.
- Optional: worker pool, metrics, and context cancellation.
### Execution Flow (Illustration)

```
        Root URL (depth 0)
               |
       Fetch URLs concurrently
               |
   +-----------+-----------+
   |           |           |
URL1 (1)     URL2 (1)    URL3 (1)
   |           |           |
 Fetch URLs concurrently for each, up to max depth
   |           |           |
  ...
```
- Each URL is fetched only once.
- Goroutines fetch URLs in parallel, results printed as they arrive.
- Depth ensures crawler does not go beyond the user-specified limit.
