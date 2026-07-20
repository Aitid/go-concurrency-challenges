You are building a service that needs to process a large number of independent jobs.

Each job takes some time to complete. You cannot process all jobs concurrently because it could overload downstream services.

Implement a worker pool in Go that:
- Accepts jobs from a producer
- Processes jobs concurrently
- Limits the number of concurrent workers
- Returns results
- Supports cancellation using `context.Context`
- Shuts down cleanly without goroutine leaks
# Requirements

```go
type Job struct {
    ID int
}

type Result struct {
    JobID int
    Value string
}

type WorkerPool struct {
    // your fields
}

func NewWorkerPool(workerCount int) *WorkerPool

func (p *WorkerPool) Submit(job Job) error

func (p *WorkerPool) Results() <-chan Result

func (p *WorkerPool) Start(ctx context.Context)

func (p *WorkerPool) Stop()
```
# Expected behavior
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

pool := NewWorkerPool(3)

pool.Start(ctx)

for i := 0; i < 10; i++ {
    pool.Submit(Job{
        ID: i,
    })
}

for result := range pool.Results() {
    fmt.Println(result)
}

pool.Stop()
```

- What happens if Submit is called after Stop?
- How do you add a queue limit?
- How do you return errors?
- How do you preserve job ordering?
- How would you dynamically change worker count?
