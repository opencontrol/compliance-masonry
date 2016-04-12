// +build appengine

package ufs

import (
	"path/filepath"

	"github.com/go-utils/ustr"
)

//	A convenient wrapper around `go-forks/fsnotify.Watcher`.
//
//	**NOTE**: `godocdown` picked `watcher-sandboxed.go` shim instead of `watcher-default.go`:
//	Refer to http://godoc.org/github.com/go-utils/ufs#Watcher for *actual* docs on `Watcher`.
type Watcher struct {
}

//	Returns a new `Watcher`, `err` is always nil.
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{}
	return
}

//	Closes the underlying `me.Watcher`.
func (me *Watcher) Close() (err error) {
	return
}

func (me *Watcher) Go() {
}

func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	if runHandlerNow {
		errs = watchRunHandler(filepath.Clean(dirPath), namePattern, handler)
	}
	return
}
