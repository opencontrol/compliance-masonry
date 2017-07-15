package uslice

//#begin-gt -gen.gt N:F64 T:float64

//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
func F64AppendUnique(ref *[]float64, v float64) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func F64AppendUniques(ref *[]float64, vals ...float64) {
	for _, v := range vals {
		F64AppendUnique(ref, v)
	}
}

//	Returns the position of `val` in `slice`.
func F64At(slice []float64, val float64) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Converts `src` to `dst`.
//
//	If `sparse` is `true`, then only successfully converted `float64` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func F64Convert(src []interface{}, sparse bool) (dst []float64) {
	if sparse {
		var (
			val float64
			ok  bool
		)
		for _, v := range src {
			if val, ok = v.(float64); ok {
				dst = append(dst, val)
			}
		}
	} else {
		dst = make([]float64, len(src))
		for i, v := range src {
			dst[i], _ = v.(float64)
		}
	}
	return
}

//	Sets each `float64` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
func F64Each(sl []float64, apply ...func(float64) float64) []float64 {
	for _, fn := range apply {
		for i, _ := range sl {
			sl[i] = fn(sl[i])
		}
	}
	return sl
}

//	Calls `F64SetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func F64EnsureCap(ref *[]float64, capacity int) {
	if cap(*ref) < capacity {
		F64SetCap(ref, capacity)
	}
}

//	Calls `F64SetLen` only if the current `len(*ref)` is less than the specified `length`.
func F64EnsureLen(ref *[]float64, length int) {
	if len(*ref) < length {
		F64SetLen(ref, length)
	}
}

//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func F64Equivalent(one, two []float64) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if F64At(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns whether `val` is in `slice`.
func F64Has(slice []float64, val float64) bool {
	return F64At(slice, val) >= 0
}

//	Returns whether at least one of the specified `vals` is contained in `slice`.
func F64HasAny(slice []float64, vals ...float64) bool {
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
func F64Remove(ref *[]float64, v float64, all bool) {
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
func F64SetCap(ref *[]float64, capacity int) {
	nu := make([]float64, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func F64SetLen(ref *[]float64, length int) {
	nu := make([]float64, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all specified `withoutVals` from `slice`.
func F64Without(slice []float64, keepOrder bool, withoutVals ...float64) []float64 {
	if len(withoutVals) > 0 {
		var pos int
		for _, w := range withoutVals {
			for pos = F64At(slice, w); pos >= 0; pos = F64At(slice, w) {
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
