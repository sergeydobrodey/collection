# 🚀 Collection - Powerful Go Utilities for Modern Development

[![Go Reference](https://pkg.go.dev/badge/github.com/sergeydobrodey/collection.svg)](https://pkg.go.dev/github.com/sergeydobrodey/collection)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/sergeydobrodey/collection) 
[![Go Report Card](https://goreportcard.com/badge/github.com/sergeydobrodey/collection)](https://goreportcard.com/report/github.com/sergeydobrodey/collection)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/sergeydobrodey/collection/master/LICENSE) 
[![Build Status](https://github.com/sergeydobrodey/collection/actions/workflows/test.yml/badge.svg)](https://github.com/sergeydobrodey/collection/actions/workflows/test.yml)
[![Go Coverage](https://github.com/sergeydobrodey/collection/wiki/coverage.svg)](https://raw.githack.com/wiki/sergeydobrodey/collection/coverage.html)


**Collection** is a comprehensive Go package that brings functional programming paradigms to Go with type-safe generics. Transform your slice and map operations from verbose loops into elegant, chainable operations that are both performant and readable.

## ✨ Why Choose Collection?

- **🎯 Type-Safe Generics**: Full Go 1.18+ generics support with compile-time type safety
- **⚡ High Performance**: Optimized implementations with minimal memory allocations
- **🔒 Thread-Safe**: Built-in concurrent-safe map operations
- **🔗 Functional Style**: Chainable operations inspired by functional programming
- **📦 Zero Dependencies**: Pure Go implementation with no external dependencies
- **🧪 Well Tested**: Comprehensive test suite ensuring reliability

## 🚀 Quick Start

```bash
go get github.com/sergeydobrodey/collection
```

```go
import "github.com/sergeydobrodey/collection"
```

## 💡 Core Features

### 🔄 Slice Transformations

Transform your data with ease using functional programming patterns:

```go
// Transform slice elements
users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
names := collection.TransformBy(users, func(u User) string { return u.Name })
// Result: ["Alice", "Bob"]

// Filter with predicates
adults := collection.FilterBy(users, func(u User) bool { return u.Age >= 18 })

// Chain operations naturally
activeUserNames := collection.TransformBy(
    collection.FilterBy(users, func(u User) bool { return u.Active }),
    func(u User) string { return u.Name },
)
```

### 🗺️ Powerful Map Operations

```go
// Convert slices to maps instantly
userMap := collection.SliceToMap(users, func(u User) int { return u.ID })
// Result: map[1:User{1, "Alice"} 2:User{2, "Bob"}]

// Transform map values
uppercaseNames := collection.MapTransformBy(userMap, func(u User) string {
    return strings.ToUpper(u.Name)
})
```

### 🔒 Thread-Safe Concurrent Operations

```go
// Thread-safe map for concurrent access
syncMap := &collection.SyncMap[string, int]{}

// Safe concurrent operations
syncMap.Store("key", 42)
value, ok := syncMap.Load("key")
syncMap.CompareAndSwap("key", 42, 100)
```

### 🌊 Async & Channel Operations

```go
// Merge multiple channels elegantly
worker1Done := make(chan struct{})
worker2Done := make(chan struct{})
allDone := collection.ChannelsMerge(worker1Done, worker2Done)

// Async transformations with context
results, err := collection.AsyncTryTransformBy(ctx, urls, fetchData)
```

## 📚 Comprehensive API Reference

### Slice Operations
| Function | Description | Example Use Case |
|----------|-------------|------------------|
| `TransformBy` | Transform elements to new type | Convert structs to IDs |
| `FilterBy` | Filter elements by predicate | Get active users |
| `Aggregate` | Reduce slice to single value | Sum, concatenate, etc. |
| `GroupBy` | Group elements by key function | Group users by department |
| `ChunkBy` | Split slice into smaller chunks | Batch processing |
| `Distinct` | Remove duplicates | Unique IDs |
| `Intersection` | Find common elements | Common interests |
| `Difference` | Find unique elements | Missing items |

### Validation & Checks
| Function | Description | Example Use Case |
|----------|-------------|------------------|
| `All` | Check if all elements match predicate | All users validated |
| `Any` | Check if any element matches predicate | Any errors present |
| `Contains` | Check if slice contains element | User exists |
| `Equal` | Compare two slices for equality | Data consistency |

### Map Operations
| Function | Description | Example Use Case |
|----------|-------------|------------------|
| `MapTransformBy` | Transform map values | Normalize data |
| `MapFilterBy` | Filter map entries | Active sessions only |
| `MapToSlice` | Convert map to slice | Extract values |
| `MapKeys` / `MapValues` | Extract keys or values | Get all IDs |

### Async & Concurrency
| Function | Description | Example Use Case |
|----------|-------------|------------------|
| `AsyncTransformBy` | Parallel transformations | Concurrent API calls |
| `AsyncTryTransformBy` | Parallel with error handling | Safe concurrent operations |
| `ChannelsMerge` | Combine multiple channels | Wait for multiple workers |

## 🎯 Real-World Examples

### Data Processing Pipeline

```go
type Order struct {
    ID       int
    UserID   int
    Amount   float64
    Status   string
    Created  time.Time
}

// Complex data processing in a readable pipeline
func ProcessRecentOrders(orders []Order) map[int]float64 {
    // Get recent, completed orders grouped by user
    recentCompleted := collection.FilterBy(orders, func(o Order) bool {
        return o.Status == "completed" && 
               time.Since(o.Created) < 24*time.Hour
    })
    
    // Group by user and calculate totals
    return collection.MapTransformBy(
        collection.GroupBy(recentCompleted, func(o Order) int { return o.UserID }),
        func(userOrders []Order) float64 {
            return collection.Aggregate(userOrders, func(sum float64, o Order) float64 {
                return sum + o.Amount
            })
        },
    )
}
```

### Concurrent Data Fetching

```go
func FetchUserProfiles(userIDs []int) ([]UserProfile, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Fetch all profiles concurrently with error handling
    return collection.AsyncTryTransformBy(ctx, userIDs, func(ctx context.Context, id int) (UserProfile, error) {
        return fetchUserProfile(ctx, id)
    })
}
```

### Thread-Safe Cache

```go
type UserCache struct {
    users *collection.SyncMap[int, User]
}

func (c *UserCache) GetOrFetch(userID int) (User, error) {
    // Try to load from cache first
    if user, ok := c.users.Load(userID); ok {
        return user, nil
    }
    
    // Fetch and store atomically
    user, err := fetchUser(userID)
    if err != nil {
        return User{}, err
    }
    
    actual, loaded := c.users.LoadOrStore(userID, user)
    return actual, nil
}
```

## 🏆 Performance & Benchmarks

Collection is designed for performance with minimal allocations:

- **Memory Efficient**: Reuses slices where possible, minimal allocations
- **CPU Optimized**: Efficient algorithms with O(n) complexity for most operations
- **Concurrent Safe**: Lock-free operations where possible in SyncMap

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **🐛 Report Bugs**: Open an issue with reproduction steps
2. **💡 Feature Requests**: Suggest new utility functions
3. **📖 Documentation**: Improve examples and documentation
4. **🧪 Tests**: Add test cases for edge cases

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Show Your Support

If this library helps you build better Go applications, please consider:
- ⭐ Starring the repository
- 🐛 Reporting issues
- 📢 Sharing with the Go community
- 💡 Contributing new features

---

**Made with ❤️ for the Go community**

*Transform your Go code from imperative to functional - one collection at a time.*
