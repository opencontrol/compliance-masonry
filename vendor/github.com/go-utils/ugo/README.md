# ugo
--
    import "github.com/go-utils/ugo"

Go programming helpers for common miscellaneous needs.

## Usage

```go
var (
	//	The string format used in LogError().
	LogErrorFormat = "%v"

	//	Look-up hash-table for the `OSName` function.
	OSNames = map[string]string{
		"windows":   "Windows",
		"darwin":    "Mac OS X",
		"linux":     "Linux",
		"freebsd":   "FreeBSD",
		"appengine": "Google App Engine",
	}
)
```

#### func  GoPaths

```go
func GoPaths() []string
```
Returns all paths listed in the `GOPATH` environment variable.

#### func  GopathSrc

```go
func GopathSrc(subDirNames ...string) (gps string)
```
Returns the `path/filepath.Join`-ed full directory path for a specified
`$GOPATH/src` sub-directory. Example: `util.GopathSrc("tools", "importers",
"sql")` yields `c:\gd\src\tools\importers\sql` if `$GOPATH` is `c:\gd`.

#### func  GopathSrcGithub

```go
func GopathSrcGithub(gitHubName string, subDirNames ...string) string
```
Returns the `path/filepath.Join`-ed full directory path for a specified
`$GOPATH/src/github.com` sub-directory. Example:
`util.GopathSrcGithub("go-utils", "unum")` yields
`c:\gd\src\github.com\go-utils\unum` if `$GOPATH` is `c:\gd`.

#### func  HostName

```go
func HostName() (hostName string)
```
Returns the result of `os.Hostname` if any, else `localhost`.

#### func  Ifb

```go
func Ifb(cond, ifTrue, ifFalse bool) bool
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifd

```go
func Ifd(cond bool, ifTrue, ifFalse float64) float64
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifi

```go
func Ifi(cond bool, ifTrue, ifFalse int) int
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifi16

```go
func Ifi16(cond bool, ifTrue, ifFalse int16) int16
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifi32

```go
func Ifi32(cond bool, ifTrue, ifFalse int32) int32
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifi64

```go
func Ifi64(cond bool, ifTrue, ifFalse int64) int64
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifs

```go
func Ifs(cond bool, ifTrue string, ifFalse string) string
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifu32

```go
func Ifu32(cond bool, ifTrue, ifFalse uint32) uint32
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifu64

```go
func Ifu64(cond bool, ifTrue, ifFalse uint64) uint64
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifw

```go
func Ifw(cond bool, ifTrue, ifFalse io.Writer) io.Writer
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  Ifx

```go
func Ifx(cond bool, ifTrue, ifFalse interface{}) interface{}
```
Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.

#### func  LogError

```go
func LogError(err error)
```
A convenience short-hand for `log.Println(fmt.Sprintf(LogErrorFormat, err))` if
`err` isn't `nil`.

#### func  MaxProcs

```go
func MaxProcs()
```
Short-hand for: `runtime.GOMAXPROCS(2 * runtime.NumCPU())`.

#### func  OSName

```go
func OSName(goOS string) (name string)
```
Returns the human-readable operating system name represented by the specified
`goOS` name, by looking up the corresponding entry in `OSNames`.

#### func  ParseVersion

```go
func ParseVersion(verstr string) (majorMinor [2]int, both float64)
```
Attempts to extract major and minor version components from a string that begins
with a version number. Example: returns []int{3, 2} and float64(3.2) for a
`verstr` that is `3.2.0 - Build 8.15.10.2761`.

#### func  UserHomeDirPath

```go
func UserHomeDirPath() (dirPath string)
```
Returns the path to the current user's home directory. Might be `C:\Users\Kitty`
under Windows, `/home/Kitty` under Linux or `/Users/Kitty` under Mac OS X.
Specifically, returns the value of either the `%userprofile%` (Windows) or the
`$HOME` (others) environment variable, whichever one is set.

#### type MutexIf

```go
type MutexIf struct {
	sync.Mutex
}
```

A `sync.Mutex` wrapper for convenient conditional `defer`d un/locking.

Example: `defer mut.UnlockIf(mut.LockIf(mycondition))`

#### func (*MutexIf) Lock

```go
func (me *MutexIf) Lock() bool
```

#### func (*MutexIf) LockIf

```go
func (me *MutexIf) LockIf(lock bool) bool
```
Calls `me.Lock` if `lock` is `true`, then returns `lock`.

#### func (*MutexIf) UnlockIf

```go
func (me *MutexIf) UnlockIf(unlock bool)
```
Calls `me.Unlock` if `unlock` is `true`.

--
**godocdown** http://github.com/robertkrimen/godocdown
