package ustr

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/go-utils/ugo"
	"github.com/go-utils/uslice"
)

/*
//	Sets all values in m to the empty string.
func ClearMap(m map[string]string) {
	for k, v := range m {
		if len(v) > 0 {
			m[k] = v[:0]
		}
	}
}
*/

//	Passes the specified `vals` to `strings.Join`.
func Concat(vals ...string) string {
	return strings.Join(vals, "")
}

/*
//	A simple string-similarity algorithm.
func Distance(s1, s2 string) int {
	var (
		cost, min1, min2, min3, i, j int
		d                            = make([][]int, len(s1)+1)
	)
	for i = 0; i < len(d); i++ {
		d[i] = make([]int, len(s2)+1)
		d[i][0] = i
	}
	for i = 0; i < len(d[0]); i++ {
		d[0][i] = i
	}
	for i = 1; i < len(d); i++ {
		for j = 1; j < len(d[0]); j++ {
			cost = ugo.Ifi(s1[i-1] == s2[j-1], 0, 1)
			min1 = d[i-1][j] + 1
			min2 = d[i][j-1] + 1
			min3 = d[i-1][j-1] + cost
			d[i][j] = int(math.Min(math.Min(float64(min1), float64(min2)), float64(min3)))
		}
	}
	return d[len(s1)][len(s2)]
}
*/

//	Extracts all "identifiers" (as per `ExtractFirstIdentifier`) in `src` and starting with `prefix` (no duplicates, ordered by occurrence).
func ExtractAllIdentifiers(src, prefix string) (identifiers []string) {
	minPos := 0
	id := ExtractFirstIdentifier(src, prefix, minPos)
	for len(id) > 0 {
		if minPos = strings.Index(src, id) + 1; !uslice.StrHas(identifiers, id) {
			identifiers = append(identifiers, id)
		}
		id = ExtractFirstIdentifier(src, prefix, minPos)
	}
	return
}

//	Extracts the first occurrence (at or after `minPos`) of the "identifier" starting with `prefix` in `src`.
func ExtractFirstIdentifier(src, prefix string, minPos int) (identifier string) {
	sub := src[minPos:]
	pos := strings.Index(sub, prefix)
	if pos >= 0 {
		for i, r := range sub[pos:] {
			if !(unicode.IsNumber(r) || unicode.IsLetter(r) || r == '_') {
				identifier = sub[pos : pos+i]
				break
			}
		}
	}
	return
}

//	Returns the first `string` in `vals` to match the specified `predicate`.
//
//	`step`: 1 to test all values, a higher value to skip n values after each test, negative for reverse slice traversal, or use 0 to get stuck in an infinite loop.
func First(predicate func(s string) bool, step int, vals ...string) string {
	l := len(vals)
	reverse := step < 0
	for i := ugo.Ifi(reverse, l-1, 0); ugo.Ifb(reverse, i >= 0, i < l); i += step {
		if predicate(vals[i]) {
			return vals[i]
		}
	}
	return ""
}

//	Returns the first non-empty `string` in `vals`.
func FirstNonEmpty(vals ...string) (val string) {
	// return First(func(s string) bool { return len(s) > 0 }, step, vals...)
	for _, val = range vals {
		if len(val) > 0 {
			return
		}
	}
	return
}

//	Convenience short-hand for `strings.Contains`.
func Has(s, substr string) bool {
	return strings.Contains(s, substr)
}

//	Returns whether `s` contains any of the specified sub-strings.
func HasAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

//	Returns whether `s1` contains `s2` or lower-case `s1` contains lower-case `s2`.
func HasAnyCase(s1, s2 string) bool {
	return strings.Contains(s1, s2) || strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
}

//	Returns whether `s` starts with any one of the specified `prefixes`.
func HasAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

