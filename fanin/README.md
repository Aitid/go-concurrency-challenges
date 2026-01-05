# Fan-in (Concurrency Pattern) Challenge

This pattern merges multiple channels into a single channel.  
The goal is to implement the function:

```go
func Merge[T any](ctx context.Context, chs ...<-chan T) <-chan T {}
```
so that it passes all tests, including edge cases:
- properly closes the output channel (out)
- handles context cancellation
- works with nil channels
- works with slow consumers
- handles zero input channels
- does not leak goroutines

---

## How to run tests
1. Open a terminal and navigate to the root of the repository:
```bash
cd path/to/go-concurrency-challenges
```
2. Pull dependencies (if not already done):
```bash
go mod tidy
```
3. Run all fan-in tests:
```bash
go test -v -race ./fanin/
```
4. Run a specific test (e.g., happy path):
```bash
go test -v -race ./fanin/ -run TestMerge/happy_path
```
All tests are designed to verify real understanding of the pattern.
Try to implement Merge without hints or LLM assistance to test your knowledge level!

