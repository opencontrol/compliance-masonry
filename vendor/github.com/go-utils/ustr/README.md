# ustr
--
    import "github.com/go-utils/ustr"

Go programming helpers for common string-processing needs.

## Usage

#### func  Concat

```go
func Concat(vals ...string) string
```
Passes the specified `vals` to `strings.Join`.

#### func  ExtractAllIdentifiers

```go
func ExtractAllIdentifiers(src, prefix string) (identifiers []string)
```
Extracts all "identifiers" (as per `ExtractFirstIdentifier`) in `src` and
starting with `prefix` (no duplicates, ordered by occurrence).

#### func  ExtractFirstIdentifier

```go
func ExtractFirstIdentifier(src, prefix string, minPos int) (identifier string)
```
Extracts the first occurrence (at or after `minPos`) of the "identifier"
starting with `prefix` in `src`.

#### func  First

```go
func First(predicate func(s string) bool, step int, vals ...string) string
```
Returns the first `string` in `vals` to match the specified `predicate`.

`step`: 1 to test all values, a higher value to skip n values after each test,
negative for reverse slice traversal, or use 0 to get stuck in an infinite loop.

#### func  FirstNonEmpty

```go
func FirstNonEmpty(vals ...string) (val string)
```
Returns the first non-empty `string` in `vals`.

#### func  Has

```go
func Has(s, substr string) bool
```
Convenience short-hand for `strings.Contains`.

#### func  HasAny

```go
func HasAny(s string, subs ...string) bool
```
Returns whether `s` contains any of the specified sub-strings.

#### func  HasAnyCase

```go
func HasAnyCase(s1, s2 string) bool
```
Returns whether `s1` contains `s2` or lower-case `s1` contains lower-case `s2`.

#### func  HasAnyPrefix

```go
func HasAnyPrefix(s string, prefixes ...string) bool
```
Returns whether `s` starts with any one of the specified `prefixes`.

#### func  HasAnySuffix

```go
func HasAnySuffix(s string, suffixes ...string) bool
```
Returns whether `s` ends with any one of the specified `suffixes`.

#### func  HasOnce

```go
func HasOnce(str1, str2 string) bool
```
Returns whether `str2` is contained in `str1` exactly once.

#### func  Ifm

```go
func Ifm(cond bool, ifTrue, ifFalse map[string]string) map[string]string
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifs

```go
func Ifs(cond bool, ifTrue, ifFalse string) string
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  IndexAny

```go
func IndexAny(s string, seps ...string) (pos int)
```
For all `seps`, records its position of first occurrence in `s`, then returns
the smallest such position.

#### func  IsAscii

```go
func IsAscii(str string) bool
```
Returns whether `str` is ASCII-compatible.

#### func  IsLower

```go
func IsLower(s string) bool
```
Returns whether all `unicode.IsLetter` runes in `s` are lower-case.

#### func  IsOneOf

```go
func IsOneOf(s string, all ...string) bool
```
Returns whether `s` is in `all`.

#### func  IsUpper

```go
func IsUpper(s string) bool
```
Returns whether all `unicode.IsLetter` runes in `s` are upper-case.

#### func  IsUpperAscii

```go
func IsUpperAscii(s string) bool
```

#### func  LettersOnly

```go
func LettersOnly(s string) string
```
Returns a representation of `s` with all non-`unicode.IsLetter` runes removed.

#### func  MatchesAny

```go
func MatchesAny(value string, patterns ...string) bool
```
Uses a `Matcher` to determine whether `value` matches any one of the specified
simple-`patterns`.

#### func  NonEmpties

```go
func NonEmpties(breakAtFirstEmpty bool, vals ...string) (slice []string)
```
Returns a slice that contains the non-empty items in `vals`.

#### func  ParseBool

```go
func ParseBool(s string) bool
```
Returns `strconv.ParseBool` or `false`.

#### func  ParseFloat

```go
func ParseFloat(s string) float64
```
Returns `strconv.ParseFloat` or `0`.

#### func  ParseFloats

```go
func ParseFloats(vals ...string) []float64
```
Returns the parsed `float64`s from `vals` in the same order, or `nil` if one of
them failed to parse.

#### func  ParseInt

```go
func ParseInt(s string) int64
```
Returns `strconv.ParseInt` or `0`.

#### func  ParseUint

```go
func ParseUint(s string) uint64
```
Returns `strconv.ParseUint` or `0`.

#### func  Pluralize

```go
func Pluralize(s string) string
```
A most simplistic (not linguistically-correct) English-language pluralizer that
may be useful for code or doc generation.

