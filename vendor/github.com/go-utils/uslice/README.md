# uslice
--
    import "github.com/go-utils/uslice"

Go programming helpers for common 'pseudo-generic' typed-slice needs.

## Usage

#### func  BoolAppendUnique

```go
func BoolAppendUnique(ref *[]bool, v bool)
```
Appends `v` to `*ref` only if `*ref` does not already contain `v`.

#### func  BoolAppendUniques

```go
func BoolAppendUniques(ref *[]bool, vals ...bool)
```
Appends each value in `vals` to `*ref` only `*ref` does not already contain it.

#### func  BoolAt

```go
func BoolAt(slice []bool, val bool) int
```
Returns the position of `val` in `slice`.

#### func  BoolConvert

```go
func BoolConvert(src []interface{}, sparse bool) (dst []bool)
```
Converts `src` to `dst`.

If `sparse` is `true`, then only successfully converted `bool` values are placed
in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in
length or indices.

If `sparse` is `false`, `dst` has the same length as `src` and non-convertable
values remain zeroed.

#### func  BoolEach

```go
func BoolEach(sl []bool, apply ...func(bool) bool) []bool
```
Sets each `bool` in `sl` to the result of passing it to each `apply` func.
Although `sl` is modified in-place, it is also returned for convenience.

#### func  BoolEnsureCap

```go
func BoolEnsureCap(ref *[]bool, capacity int)
```
Calls `BoolSetCap` only if the current `cap(*ref)` is less than the specified
`capacity`.

#### func  BoolEnsureLen

```go
func BoolEnsureLen(ref *[]bool, length int)
```
Calls `BoolSetLen` only if the current `len(*ref)` is less than the specified
`length`.

#### func  BoolEquivalent

```go
func BoolEquivalent(one, two []bool) bool
```
Returns whether `one` and `two` only contain identical values, regardless of
ordering.

#### func  BoolHas

```go
func BoolHas(slice []bool, val bool) bool
```
Returns whether `val` is in `slice`.

#### func  BoolHasAny

```go
func BoolHasAny(slice []bool, vals ...bool) bool
```
Returns whether at least one of the specified `vals` is contained in `slice`.

#### func  BoolRemove

```go
func BoolRemove(ref *[]bool, v bool, all bool)
```
Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if
`all` is `true`.

#### func  BoolSetCap

```go
func BoolSetCap(ref *[]bool, capacity int)
```
Sets `*ref` to a copy of `*ref` with the specified `capacity`.

#### func  BoolSetLen

```go
func BoolSetLen(ref *[]bool, length int)
```
Sets `*ref` to a copy of `*ref` with the specified `length`.

#### func  BoolWithout

```go
func BoolWithout(slice []bool, keepOrder bool, withoutVals ...bool) []bool
```
Removes all specified `withoutVals` from `slice`.

#### func  F64AppendUnique

```go
func F64AppendUnique(ref *[]float64, v float64)
```
Appends `v` to `*ref` only if `*ref` does not already contain `v`.

#### func  F64AppendUniques

```go
func F64AppendUniques(ref *[]float64, vals ...float64)
```
Appends each value in `vals` to `*ref` only `*ref` does not already contain it.

#### func  F64At

```go
func F64At(slice []float64, val float64) int
```
Returns the position of `val` in `slice`.

#### func  F64Convert

```go
func F64Convert(src []interface{}, sparse bool) (dst []float64)
```
Converts `src` to `dst`.

