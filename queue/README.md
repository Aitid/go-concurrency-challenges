# 🟢 Concurrent Queue
### Problem Description
Implement a thread-safe queue with a fixed maximum capacity. The queue must support concurrent `Push`, `Pop`, and `Peek` operations from multiple goroutines without data races or corruption.

```go
type Queue interface {
    Push(val int) error  // Adds element to back; returns ErrQueueFull if full
    Pop() (int, error)   // Removes and returns front element; returns ErrQueueEmpty if empty
    Peek() (int, error)  // Returns front element without removing; returns ErrQueueEmpty if empty
}
```
### Requirements
- Fixed maximum size defined at initialization.
- Concurrent-safe: multiple producers and consumers may call `Push`, `Pop`, and `Peek` concurrently.
- Non-blocking behavior: operations should return immediately; do not wait for space or items
- Minimal overhead.
- `Push(val int)`:
    - Adds element to the back of the queue.
    - Returns `ErrQueueFull` if the queue is full.
- `Pop()`:
    - Removes and returns the front element.
    - Returns `ErrQueueEmpty` if the queue is empty.
- Peek():
    - Returns the front element without removing it.
    - Returns `ErrQueueEmpty` if the queue is empty.
### Hints
- Use sync.Mutex to protect internal state.
- Store elements in a slice or ring buffer.
- Keep track of current size and head/tail indices.
- Avoid holding locks longer than necessary.
### Skills Tested
- Safe shared-state management with `sync.Mutex`.
- Understanding of race conditions.
- Data structure implementation (queue / ring buffer).
- Clean API design and error handling.

