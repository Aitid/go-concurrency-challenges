### Problem Description
You are given a binary tree structure:
```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```
- tree.New(k) constructs a sorted tree with values k, 2k, 3k, ..., 10k.
- The tree structure may be randomized, but in-order traversal produces sorted values.

Your tasks:
### 1. Implement Walk

```go
func Walk(t *Tree, ch chan int)
```
- Traverses the tree `t` in-order.
- Sends each value to ch.
- Closes the channel after traversal.
Example Usage:
```go
ch := make(chan int)
go Walk(tree.New(1), ch)
for v := range ch {
    fmt.Println(v) // 1, 2, 3, ..., 10
}
```
### 2. Implement Same
```go
func Same(t1, t2 *Tree) bool
```
- Determines whether `t1` and `t2` contain the same values, regardless of structure.
- Uses `Walk` to read tree values from channels and compares them.
Examples:
```go
Same(tree.New(1), tree.New(1)) // true
Same(tree.New(1), tree.New(2)) // false
```
### Hints
- Use recursive in-order traversal for `Walk`.
- Always close the channel after sending all values.
- For Same, use two channels and iterate through both simultaneously, comparing values one by one.
- Return false early if any mismatch occurs.
- Optional: use buffered channels to avoid blocking if needed.
### Skills Tested
- Goroutines & Channels: run `Walk` in its own goroutine.
- Tree Traversal: recursive in-order traversal.
- Synchronization: closing channels, iterating until done.
- Comparison Logic: correctly compare sequences from two trees.
- Go Idioms: clean, idiomatic Go with concurrency primitives.
### Execution Flow (Illustration)

```
       Tree
        |
       Walk(t, ch) ---> goroutine sends values to channel
        |
   +----+----+
   |         |
  1,2,3,... 10
        |
  Main goroutine reads from channel
```
For Same, two Walk goroutines produce two channels → compare values as they arrive.

# Source
https://go.dev/tour/concurrency/7
