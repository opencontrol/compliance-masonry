package ustr

import (
	"strings"
)

//	Uses a `Matcher` to determine whether `value` matches any one of the specified simple-`patterns`.
func MatchesAny(value string, patterns ...string) bool {
	var m Matcher
	m.AddPatterns(patterns...)
	return m.IsMatch(value)
}

type matcherPattern struct {
	pattern, prefix, suffix, contains string
	any                               bool
}

//	Matches a string against "simple-patterns": patterns that can have asterisk (*) wildcards only
//	at the beginning ("ends-with"), at the end ("begins-with"), or both ("contains"), or not at all ("equals").
//
//	For more complex pattern-matching needs, go forth and unleash the full force of the standard library's `regexp` package.
//	But I found that in a big portion of pattern-matching use-cases, I'm just doing "begins-or-ends-or-contains-or-equals" testing.
//	Hence the conception of the "simple-pattern".
//
//	There is also an alternative `Pattern` type in this package. Use `Matcher` to match strings against multiple patterns
//	at once, especially if the patterns don't change often and the matchings occur frequently / repeatedly.
//	In simpler, rarer one-off matchings, `Pattern` is preferable for simpler "setup-less" matching.
type Matcher struct {
	patterns     []matcherPattern
	hasWildcards bool
}

//	Adds the specified simple-`patterns` to me.
func (me *Matcher) AddPatterns(patterns ...string) {
	var s string
	patts := make([]matcherPattern, len(patterns))
	for i := 0; i < len(patterns); i++ {
		s = patterns[i]
		if patts[i].pattern, patts[i].any = s, len(s) == 0 || s == "*"; !patts[i].any {
			if strings.HasPrefix(s, "*") && strings.HasSuffix(s, "*") {
				patts[i].contains = s[1 : len(s)-1]
			} else if strings.HasPrefix(s, "*") {
				patts[i].suffix = s[1:]
			} else if strings.HasSuffix(s, "*") {
				patts[i].prefix = s[:len(s)-1]
			}
		}
		if patts[i].any || len(patts[i].contains) > 0 || len(patts[i].prefix) > 0 || len(patts[i].suffix) > 0 {
			me.hasWildcards = true
		}
	}
	me.patterns = append(me.patterns, patts...)
}

//	Returns whether any of the simple-patterns specified for `me` declares a (usable) *-wildcard.
func (me *Matcher) HasWildcardPatterns() bool {
	return me.hasWildcards
}

//	Matches `s` against all patterns in `me`.
func (me *Matcher) IsMatch(s string) bool {
	for i := 0; i < len(me.patterns); i++ {
		if me.patterns[i].any || s == me.patterns[i].pattern {
			return true
		}
		if me.hasWildcards {
			if len(me.patterns[i].prefix) > 0 && strings.HasPrefix(s, me.patterns[i].prefix) {
				return true
			}
			if len(me.patterns[i].suffix) > 0 && strings.HasSuffix(s, me.patterns[i].suffix) {
				return true
			}
			if len(me.patterns[i].contains) > 0 && strings.Contains(s, me.patterns[i].contains) {
				return true
			}
		}
	}
	return false
}

//	An "leaner" alternative to `Matcher` (see docs for `Matcher`). This represents a
//	single "simple-pattern" and provides matching methods for one or multiple values.
type Pattern string

//	Returns whether all specified `values` match this simple-pattern.
func (me Pattern) AllMatch(values ...string) (allMatch bool) {
	allMatch = true
	if len(me) == 0 || me == "*" {
		return
	}
	for _, val := range values {
		if !me.IsMatch(val) {
			allMatch = false
			break
		}
	}
	return
}

//	Returns the first of the specified `values` to match this simple-pattern, or empty if none of them match.
func (me Pattern) AnyMatches(values ...string) (firstMatch string) {
	for _, val := range values {
		if me.IsMatch(val) {
			firstMatch = val
			break
		}
	}
	return
}

//	Returns whether the specified `value` matches this simple-pattern.
func (me Pattern) IsMatch(value string) bool {
	meLen := len(me)
	if meLen == 0 || me == "*" {
		return true
	}
	prefix, suffix := me[0] == '*', me[meLen-1] == '*'
	if prefix && suffix {
		return strings.Contains(value, string(me)[1:meLen-2])
	} else if prefix {
		return strings.HasSuffix(value, string(me)[1:])
	} else if suffix {
		return strings.HasPrefix(value, string(me)[:meLen-1])
	}
	return value == string(me)
}
