package uslice

import "strings"

//	Returns the position of lower-case `val` in lower-case `vals`.
func StrAtIgnoreCase(vals []string, val string) int {
	lv := strings.ToLower(val)
	for i, v := range vals {
		if (v == val) || (strings.ToLower(v) == lv) {
			return i
		}
	}
	return -1
}

//	Returns whether lower-case `val` is in lower-case `vals`.
func StrHasIgnoreCase(vals []string, val string) bool {
	return StrAtIgnoreCase(vals, val) >= 0
}

//#begin-gt -gen.gt N:Str T:string

//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
func StrAppendUnique(ref *[]string, v string) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func StrAppendUniques(ref *[]string, vals ...string) {
	for _, v := range vals {
		StrAppendUnique(ref, v)
	}
}

//	Returns the position of `val` in `slice`.
func StrAt(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Converts `src` to `dst`.
//
//	If `sparse` is `true`, then only successfully converted `string` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func StrConvert(src []interface{}, sparse bool) (dst []string) {
	if sparse {
		var (
			val string
			ok  bool
		)
		for _, v := range src {
			if val, ok = v.(string); ok {
				dst = append(dst, val)
			}
		}
	} else {
		dst = make([]string, len(src))
		for i, v := range src {
			dst[i], _ = v.(string)
		}
	}
	return
}

//	Sets each `string` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
func StrEach(sl []string, apply ...func(string) string) []string {
	for _, fn := range apply {
		for i, _ := range sl {
			sl[i] = fn(sl[i])
		}
	}
	return sl
}

//	Calls `StrSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func StrEnsureCap(ref *[]string, capacity int) {
	if cap(*ref) < capacity {
		StrSetCap(ref, capacity)
	}
}

//	Calls `StrSetLen` only if the current `len(*ref)` is less than the specified `length`.
func StrEnsureLen(ref *[]string, length int) {
	if len(*ref) < length {
		StrSetLen(ref, length)
	}
}

//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func StrEquivalent(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if StrAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns whether `val` is in `slice`.
func StrHas(slice []string, val string) bool {
	return StrAt(slice, val) >= 0
}

//	Returns whether at least one of the specified `vals` is contained in `slice`.
func StrHasAny(slice []string, vals ...string) bool {
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
func StrRemove(ref *[]string, v string, all bool) {
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
func StrSetCap(ref *[]string, capacity int) {
	nu := make([]string, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func StrSetLen(ref *[]string, length int) {
	nu := make([]string, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all specified `withoutVals` from `slice`.
func StrWithout(slice []string, keepOrder bool, withoutVals ...string) []string {
	if len(withoutVals) > 0 {
		var pos int
		for _, w := range withoutVals {
			for pos = StrAt(slice, w); pos >= 0; pos = StrAt(slice, w) {
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
