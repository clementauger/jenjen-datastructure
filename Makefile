examples:
	(cd examples/heap;\
		jenjen -template=github.com/clementauger/jenjen-datastructure/heap \
		"U => -, T => -, Heap=>MinIntHeap , MinIntHeap:U=>int, MinIntHeap:T=> minInt"
	)
bench:
	(cd examples/heap; go test -bench=. -benchmem -count=2)
