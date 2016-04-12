package uslice

//#begin-gt -gen.gt N:Bool T:bool

//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
func BoolAppendUnique(ref *[]bool, v bool) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func BoolAppendUniques(ref *[]bool, vals ...bool) {
	for _, v := range vals {
		BoolAppendUnique(ref, v)
	}
}

//	Returns the position of `val` in `slice`.
func BoolAt(slice []bool, val bool) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Converts `src` to `dst`.
//
//	If `sparse` is `true`, then only successfully converted `bool` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func BoolConvert(src []interface{}, sparse bool) (dst []bool) {
	if sparse {
		var (
			val bool
			ok  bool
		)
		for _, v := range src {
			if val, ok = v.(bool); ok {
				dst = append(dst, val)
			}
		}
	} else {
		dst = make([]bool, len(src))
		for i, v := range src {
			dst[i], _ = v.(bool)
		}
	}
	return
}

//	Sets each `bool` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
func BoolEach(sl []bool, apply ...func(bool) bool) []bool {
	for _, fn := range apply {
		for i, _ := range sl {
			sl[i] = fn(sl[i])
		}
	}
	return sl
}

//	Calls `BoolSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func BoolEnsureCap(ref *[]bool, capacity int) {
	if cap(*ref) < capacity {
		BoolSetCap(ref, capacity)
	}
}

//	Calls `BoolSetLen` only if the current `len(*ref)` is less than the specified `length`.
func BoolEnsureLen(ref *[]bool, length int) {
	if len(*ref) < length {
		BoolSetLen(ref, length)
	}
}

//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func BoolEquivalent(one, two []bool) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if BoolAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns whether `val` is in `slice`.
func BoolHas(slice []bool, val bool) bool {
	return BoolAt(slice, val) >= 0
}

//	Returns whether at least one of the specified `vals` is contained in `slice`.
func BoolHasAny(slice []bool, vals ...bool) bool {
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
func BoolRemove(ref *[]bool, v bool, all bool) {
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
func BoolSetCap(ref *[]bool, capacity int) {
	nu := make([]bool, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func BoolSetLen(ref *[]bool, length int) {
	nu := make([]bool, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all specified `withoutVals` from `slice`.
func BoolWithout(slice []bool, keepOrder bool, withoutVals ...bool) []bool {
	if len(withoutVals) > 0 {
		var pos int
		for _, w := range withoutVals {
			for pos = BoolAt(slice, w); pos >= 0; pos = BoolAt(slice, w) {
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
