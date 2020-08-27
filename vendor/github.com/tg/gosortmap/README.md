# gosortmap [![GoDoc](https://godoc.org/github.com/tg/gosortmap?status.svg)](https://godoc.org/github.com/tg/gosortmap) [![Build Status](https://circleci.com/gh/tg/gosortmap.png?style=shield&circle-token=:circle-token)](https://travis-ci.org/tg/gosortmap)
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
`sort.Sort` directly. Here are the results for sorting `map[int]int` with 1000 values,
indicating manual sorting can be up to 2x faster:

```
BenchmarkByKey-8           	     300	   4094204 ns/op
BenchmarkByKey_manual-8    	     500	   2367496 ns/op

BenchmarkByFunc-8          	     500	   3623426 ns/op
BenchmarkByFunc_manual-8   	    1000	   2122365 ns/op
```
(go 1.11.2, i7-6820HQ)
