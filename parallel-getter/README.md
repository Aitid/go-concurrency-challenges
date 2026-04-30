### Problem Description
You are given an interface:
```go
type Getter interface {
    Get(ctx context.Context, address, key string) (string, error)
}
```

You have multiple `Getter` instances, each pointing to a remote source. Implement:

```go
func DistributiveGet(ctx context.Context, getter Getter, addresses []string, key string) (string, error)
```

### Requirements
1. Call Getter.Get() for all addresses in parallel.
2. Return the first successful response immediately.
3. If all requests fail, return an error.
4. Support cancellation via context.Context — remaining requests should stop once a result is obtained or timeout occurs.
### Optional Enhancements (Senior Level)
- Aggregate all errors and return detailed failure information if all requests fail.
- Implement timeouts per request independent of global context.
- Limit maximum concurrency if the list of getters is very large.
### Hints
- Use a result channel to collect successful responses.
- Launch a goroutine per Getter to call Get().
- Use select to wait for the first success or context cancellation.
- Cancel other goroutines when the first successful result arrives using context.WithCancel.
- Be careful to avoid goroutine leaks.
### Skills Tested
- Concurrency primitives: goroutines, channels, select.
- Context usage: cancellation, timeout handling.
- Parallelism control: efficient coordination of multiple concurrent requests.
- Error handling: propagate errors correctly if all sources fail.
- Optional: concurrency limits and structured error aggregation.
### Execution Flow (Illustration)

```
  +-----------------+
  |   Main goroutine |
  +-----------------+
          |
          v
   Launch goroutines for each Getter
          |
          v
+-----------------+   +-----------------+   +-----------------+
| Getter 1        |   | Getter 2        |   | Getter 3        |
| Get(key,address)|   | Get(key,address)|   | Get(key,address)|
+-----------------+   +-----------------+   +-----------------+
          |                 |                  |
          v                 v                  v
       Send result       Send result        Send result
          |                 |                  |
          +-------- select -------------------+
                       |
         +-------------+--------------+
         |                            |
     First success                  ctx done / all fail
```

