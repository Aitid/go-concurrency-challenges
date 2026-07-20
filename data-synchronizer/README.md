You are building a service that synchronizes data from multiple external sources into one destination.

For example:
- Fetch users from multiple APIs
- Merge results
- Remove duplicates
- Write the final data to a database

Each source can be slow or fail independently.

Implement a concurrent data synchronizer in Go.

The synchronizer must:
- Run multiple sync tasks concurrently
- Limit concurrency
- Collect results from all workers
- Handle errors
- Support cancellation using `context.Context`
- Avoid goroutine leaks
# Example scenario
You have several data sources:

```
          Source A
             |
             |
          Source B -----> Synchronizer -----> Database
             |
             |
          Source C
```

Each source returns user data.
Input:
```go
sources := []Source{
    APIUsers,
    LegacyUsers,
    PartnerUsers,
}
```
Output:
```go
[]User{
    {ID:1, Name:"Alice"},
    {ID:2, Name:"Bob"},
}
```
# API

```go
type User struct {
    ID   int
    Name string
}

type Source interface {
    Fetch(ctx context.Context) ([]User, error)
}


type Synchronizer struct {
    // fields
}


func NewSynchronizer(
    workerCount int,
) *Synchronizer


func (s *Synchronizer) Sync(
    ctx context.Context,
    sources []Source,
) ([]User, error)
```

# Usage
```go
ctx := context.Background()

syncer := NewSynchronizer(3)

users, err := syncer.Sync(
    ctx,
    []Source{
        sourceA,
        sourceB,
        sourceC,
    },
)

if err != nil {
    log.Fatal(err)
}

fmt.Println(users)
```

What if sources are continuously changing?
How do you retry failed sources?
How do you make it resumable?
How would you monitor it?
