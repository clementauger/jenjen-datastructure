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
$ make bench
(cd examples/heap; go test -bench=. -benchmem -count=2)
goos: linux
goarch: amd64
pkg: github.com/clementauger/jenjen-datastructure/examples/heap
BenchmarkMinIntHeapDup-4      	   10000	    107857 ns/op	       8 B/op	       0 allocs/op
BenchmarkMinIntHeapDup-4      	   11208	    107012 ns/op	       7 B/op	       0 allocs/op
BenchmarkMinIntHeapNoDup-4    	    1945	    607806 ns/op	      42 B/op	       0 allocs/op
BenchmarkMinIntHeapNoDup-4    	    1950	    608287 ns/op	      42 B/op	       0 allocs/op
BenchmarkRegularHeapDup-4     	    3565	    324298 ns/op	      22 B/op	       0 allocs/op
BenchmarkRegularHeapDup-4     	    3586	    324760 ns/op	      22 B/op	       0 allocs/op
BenchmarkRegularHeapNoDup-4   	     561	   2115989 ns/op	  156050 B/op	   19488 allocs/op
BenchmarkRegularHeapNoDup-4   	     558	   2125410 ns/op	  156051 B/op	   19488 allocs/op
PASS
ok  	github.com/clementauger/jenjen-datastructure/examples/heap	10.915s
```
