### Problem Description
Implement a parking lot system with a fixed number of spaces that supports concurrent access:
- `Park()` occupies a free space.
- `Leave()` frees a space.
- If no spaces are available, requests **wait** (or fail if cancelled).
- Optional: collect metrics and support priorities.
### Requirements
1. Initialize parking lot with a maximum number of spaces.
2. `Park()` occupies a free space if available.
3. `Leave()` releases a space.
4. If no spaces are available, `Park()` blocks until a space is free.
5. `Park(ctx)` tries to occupy a free space:
    - Succeeds if a space is free.
    - Waits until a space is free.
    - Returns an error if context is cancelled or timeout expires.
6. `Leave()` releases a space.
7. Non-blocking methods (e.g., `TryPark()`) return immediately.
8. Thread-safe under high concurrency.    
### Example API
```go
type Parking struct { ... }

func NewParking(capacity int) *Parking
func (p *Parking) Park(ctx context.Context) error
func (p *Parking) TryPark() bool
func (p *Parking) Leave()
```
### Execution Flow (Illustration)
```
Incoming car → Park(ctx) / TryPark()
            |
            v
       Is space available?
        /         \
      Yes          No
       |            |
  Occupy space   Wait until space frees or context cancels
       |
      Metrics updated (occupied count, response time, timeouts)
       |
Car leaves → Leave() → frees space → unblocks waiting cars
```
### Hints
- Use a channel as a semaphore to track free spaces (`chan struct{}`)
- Use mutexes if you store metrics or internal counters.
- `TryPark()` can select on the channel with default to avoid blocking.
- `Park(ctx)` can select on the channel and ctx.Done() to handle cancellation.
- For priority parking, consider separate queues or channels for VIPs.
- Track **metrics** by updating counters after acquiring/releasing a space, not before.
### Skills Tested
- Concurrency patterns: channels, semaphores, mutexes
- Thread-safe state management
- `context.Context` usage for timeouts and cancellations
- Designing non-blocking methods
- Metrics collection under concurrency