//	Returns whether `s` ends with any one of the specified `suffixes`.
func HasAnySuffix(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

//	Returns whether `str2` is contained in `str1` exactly once.
func HasOnce(str1, str2 string) bool {
	first, last := strings.Index(str1, str2), strings.LastIndex(str1, str2)
	return (first >= 0) && (first == last)
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifm(cond bool, ifTrue, ifFalse map[string]string) map[string]string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifs(cond bool, ifTrue, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	For all `seps`, records its position of first occurrence in `s`, then returns the smallest such position.
func IndexAny(s string, seps ...string) (pos int) {
	pos = -1
	for index, sep := range seps {
		if index = strings.Index(s, sep); pos < 0 || (index >= 0 && index < pos) {
			pos = index
		}
	}
	return
}

//	Returns whether `str` is ASCII-compatible.
func IsAscii(str string) bool {
	for _, c := range str {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

//	Returns whether all `unicode.IsLetter` runes in `s` are lower-case.
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

//	Returns whether `s` is in `all`.
func IsOneOf(s string, all ...string) bool {
	for _, a := range all {
		if s == a {
			return true
		}
	}
	return false
}

//	Returns whether all `unicode.IsLetter` runes in `s` are upper-case.
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func IsUpperAscii(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII || (unicode.IsLetter(r) && !unicode.IsUpper(r)) {
			return false
		}
	}
	return true
}

//	Returns a representation of `s` with all non-`unicode.IsLetter` runes removed.
func LettersOnly(s string) string {
	var buf Buffer
	for _, r := range s {
		if unicode.IsLetter(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

//	Returns a slice that contains the non-empty items in `vals`.
func NonEmpties(breakAtFirstEmpty bool, vals ...string) (slice []string) {
	for _, s := range vals {
		if len(s) > 0 {
			slice = append(slice, s)
		} else if breakAtFirstEmpty {
			break
		}
	}
	return
}

//	Returns `strconv.ParseBool` or `false`.
func ParseBool(s string) bool {
	v, _ := strconv.ParseBool(s)
	return v
}

//	Returns `strconv.ParseFloat` or `0`.
func ParseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

//	Returns the parsed `float64`s from `vals` in the same order, or `nil` if one of them failed to parse.
func ParseFloats(vals ...string) []float64 {
	var (
		f   float64
		err error
	)
	all := make([]float64, 0, len(vals))
	for _, s := range vals {
		if f, err = strconv.ParseFloat(s, 64); err == nil {
			all = append(all, f)
		} else {
			return nil
		}
	}
	return all
}

//	Returns `strconv.ParseInt` or `0`.
func ParseInt(s string) int64 {
	v, _ := strconv.ParseInt(s, 0, 64)
	return v
}

//	Returns `strconv.ParseUint` or `0`.
func ParseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 0, 64)
	return v
}

//	A most simplistic (not linguistically-correct) English-language pluralizer that may be useful for code or doc generation.
//
//	If `s` ends with "s", only appends "es": bus -> buses, mess -> messes etc.
//
//	If `s` ends with "y" (but not "ay", "ey", "oy", "uy" or "iy"), removes "y" and appends "ies": autonomy -> autonomies, dictionary -> dictionaries etc.
//
//	Otherwise, appends "s": gopher -> gophers, laptop -> laptops etc.
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	if (len(s) > 1) && strings.HasSuffix(s, "y") && !IsOneOf(s[(len(s)-2):], "ay", "ey", "oy", "uy", "iy") {
		return s[0:(len(s)-1)] + "ies"
	}
	return s + "s"
}

//	Prepends `prefix + sep` to `v` only if `prefix` isn't empty.
func PrefixWithSep(prefix, sep, v string) string {
	if len(prefix) > 0 {
		return prefix + sep + v
	}
	return v
}

//	Prepends `p` to `s` only if `s` doesn't already have that prefix.
func PrependIf(s, p string) string {
	if strings.HasPrefix(s, p) {
		return s
	}
	return p + s
}

//	All occurrences in `s` of multiple subsequent spaces in a row are collapsed into one single space.
func ReduceSpaces(s string) string {
	for strings.Index(s, "  ") >= 0 {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

//	Replaces in `str` all occurrences of all `repls` hash-map keys with their respective associated (mapped) value.
func Replace(str string, repls map[string]string) string {
	for k, v := range repls {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

//	Creates a Pascal-cased "identifier" version of the specified string.
func SafeIdentifier(s string) string {
	var (
		isL, isD, last bool
		buf            Buffer
	)
	for i, r := range s {
		if isL, isD = unicode.IsLetter(r), unicode.IsDigit(r); isL || isD || ((r == '_') && (i == 0)) {
			if (i > 0) && (isL != last) {
				buf.WriteRune(' ')
			}
			buf.WriteRune(r)
		} else {
			buf.WriteRune(' ')
		}
		last = isL
	}
	words := Split(strings.Title(buf.String()), " ")
	for i, w := range words {
		if (len(w) > 1) && IsUpper(w) {
			words[i] = strings.Title(strings.ToLower(w))
		}
	}
	return strings.Join(words, "")
}

//	Returns an empty slice is `v` is emtpy, otherwise like `strings.Split`
func Split(v, s string) (sl []string) {
	if len(v) > 0 {
		sl = strings.Split(v, s)
	}
	return
}

//	Strips `prefix` off `val` if possible.
func StripPrefix(val, prefix string) string {
	for strings.HasPrefix(val, prefix) {
		val = val[len(prefix):]
	}
	return val
}

//	Strips `suffix` off `val` if possible.
func StripSuffix(val, suffix string) string {
	for strings.HasSuffix(val, suffix) {
		val = val[:len(val)-len(suffix)]
	}
	return val
}

/*
func ToFloat32 (str string) float32 {
	var f, err = strconv.ParseFloat(str, 32)
	if err == nil { return float32(f) }
	return 0.0
}

func ToFloat64 (str string) float64 {
	var f, err = strconv.ParseFloat(str, 64)
	if err == nil { return f }
	return 0.0
}

func ToFloat64s (strs ... string) []float64 {
	var f = make([]float64, len(strs))
	for i, s := range strs { f[i] = ToFloat64(s) }
	return f
}

func ToInt (str string) int {
	var i, err = strconv.Atoi(str)
	if err == nil { return i }
	return 0
}

func ToString (any interface{}, nilVal string) string {
	if any == nil {
		return nilVal
	}
	if s, isS := any.(string); isS {
		return s
	}
	if f, isF := any.(fmt.Stringer); isF {
		return f.String()
	}
	return fmt.Sprintf("%v", any)
}

func ToStrings (any interface{}) []string {
	if sl, isSl := any.([]string); isSl {
		return sl
	}
	return nil
}
*/

//	Returns the lower-case representation of `s` only if it is currently fully upper-case as per `IsUpper`.
func ToLowerIfUpper(s string) string {
	if IsUpper(s) {
		return strings.ToLower(s)
	}
	return s
}

//	Returns the upper-case representation of `s` only if it is currently fully lower-case as per `IsLower`.
func ToUpperIfLower(s string) string {
	if IsLower(s) {
		return strings.ToUpper(s)
	}
	return s
}
