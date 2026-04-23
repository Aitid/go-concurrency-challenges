### Problem Description

Implement a concurrent Token Bucket rate limiter:
- Controls the rate of requests by consuming tokens from a bucket.
- Tokens refill automatically at a fixed rate, up to bucket capacity.
- Allows multiple goroutines to check and consume tokens safely.

You should implement:
```go
type TokenBucket interface {
    Allow() bool                // Returns true if a token is available, consumes it
    Wait(ctx context.Context) error  // Optional: blocks until a token is available or context cancelled
}
```
### Requirements

1. Thread-safe: multiple goroutines may call `Allow()` or `Wait()` concurrently.
2. Token consumption: `Allow()` consumes one token if available; otherwise, returns false.
3. Automatic refill: tokens are added at a fixed interval up to the bucket capacity.
4. Optional blocking: `Wait(ctx)` should block until a token is available or context expires.
5. Performance: limiter should scale under high concurrency without contention.
### Hints
- Use `sync.Mutex` or atomic counters for thread safety.
- Use `time.Ticker` or `time.After` for periodic refill.
- Consider storing the last refill timestamp to calculate tokens efficiently.
- For `Wait(ctx)`, use channels or `select` to respect context cancellation.

### Optional Enhancements (Senior Level)
- Allow dynamic configuration of bucket capacity or refill rate.
- Track metrics: number of allowed requests, rejected requests, average wait time.
- Support distributed usage via Redis or another shared storage.
### Skills Tested
- Concurrency patterns: goroutines, mutexes, atomic operations, channels.
- Thread-safe state management under high load.
- Time-based logic and scheduling in Go.
- Context cancellation handling (`Wait(ctx)`).
- API design for rate limiting and monitoring.

### Execution Flow (Illustration)

```
            +----------------+
Caller 1    | Allow()        |
Caller 2    | Allow()        |
            +--------+-------+
                     |
            Check token bucket
                     |
            +--------+--------+
            | Tokens > 0 ?   |
            |  Yes → consume |
            |  No  → return false / wait
            |
     +-------+--------+
     | Optional Wait()|
     | Block until token available or ctx expires
     +-------+--------+
                     |
       Refill tokens at fixed rate, up to capacity
```
