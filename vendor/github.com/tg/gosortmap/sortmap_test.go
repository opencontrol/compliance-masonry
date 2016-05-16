package sortmap_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/tg/gosortmap"
)

func ExampleByKey() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByKey(m) {
		fmt.Printf("%s\t%d\n", e.Key, e.Value)
	}
	// Output:
	// apple	4
	// banana	1
	// cabbage	3
	// daikon	2
}

func ExampleByKeyDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByKeyDesc(m) {
		fmt.Printf("%s\t%d\n", e.Key, e.Value)
	}
	// Output:
	// daikon	2
	// cabbage	3
	// banana	1
	// apple	4
}

func ExampleByValue() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByValue(m) {
		fmt.Printf("%s\t%d\n", e.Key, e.Value)
	}
	// Output:
	// banana	1
	// daikon	2
	// cabbage	3
	// apple	4
}

func ExampleByValueDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByValueDesc(m) {
		fmt.Printf("%s\t%d\n", e.Key, e.Value)
	}
	// Output:
	// apple	4
	// cabbage	3
	// daikon	2
	// banana	1
}

func ExampleTopElements() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	fmt.Println(sortmap.ByValueDesc(m).Top(2))
	// Output:
	// [{apple 4} {cabbage 3}]
}

func ExampleByTime() {
	p := func(s string) time.Time {
		t, err := time.Parse(time.Kitchen, s)
		if err != nil {
			panic(err)
		}
		return t
	}
	m := map[time.Time]int{p("3:04PM"): 4, p("6:48AM"): 2, p("1:10PM"): 3, p("1:10AM"): 1}
	for _, e := range sortmap.ByKey(m) {
		fmt.Printf("%s\t%d\n", e.Key.(time.Time).Format(time.Kitchen), e.Value)
	}
	// Output:
	// 1:10AM	1
	// 6:48AM	2
	// 1:10PM	3
	// 3:04PM	4
}

var benchMap = func() map[int]int {
	m := make(map[int]int)
	for n := 0; n < 10000; n++ {
		m[rand.Int()] = rand.Int()
	}
	return m
}()

type kv struct{ k, v int }
type kvs []kv

func (m kvs) Len() int           { return len(m) }
func (m kvs) Less(i, j int) bool { return m[i].k < m[j].k }
func (m kvs) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type kvs_nosort []kv

func (m kvs_nosort) Len() int           { return len(m) }
func (m kvs_nosort) Less(i, j int) bool { return false }
func (m kvs_nosort) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func BenchmarkManualSorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(kvs_nosort, 0, len(benchMap))
		for k, v := range benchMap {
			m = append(m, kv{k, v})
		}
		sort.Sort(m)
	}
}

func BenchmarkManualFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(kvs, 0, len(benchMap))
		for k, v := range benchMap {
			m = append(m, kv{k, v})
		}
		sort.Sort(m)
	}
}

func BenchmarkManualKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		keys := make([]int, 0, len(benchMap))
		for k, _ := range benchMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		values := make([]int, len(keys))
		for n := range keys {
			values[n] = benchMap[keys[n]]
		}
	}
}

func BenchmarkSortSorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.Item) bool { return false })
	}
}

func BenchmarkSortFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.Item) bool { return x.Key.(int) < y.Key.(int) })
	}
}

func BenchmarkSortKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByKey(benchMap)
	}
}
