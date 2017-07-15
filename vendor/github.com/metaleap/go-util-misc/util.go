package ugo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)


type CmdTry struct {
	Args	[]string
	Ran		*bool
}


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

var (
	_userHomeDirPath string
	_userDataDirPath string
)

func SetupJsonProtoPipes (bufferCapacity int, clenProto bool, needJsonOut bool) (stdin *bufio.Scanner, rawOut *bufio.Writer, jsonOut *json.Encoder) {
	stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufferCapacity), bufferCapacity)
	if clenProto {
		stdin.Split(func(data []byte, ateof bool) (advance int, token []byte, err error) {
			if i_cl1 := bytes.Index(data, []byte("Content-Length: "))  ;  i_cl1>=0 {
				datafromclen := data[i_cl1+16:]  ;  if i_cl2 := bytes.IndexAny(datafromclen, "\r\n")  ;  i_cl2>0 {
					if clen,e := strconv.Atoi(string(datafromclen[:i_cl2]))  ;  e!=nil { err = e } else {
						if i_js1 := bytes.Index(datafromclen, []byte("{\""))  ;  i_js1 > i_cl2 {
							if i_js2 := i_js1+clen  ;  len(datafromclen)>=i_js2 {
								advance = i_cl1 + 16 + i_js2  ;  token = datafromclen[i_js1:i_js2]
							}
						}
					}
				}
			}
			return
		})
	}
	rawOut = bufio.NewWriterSize(os.Stdout, bufferCapacity)
	if needJsonOut {
		jsonOut = json.NewEncoder(rawOut)
		jsonOut.SetEscapeHTML(false)
		jsonOut.SetIndent("","")
	}
	return
}

func CmdTryStart (cmdname string, cmdargs ...string) (err error) {
	cmd := exec.Command(cmdname, cmdargs...)
	err = cmd.Start()  ;  defer cmd.Wait()
	if cmd.Process != nil { cmd.Process.Kill() }
	return
}


func CmdsTryStart (cmds map[string]*CmdTry) {
	var w sync.WaitGroup
	run := func (cmd string, try *CmdTry) {
		defer w.Done()  ;  *try.Ran = nil==CmdTryStart(cmd, try.Args...)
	}
	for cmdname,cmdmore := range cmds {  w.Add(1)  ;  go run(cmdname, cmdmore)  }
	w.Wait()
}


func CmdExecStdin (stdin string, dir string, cmdname string, cmdargs ...string) (stdout string, stderr string, err error) {
	if len(cmdname)>0 && strings.Contains(cmdname, " ") && len(cmdargs)==0 {
		cmdargs = strings.Split(cmdname, " ")
		cmdname = cmdargs[0]
		cmdargs = cmdargs[1:]
	}
	cmd := exec.Command(cmdname, cmdargs...)
	cmd.Dir = dir
	if len(stdin)>0 { cmd.Stdin = strings.NewReader(stdin) }
	var bufout, buferr bytes.Buffer
	cmd.Stdout = &bufout
	cmd.Stderr = &buferr
	if err = cmd.Run() ; err != nil {
		if _, isexiterr := err.(*exec.ExitError) ; isexiterr || strings.Contains(err.Error(), "pipe has been ended") {
			err = nil
		}
	}
	stdout = bufout.String()
	stderr = strings.TrimSpace(buferr.String())
	return
}

func CmdExecIn (dir string, cmdname string, cmdargs ...string) (out string, err error) {
	var output []byte
	cmd := exec.Command(cmdname, cmdargs...)
	cmd.Dir = dir
	output,err = cmd.CombinedOutput() // wish Output() would suffice, but sadly some tools abuse stderr for all sorts of non-error 'metainfotainment' (hi godoc & gofmt!)
	out = strings.TrimSpace(string(output)) // do this regardless of err, because it might well be benign such as "exitcode 2", in which case output is still wanted
	return
}

func CmdExecInOr (def string, dir string, cmdname string, cmdargs ...string) string {
	out,err := CmdExecIn(dir, cmdname, cmdargs...)
	if err != nil { return def }
	return out
}

func CmdExec (cmdname string, cmdargs ...string) (string, error) {
	return CmdExecIn("", cmdname, cmdargs...)
}

func CmdExecOr (def string, cmdname string, cmdargs ...string) string {
	return CmdExecInOr(def, "", cmdname, cmdargs...)
}

func WaitOn (funcs ...func()) {
	if l := len(funcs) ; l==0 { return } else if l==1 { funcs[0]() ; return }
	var wait sync.WaitGroup
	run := func(fn func()) {
		defer wait.Done()
		fn()
	}
	for _,fn := range funcs {
		wait.Add(1)
		go run(fn)
	}
	wait.Wait()
}

func WaitOn_ (funcs ...func()) {
	for _,fn := range funcs { fn() }
}

func E (msg string) error {
	return errors.New(msg)
}

func F (thing interface{}) float64 {
	f,_ := thing.(float64)  ;  return f
}

func S (thing interface{}) string {
	s,_ := thing.(string)  ;  return s
}

