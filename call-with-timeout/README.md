# Problem Description
You are given a function simulating a long-running API request:
```go
func ExternalAPIRequest() (string, error) {
    // Simulate a slow API request
    time.Sleep(5 * time.Second)
    return "data from API", nil
}
```
Implement a wrapper that adds timeout and cancellation support.
### Requirements
1. Call fn in a way that does not block indefinitely.
2. Return the result of fn if it completes within the timeout.
3. Return context.DeadlineExceeded if the timeout expires.
4. Respect context cancellation: If ctx is canceled, abort the call and return an error.
5. Optional: support retries, logging, and execution time metrics.

Implement a function:
```go
func CallWithTimeout(ctx context.Context, fn func() (string, error), timeout time.Duration) (string, error)
```
## Hints
- Use goroutines to run `fn` concurrently.
- Communicate results via channels.
- Use select to wait on:
    - result channel,
    - error channel,
    - ctx.Done() (for timeout/cancellation).
- Avoid goroutine leaks.

## Optional Enhancements
- Retry failed requests a configurable number of times.
- Log start/end time for each attempt.
- Accept a custom context.Context to allow external cancellation.

## Skills Tested
- Goroutines & channels
- context.Context usage (timeouts & cancellation)
- Concurrency safety & non-blocking design
- Optional: retry logic, metrics, logging