If `s` ends with "s", only appends "es": bus -> buses, mess -> messes etc.

If `s` ends with "y" (but not "ay", "ey", "oy", "uy" or "iy"), removes "y" and
appends "ies": autonomy -> autonomies, dictionary -> dictionaries etc.

Otherwise, appends "s": gopher -> gophers, laptop -> laptops etc.

#### func  PrefixWithSep

```go
func PrefixWithSep(prefix, sep, v string) string
```
Prepends `prefix + sep` to `v` only if `prefix` isn't empty.

#### func  PrependIf

```go
func PrependIf(s, p string) string
```
Prepends `p` to `s` only if `s` doesn't already have that prefix.

#### func  ReduceSpaces

```go
func ReduceSpaces(s string) string
```
All occurrences in `s` of multiple subsequent spaces in a row are collapsed into
one single space.

#### func  Replace

```go
func Replace(str string, repls map[string]string) string
```
Replaces in `str` all occurrences of all `repls` hash-map keys with their
respective associated (mapped) value.

#### func  SafeIdentifier

```go
func SafeIdentifier(s string) string
```
Creates a Pascal-cased "identifier" version of the specified string.

#### func  Split

```go
func Split(v, s string) (sl []string)
```
Returns an empty slice is `v` is emtpy, otherwise like `strings.Split`

#### func  StripPrefix

```go
func StripPrefix(val, prefix string) string
```
Strips `prefix` off `val` if possible.

#### func  StripSuffix

```go
func StripSuffix(val, suffix string) string
```
Strips `suffix` off `val` if possible.

#### func  ToLowerIfUpper

```go
func ToLowerIfUpper(s string) string
```
Returns the lower-case representation of `s` only if it is currently fully
upper-case as per `IsUpper`.

#### func  ToUpperIfLower

```go
func ToUpperIfLower(s string) string
```
Returns the upper-case representation of `s` only if it is currently fully
lower-case as per `IsLower`.

#### type Buffer

```go
type Buffer struct {
	bytes.Buffer
}
```

A convenient wrapper for `bytes.Buffer`.

#### func (*Buffer) Write

```go
func (me *Buffer) Write(format string, args ...interface{})
```
Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format,
args...))`

#### func (*Buffer) Writeln

```go
func (me *Buffer) Writeln(format string, args ...interface{})
```
Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format+"\n",
args...))`

#### type Matcher

```go
type Matcher struct {
}
```

Matches a string against "simple-patterns": patterns that can have asterisk (*)
wildcards only at the beginning ("ends-with"), at the end ("begins-with"), or
both ("contains"), or not at all ("equals").

For more complex pattern-matching needs, go forth and unleash the full force of
the standard library's `regexp` package. But I found that in a big portion of
pattern-matching use-cases, I'm just doing
"begins-or-ends-or-contains-or-equals" testing. Hence the conception of the
"simple-pattern".

There is also an alternative `Pattern` type in this package. Use `Matcher` to
match strings against multiple patterns at once, especially if the patterns
don't change often and the matchings occur frequently / repeatedly. In simpler,
rarer one-off matchings, `Pattern` is preferable for simpler "setup-less"
matching.

#### func (*Matcher) AddPatterns

```go
func (me *Matcher) AddPatterns(patterns ...string)
```
Adds the specified simple-`patterns` to me.

#### func (*Matcher) HasWildcardPatterns

```go
func (me *Matcher) HasWildcardPatterns() bool
```
Returns whether any of the simple-patterns specified for `me` declares a
(usable) *-wildcard.

#### func (*Matcher) IsMatch

```go
func (me *Matcher) IsMatch(s string) bool
```
Matches `s` against all patterns in `me`.

#### type Pattern

```go
type Pattern string
```

An "leaner" alternative to `Matcher` (see docs for `Matcher`). This represents a
single "simple-pattern" and provides matching methods for one or multiple
values.

#### func (Pattern) AllMatch

```go
func (me Pattern) AllMatch(values ...string) (allMatch bool)
```
Returns whether all specified `values` match this simple-pattern.

#### func (Pattern) AnyMatches

```go
func (me Pattern) AnyMatches(values ...string) (firstMatch string)
```
Returns the first of the specified `values` to match this simple-pattern, or
empty if none of them match.

#### func (Pattern) IsMatch

```go
func (me Pattern) IsMatch(value string) bool
```
Returns whether the specified `value` matches this simple-pattern.

--
**godocdown** http://github.com/robertkrimen/godocdown
