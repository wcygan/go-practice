# Practice

Practicing Data Structures, Algorithms, Concurrency, and more in [Go](https://go.dev/)!

## Table of Contents

- **[Build and Run](#build-and-run)**<br>
- **[References](#references)**<br>

## Build and Run

From the go-practice root directory, `Practice/go-practice`, you can execute these commands:

### Execute all tests:

```
$ go test ./...
```

### Execute a specific test (with coverage for the entire project):

```
$ go test ./... -run TestName -cover
```

Tip: adding the `-cover` flag tells the test runner to keep track of observed program coverage.

## References

I'm using the following material as a reference:

1. [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
2. [The Go Programming Language](https://www.gopl.io/)
3.  [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)
4. [Network Programming with Go](https://nostarch.com/networkprogrammingwithgo)
5. [Distributed Services with Go](https://pragprog.com/titles/tjgo/distributed-services-with-go/)