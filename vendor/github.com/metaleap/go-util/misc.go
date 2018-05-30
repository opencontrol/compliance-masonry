// Go programming helpers for common miscellaneous needs.
package umisc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	//	The string format used in LogError().
	LogErrorFormat = "%v"
)

func E(msg string) error {
	return errors.New(msg)
}

func F(thing interface{}) float64 {
	f, _ := thing.(float64)
	return f
}

func S(thing interface{}) string {
	s, _ := thing.(string)
	return s
}

func Str(thing interface{}) string {
	return fmt.Sprint(thing)
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfB(cond, ifTrue, ifFalse bool) bool {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfF64(cond bool, ifTrue, ifFalse float64) float64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfI(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfI16(cond bool, ifTrue, ifFalse int16) int16 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfI32(cond bool, ifTrue, ifFalse int32) int32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfI64(cond bool, ifTrue, ifFalse int64) int64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfS(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfU32(cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfU64(cond bool, ifTrue, ifFalse uint64) uint64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfW(cond bool, ifTrue, ifFalse io.Writer) io.Writer {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func IfX(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func JsonDecodeFromFile(fromfilepath string, into interface{}) (err error) {
	var f *os.File
	if f, err = os.Open(fromfilepath); err == nil {
		defer f.Close()
		err = json.NewDecoder(f).Decode(into)
	}
	return
}

func JsonEncodeToFile(from interface{}, tofilepath string) (err error) {
	var f *os.File
	if f, err = os.Create(tofilepath); err == nil {
		defer f.Close()
		err = json.NewEncoder(f).Encode(from)
	}
	return
}

//	A convenience short-hand for `log.Println(fmt.Sprintf(LogErrorFormat, err))` if `err` isn't `nil`.
func LogError(err error) {
	if err != nil {
		log.Println(Strf(LogErrorFormat, err))
	}
}

//	Attempts to extract major and minor version components from a string that begins with a version number.
//	Example: returns []int{3, 2} and float64(3.2) for a `verstr` that is `3.2.0 - Build 8.15.10.2761`.
func ParseVersion(verstr string) (majorMinor [2]int, both float64) {
	var (
		pos, j int
		i      uint64
		err    error
	)
	for _, p := range strings.Split(verstr, ".") {
		if pos = strings.Index(p, " "); pos > 0 {
			p = p[:pos]
		}
		if i, err = strconv.ParseUint(p, 10, 8); err == nil {
			majorMinor[j] = int(i)
			if j++; j >= len(majorMinor) {
				break
			}
		} else {
			break
		}
	}
	if len(majorMinor) > 0 {
		both = float64(majorMinor[0])
	}
	if len(majorMinor) > 1 {
		both += (float64(majorMinor[1]) * 0.1)
	}
	return
}

func Strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