If `sparse` is `true`, then only successfully converted `float64` values are
placed in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src`
in length or indices.

If `sparse` is `false`, `dst` has the same length as `src` and non-convertable
values remain zeroed.

#### func  F64Each

```go
func F64Each(sl []float64, apply ...func(float64) float64) []float64
```
Sets each `float64` in `sl` to the result of passing it to each `apply` func.
Although `sl` is modified in-place, it is also returned for convenience.

#### func  F64EnsureCap

```go
func F64EnsureCap(ref *[]float64, capacity int)
```
Calls `F64SetCap` only if the current `cap(*ref)` is less than the specified
`capacity`.

#### func  F64EnsureLen

```go
func F64EnsureLen(ref *[]float64, length int)
```
Calls `F64SetLen` only if the current `len(*ref)` is less than the specified
`length`.

#### func  F64Equivalent

```go
func F64Equivalent(one, two []float64) bool
```
Returns whether `one` and `two` only contain identical values, regardless of
ordering.

#### func  F64Has

```go
func F64Has(slice []float64, val float64) bool
```
Returns whether `val` is in `slice`.

#### func  F64HasAny

```go
func F64HasAny(slice []float64, vals ...float64) bool
```
Returns whether at least one of the specified `vals` is contained in `slice`.

#### func  F64Remove

```go
func F64Remove(ref *[]float64, v float64, all bool)
```
Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if
`all` is `true`.

#### func  F64SetCap

```go
func F64SetCap(ref *[]float64, capacity int)
```
Sets `*ref` to a copy of `*ref` with the specified `capacity`.

#### func  F64SetLen

```go
func F64SetLen(ref *[]float64, length int)
```
Sets `*ref` to a copy of `*ref` with the specified `length`.

#### func  F64Without

```go
func F64Without(slice []float64, keepOrder bool, withoutVals ...float64) []float64
```
Removes all specified `withoutVals` from `slice`.

#### func  IntAppendUnique

```go
func IntAppendUnique(ref *[]int, v int)
```
Appends `v` to `*ref` only if `*ref` does not already contain `v`.

#### func  IntAppendUniques

```go
func IntAppendUniques(ref *[]int, vals ...int)
```
Appends each value in `vals` to `*ref` only `*ref` does not already contain it.

#### func  IntAt

```go
func IntAt(slice []int, val int) int
```
Returns the position of `val` in `slice`.

#### func  IntConvert

```go
func IntConvert(src []interface{}, sparse bool) (dst []int)
```
Converts `src` to `dst`.

If `sparse` is `true`, then only successfully converted `int` values are placed
in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in
length or indices.

If `sparse` is `false`, `dst` has the same length as `src` and non-convertable
values remain zeroed.

#### func  IntEach

```go
func IntEach(sl []int, apply ...func(int) int) []int
```
Sets each `int` in `sl` to the result of passing it to each `apply` func.
Although `sl` is modified in-place, it is also returned for convenience.

#### func  IntEnsureCap

```go
func IntEnsureCap(ref *[]int, capacity int)
```
Calls `IntSetCap` only if the current `cap(*ref)` is less than the specified
`capacity`.

#### func  IntEnsureLen

```go
func IntEnsureLen(ref *[]int, length int)
```
Calls `IntSetLen` only if the current `len(*ref)` is less than the specified
`length`.

#### func  IntEquivalent

```go
func IntEquivalent(one, two []int) bool
```
Returns whether `one` and `two` only contain identical values, regardless of
ordering.

#### func  IntHas

```go
func IntHas(slice []int, val int) bool
```
Returns whether `val` is in `slice`.

#### func  IntHasAny

```go
func IntHasAny(slice []int, vals ...int) bool
```
Returns whether at least one of the specified `vals` is contained in `slice`.

#### func  IntRemove

```go
func IntRemove(ref *[]int, v int, all bool)
```
Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if
`all` is `true`.

#### func  IntSetCap

```go
func IntSetCap(ref *[]int, capacity int)
```
Sets `*ref` to a copy of `*ref` with the specified `capacity`.

#### func  IntSetLen

```go
func IntSetLen(ref *[]int, length int)
```
Sets `*ref` to a copy of `*ref` with the specified `length`.

#### func  IntWithout

```go
func IntWithout(slice []int, keepOrder bool, withoutVals ...int) []int
```
Removes all specified `withoutVals` from `slice`.

#### func  StrAppendUnique

```go
func StrAppendUnique(ref *[]string, v string)
```
Appends `v` to `*ref` only if `*ref` does not already contain `v`.

#### func  StrAppendUniques

```go
func StrAppendUniques(ref *[]string, vals ...string)
```
Appends each value in `vals` to `*ref` only `*ref` does not already contain it.

#### func  StrAt

```go
func StrAt(slice []string, val string) int
```
Returns the position of `val` in `slice`.

#### func  StrAtIgnoreCase

```go
func StrAtIgnoreCase(vals []string, val string) int
```
Returns the position of lower-case `val` in lower-case `vals`.

#### func  StrConvert

```go
func StrConvert(src []interface{}, sparse bool) (dst []string)
```
Converts `src` to `dst`.

If `sparse` is `true`, then only successfully converted `string` values are
placed in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src`
in length or indices.

If `sparse` is `false`, `dst` has the same length as `src` and non-convertable
values remain zeroed.

#### func  StrEach

```go
func StrEach(sl []string, apply ...func(string) string) []string
```
Sets each `string` in `sl` to the result of passing it to each `apply` func.
Although `sl` is modified in-place, it is also returned for convenience.

#### func  StrEnsureCap

```go
func StrEnsureCap(ref *[]string, capacity int)
```
Calls `StrSetCap` only if the current `cap(*ref)` is less than the specified
`capacity`.

#### func  StrEnsureLen

```go
func StrEnsureLen(ref *[]string, length int)
```
Calls `StrSetLen` only if the current `len(*ref)` is less than the specified
`length`.

#### func  StrEquivalent

```go
func StrEquivalent(one, two []string) bool
```
Returns whether `one` and `two` only contain identical values, regardless of
ordering.

#### func  StrHas

```go
func StrHas(slice []string, val string) bool
```
Returns whether `val` is in `slice`.

#### func  StrHasAny

```go
func StrHasAny(slice []string, vals ...string) bool
```
Returns whether at least one of the specified `vals` is contained in `slice`.

#### func  StrHasIgnoreCase

```go
func StrHasIgnoreCase(vals []string, val string) bool
```
Returns whether lower-case `val` is in lower-case `vals`.

#### func  StrRemove

```go
func StrRemove(ref *[]string, v string, all bool)
```
Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if
`all` is `true`.

#### func  StrSetCap

```go
func StrSetCap(ref *[]string, capacity int)
```
Sets `*ref` to a copy of `*ref` with the specified `capacity`.

#### func  StrSetLen

```go
func StrSetLen(ref *[]string, length int)
```
Sets `*ref` to a copy of `*ref` with the specified `length`.

#### func  StrWithout

```go
func StrWithout(slice []string, keepOrder bool, withoutVals ...string) []string
```
Removes all specified `withoutVals` from `slice`.

--
**godocdown** http://github.com/robertkrimen/godocdown
