package uslice

//#begin-gt -gen.gt N:Int T:int

//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
func IntAppendUnique(ref *[]int, v int) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func IntAppendUniques(ref *[]int, vals ...int) {
	for _, v := range vals {
		IntAppendUnique(ref, v)
	}
}

//	Returns the position of `val` in `slice`.
func IntAt(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Converts `src` to `dst`.
//
//	If `sparse` is `true`, then only successfully converted `int` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func IntConvert(src []interface{}, sparse bool) (dst []int) {
	if sparse {
		var (
			val int
			ok  bool
		)
		for _, v := range src {
			if val, ok = v.(int); ok {
				dst = append(dst, val)
			}
		}
	} else {
		dst = make([]int, len(src))
		for i, v := range src {
			dst[i], _ = v.(int)
		}
	}
	return
}

//	Sets each `int` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
func IntEach(sl []int, apply ...func(int) int) []int {
	for _, fn := range apply {
		for i, _ := range sl {
			sl[i] = fn(sl[i])
		}
	}
	return sl
}

//	Calls `IntSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func IntEnsureCap(ref *[]int, capacity int) {
	if cap(*ref) < capacity {
		IntSetCap(ref, capacity)
	}
}

//	Calls `IntSetLen` only if the current `len(*ref)` is less than the specified `length`.
func IntEnsureLen(ref *[]int, length int) {
	if len(*ref) < length {
		IntSetLen(ref, length)
	}
}

//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func IntEquivalent(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if IntAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns whether `val` is in `slice`.
func IntHas(slice []int, val int) bool {
	return IntAt(slice, val) >= 0
}

//	Returns whether at least one of the specified `vals` is contained in `slice`.
func IntHasAny(slice []int, vals ...int) bool {
	for _, v1 := range vals {
		for _, v2 := range slice {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

//	Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if `all` is `true`.
func IntRemove(ref *[]int, v int, all bool) {
	for i := 0; i < len(*ref); i++ {
		if (*ref)[i] == v {
			before, after := (*ref)[:i], (*ref)[i+1:]
			*ref = append(before, after...)
			if !all {
				break
			}
		}
	}
}

//	Sets `*ref` to a copy of `*ref` with the specified `capacity`.
func IntSetCap(ref *[]int, capacity int) {
	nu := make([]int, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func IntSetLen(ref *[]int, length int) {
	nu := make([]int, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all specified `withoutVals` from `slice`.
func IntWithout(slice []int, keepOrder bool, withoutVals ...int) []int {
	if len(withoutVals) > 0 {
		var pos int
		for _, w := range withoutVals {
			for pos = IntAt(slice, w); pos >= 0; pos = IntAt(slice, w) {
				if keepOrder {
					slice = append(slice[:pos], slice[pos+1:]...)
				} else {
					slice[pos] = slice[len(slice)-1]
					slice = slice[:len(slice)-1]
				}
			}
		}
	}
	return slice
}

//#end-gt
