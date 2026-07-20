You need to implement an analytics event collector that is used by multiple components of the application.

The collector receives events from different goroutines and periodically sends them to an external analytics service.

The goal is to minimize the number of outgoing requests by batching events.

## Requirements
Implement an `AnalyticsCollector` with the following behavior:
- Events can be submitted concurrently from multiple goroutines.
- Events should be buffered in memory.
- When the buffer reaches a configured size, events should be sent immediately.
- Events should also be flushed periodically based on a configured interval.
- The collector should not block callers for a long time while sending events.
- The collector should support graceful shutdown.

## Interface
Example API:
```go
type Event struct {
    Name      string
    Timestamp time.Time
    Payload   map[string]string
}

type AnalyticsCollector interface {
    Track(event Event) error
    Start()
    Stop() error
}
```
## Additional requirements
The implementation should handle:
- concurrent calls to `Track`;
- failures while sending events;
- preventing goroutine leaks.
## Questions to consider
- What happens if the analytics service is unavailable?
- Should failed events be retried?
- What happens if the application shuts down while there are unsent events?
- How large can the internal buffer grow?
