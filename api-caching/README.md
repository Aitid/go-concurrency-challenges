## Problem Description

Implement a cache for results of calls to an external API to:
- Reduce load on the API
- Speed up responses
- Handle concurrent access safely

Example API:
```go
func ExternalAPIRequest(key string) (string, error) {
    time.Sleep(100 * time.Millisecond)
    return "value for " + key, nil
}
```
### Requirements
1. TTL support: Each cached entry expires after a configurable duration.
2. Concurrency safety: Multiple goroutines may read/write concurrently.
3. Stale data fallback: If the API fails, return expired cache data if available.
4. Size limit: Avoid uncontrolled cache growth (e.g., maxEntries or LRU).
5. Pluggable storage: Design interface to allow different backends (in-memory, Redis, etc.).

Implement a cache with the following API:
```go
type Cache interface {
    Get(ctx context.Context, key string) (string, error)
}
```
## Optional Enhancements (Senior Level)
- Thundering herd protection: Only one goroutine calls the API per key at a time.
- Metrics: Cache hits/misses, stale data returns, average response time.
- Advanced caching strategies: cache-aside, write-through, refresh-ahead.

## Hints
- Use sync.RWMutex, sync.Map, or golang.org/x/sync/singleflight for concurrency. 
- Keep TTL timestamps for automatic expiration.
- Consider container/list or LRU map for size-limited cache.
- Think about error handling and fallback to stale data.

## **Skills Tested**
- Concurrency and goroutine safety
- Cache design with TTL and size management
- Error handling and fallback strategies
- Optional: observability and performance metrics
- Clean interface design for pluggable storage
