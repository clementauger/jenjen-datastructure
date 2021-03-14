# jenjen-datastructure

A repository that exposes templates to generate faster datastructure.

# Heap

Generates a heap.

```sh
jenjen -template=github.com/clementauger/jenjen-datastructure/heap \
  - "U => -, T => -, Heap=>MinIntHeap , MinIntHeap:U=>int, MinIntHeap:T=> minInt"
```

You must provide a type that implements a `type Lesser { Less(_ T) bool }`, for example :

```go
type minInt int

func (s minInt) Less(v minInt) bool {
	return s < v
}
```

### Results

```sh
$ (cd examples/heap; go test -bench=. -benchmem -count=2)
goos: linux
goarch: amd64
pkg: github.com/clementauger/jenjen-datastructure/examples/heap
BenchmarkMinIntHeap-4    	    9439	    108133 ns/op	       8 B/op	       0 allocs/op
BenchmarkMinIntHeap-4    	   10000	    107884 ns/op	       8 B/op	       0 allocs/op
BenchmarkRegularHeap-4   	    3523	    325731 ns/op	      23 B/op	       0 allocs/op
BenchmarkRegularHeap-4   	    3433	    321977 ns/op	      23 B/op	       0 allocs/op
PASS
ok  	github.com/clementauger/jenjen-datastructure/examples/heap	4.452s
```
