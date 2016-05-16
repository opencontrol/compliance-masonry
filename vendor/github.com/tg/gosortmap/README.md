# gosortmap [![GoDoc](https://godoc.org/github.com/tg/gosortmap?status.svg)](https://godoc.org/github.com/tg/gosortmap) [![Build Status](https://travis-ci.org/tg/gosortmap.svg?branch=master)](https://travis-ci.org/tg/gosortmap)
Sort maps in Go by keys or values. Works with most built-in types; own comparator can
be provided to support custom types and ordering.
## Example
```go
m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
for _, e := range sortmap.ByValue(m) {
	fmt.Printf("%s\t%d\n", e.Key, e.Value)
}
// Output:
// banana	1
// daikon	2
// cabbage	3
// apple	4

fmt.Println(sortmap.ByValueDesc(m).Top(2))
// Output:
// [{apple 4} {cabbage 3}]
```
## Benchmark
This package favors convenience over the speed, so if the latter is preferable,
you should go with an intermediate structure implementing `sort.Interface` and use
`sort.Sort` directly. As there is an extra call on every comparison in this package
and functions are operating on `interface{}`, the execution (for `map[string]int`)
is about 4x slower than providing a direct, manual solution:
```
BenchmarkManualSorted   2000	   1004797 ns/op
BenchmarkSortSorted	     200	   6591329 ns/op

BenchmarkManualFunc	     300	   4313895 ns/op
BenchmarkSortFunc	     100	  17715101 ns/op

BenchmarkManualKey	     300	   4707508 ns/op
BenchmarkSortKey	     100	  18170250 ns/op
```
