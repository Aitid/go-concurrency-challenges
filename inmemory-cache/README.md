### Problem Description
Implement an in-memory key-value cache that supports high concurrency and efficient access:
- Concurrent reads and writes from multiple goroutines.
- Reduce lock contention using sharding (split cache into segments).
- Each entry has a TTL (time-to-live).
- Cache size is limited (e.g., via LRU or per-shard limits).
You should implement:
```go
type Cache interface {
    Get(key string) (value any, ok bool)
    Set(key string, value any, ttl time.Duration)
}
```
### Requirements
- `Get(key)` and `Set(key, value, ttl)` methods.
- Thread-safe: use `sync.RWMutex` or `sync.Mutex` per shard.
- Shard keys by hash: `shard := hash(key) % shardCount`.
- Ignore / evict expired keys on access.
- Prevent unbounded growth: enforce size or LRU limits per shard.
### Hints
- Use `struct shard { sync.RWMutex; items map[string]cacheItem }`
- Split cache into multiple shards to reduce lock contention.
- Store TTL per entry (`expiresAt time.Time`).
- Use `container/list` for LRU per shard.
- Optional: channels, goroutines, or singleflight for concurrency.
### Skills Tested
- Concurrency and thread-safe design in Go.
- Sharding and lock minimization for high-throughput access.
- TTL and cache eviction strategies (LRU / size-limited).
- Optional: thundering herd prevention with singleflight.
- Interface-oriented design for pluggable storage.
### Execution Flow (Illustration)
```
                +-------------------+
Caller 1        | Get("key1")       |
Caller 2        | Set("key2", val)  |
Caller 3        | Get("key1")       |
                +---------+---------+
                          |
                   Determine shard:
                shard := hash(key) % shardCount
                          |
                +---------+---------+
                |   Acquire shard lock  |
                | (RLock for Get, Lock for Set) |
                +---------+---------+
                          |
       +------------------+------------------+
       | Value exists? (and not expired)    |
       |      Yes → return value            |
       |      No → fetch / insert / wait   |
       +------------------+------------------+
                          |
                  Update shard map / LRU
                          |
                Release shard lock
                          |
           Return value to caller(s)

```
