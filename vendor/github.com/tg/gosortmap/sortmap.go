// Package sortmap allows for sorting maps by a custom comparator.
// For convenience, functions sorting by keys or values in ascending or descending
// order are provided – these can deal with limited types only, which are:
// bool, all built-in numerical types and string, time.Time.
//
// Functions provided by this package panic when non-map type is passed for sorting
// or, in case of the key/value sorters, the underyling type is not supported.
package sortmap

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// Item is a key-value pair representing element in the map
type Item struct {
	Key, Value interface{}
}

// Less compares two map elements and returns true if x < y
type Less func(x, y Item) bool

// flatmap is a flattened map with a comparator to be used with sort
type flatmap struct {
	items []Item
	less  Less
}

func newFlatMap(m interface{}, less Less) *flatmap {
	mv := reflect.ValueOf(m)
	keys := mv.MapKeys()
	fm := &flatmap{items: make([]Item, len(keys)), less: less}
	for n := range keys {
		fm.items[n] = Item{keys[n].Interface(), mv.MapIndex(keys[n]).Interface()}
	}
	return fm
}

func (m *flatmap) Len() int {
	return len(m.items)
}
func (m *flatmap) Less(i, j int) bool {
	return m.less(m.items[i], m.items[j])
}
func (m *flatmap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Items is a slice of map elements (key-value pairs)
type Items []Item

// Top returns slice of up to n leading elements
func (r Items) Top(n int) Items {
	if n > len(r) {
		n = len(r)
	}
	return r[:n]
}

// ByFunc sorts map using a provided comparator
func ByFunc(m interface{}, c Less) Items {
	fm := newFlatMap(m, c)
	sort.Sort(fm)
	return fm.items
}

// ByKey sorts map by keys in the ascending order
func ByKey(m interface{}) Items {
	ls := getLess(reflect.ValueOf(m).Type().Key())
	return ByFunc(m, func(x, y Item) bool { return ls(x.Key, y.Key) })
}

// ByKeyDesc sorts map by keys in the descending order
func ByKeyDesc(m interface{}) Items {
	ls := getLess(reflect.ValueOf(m).Type().Key())
	return ByFunc(m, func(x, y Item) bool { return ls(y.Key, x.Key) })
}

// ByValue sorts map by values in the ascending order
func ByValue(m interface{}) Items {
	ls := getLess(reflect.ValueOf(m).Type().Elem())
	return ByFunc(m, func(x, y Item) bool { return ls(x.Value, y.Value) })
}

// ByValueDesc sorts map by values in the descending order
func ByValueDesc(m interface{}) Items {
	ls := getLess(reflect.ValueOf(m).Type().Elem())
	return ByFunc(m, func(x, y Item) bool { return ls(y.Value, x.Value) })
}

// getLess returns default comparator for a type
func getLess(t reflect.Type) (f func(x, y interface{}) bool) {
	switch t.Kind() {
	case reflect.Bool:
		f = func(x, y interface{}) bool { return !x.(bool) && y.(bool) }
	case reflect.Int:
		f = func(x, y interface{}) bool { return x.(int) < y.(int) }
	case reflect.Int8:
		f = func(x, y interface{}) bool { return x.(int8) < y.(int8) }
	case reflect.Int16:
		f = func(x, y interface{}) bool { return x.(int16) < y.(int16) }
	case reflect.Int32:
		f = func(x, y interface{}) bool { return x.(int32) < y.(int32) }
	case reflect.Int64:
		f = func(x, y interface{}) bool { return x.(int64) < y.(int64) }
	case reflect.Uint:
		f = func(x, y interface{}) bool { return x.(uint) < y.(uint) }
	case reflect.Uint8:
		f = func(x, y interface{}) bool { return x.(uint8) < y.(uint8) }
	case reflect.Uint16:
		f = func(x, y interface{}) bool { return x.(uint16) < y.(uint16) }
	case reflect.Uint32:
		f = func(x, y interface{}) bool { return x.(uint32) < y.(uint32) }
	case reflect.Uint64:
		f = func(x, y interface{}) bool { return x.(uint64) < y.(uint64) }
	case reflect.Float32:
		f = func(x, y interface{}) bool { return x.(float32) < y.(float32) }
	case reflect.Float64:
		f = func(x, y interface{}) bool { return x.(float64) < y.(float64) }
	case reflect.String:
		f = func(x, y interface{}) bool { return x.(string) < y.(string) }
	case reflect.TypeOf(time.Time{}).Kind():
		f = func(x, y interface{}) bool { return x.(time.Time).Before(y.(time.Time)) }
	default:
		panic(fmt.Sprintf("sortmap: unsupported type: %s", t))
	}
	return
}