func SPr (thing interface{}) string {
	return fmt.Sprint(thing)
}



//	A `sync.Mutex` wrapper for convenient conditional `defer`d un/locking.
//
//	Example: `defer mut.UnlockIf(mut.LockIf(mycondition))`
type MutexIf struct {
	sync.Mutex
}

func (me *MutexIf) Lock() bool {
	me.Mutex.Lock()
	return true
}

//	Calls `me.Lock` if `lock` is `true`, then returns `lock`.
func (me *MutexIf) LockIf(lock bool) bool {
	if lock {
		me.Mutex.Lock()
	}
	return lock
}

//	Calls `me.Unlock` if `unlock` is `true`.
func (me *MutexIf) UnlockIf(unlock bool) {
	if unlock {
		me.Mutex.Unlock()
	}
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

//	Returns all paths listed in the `GOPATH` environment variable.
func GoPaths() []string {
	return filepath.SplitList(os.Getenv("GOPATH"))
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src` sub-directory.
//	Example: `util.GopathSrc("tools", "importers", "sql")` yields `c:\gd\src\tools\importers\sql` if `$GOPATH` is `c:\gd`.
func GopathSrc(subDirNames ...string) (gps string) {
	gp := []string{"", "src"}
	for _, goPath := range GoPaths() { // in 99% of setups there's only 1 GOPATH, but hey..
		gp[0] = goPath
		if gps = filepath.Join(append(gp, subDirNames...)...); dirExists(gps) {
			break
		}
	}
	return
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src/github.com` sub-directory.
//	Example: `util.GopathSrcGithub("go-utils", "unum")` yields `c:\gd\src\github.com\go-utils\unum` if `$GOPATH` is `c:\gd`.
func GopathSrcGithub(gitHubName string, subDirNames ...string) string {
	return GopathSrc(append([]string{"github.com", gitHubName}, subDirNames...)...)
}

//	Returns the result of `os.Hostname` if any, else `localhost`.
func HostName() (hostName string) {
	if hostName, _ = os.Hostname(); len(hostName) == 0 {
		hostName = "localhost"
	}
	return
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifb(cond, ifTrue, ifFalse bool) bool {
	return (cond && ifTrue) || ((!cond) && ifFalse)
	// if cond {
	// 	return ifTrue
	// }
	// return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifd(cond bool, ifTrue, ifFalse float64) float64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifi(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifi16(cond bool, ifTrue, ifFalse int16) int16 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifi32(cond bool, ifTrue, ifFalse int32) int32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifi64(cond bool, ifTrue, ifFalse int64) int64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifs(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifu32(cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifu64(cond bool, ifTrue, ifFalse uint64) uint64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifw(cond bool, ifTrue, ifFalse io.Writer) io.Writer {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifx(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	A convenience short-hand for `log.Println(fmt.Sprintf(LogErrorFormat, err))` if `err` isn't `nil`.
func LogError(err error) {
	if err != nil {
		log.Println(strf(LogErrorFormat, err))
	}
}

//	Short-hand for: `runtime.GOMAXPROCS(2 * runtime.NumCPU())`.
func MaxProcs() {
	runtime.GOMAXPROCS(2 * runtime.NumCPU())
}

//	Returns the human-readable operating system name represented by the specified
//	`goOS` name, by looking up the corresponding entry in `OSNames`.
func OSName(goOS string) (name string) {
	if name = OSNames[goOS]; len(name) == 0 {
		name = strings.ToTitle(goOS)
	}
	return
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

func strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func UserDataDirPath () string {
	dirpath := _userDataDirPath
	if len(dirpath) == 0 {
		probeenvvars := []string { "XDG_CACHE_HOME", "XDG_CONFIG_HOME", "LOCALAPPDATA", "APPDATA" }
		for _,envvar := range probeenvvars {
			if maybedirpath := os.Getenv(envvar) ; len(maybedirpath)>0 && dirExists(maybedirpath) {
				dirpath = maybedirpath
				break
			}
		}
		if len(dirpath) == 0 {
			probehomesubdirs := []string { ".cache", ".config", "Library/Caches", "Library/Application Support" }
			for _,homesubdir := range probehomesubdirs {
				if maybedirpath := filepath.Join(UserHomeDirPath(),homesubdir) ; dirExists(maybedirpath) {
					dirpath = maybedirpath
					break
				}
			}
			if len(dirpath) == 0 {
				dirpath = UserHomeDirPath()
			}
		}
		_userDataDirPath = dirpath
	}
	return dirpath
}

//	Returns the path to the current user's home directory.
func UserHomeDirPath () string {
	dirpath := _userHomeDirPath
	if len(dirpath) == 0 {
		if user,err := user.Current() ; err == nil && len(user.HomeDir) > 0 && dirExists(user.HomeDir) {
			dirpath = user.HomeDir
		} else if dirpath = os.Getenv("USERPROFILE") ; len(dirpath) == 0 || !dirExists(dirpath) {
			dirpath = os.Getenv("HOME")
		}
		_userHomeDirPath = dirpath
	}
	return dirpath
}
