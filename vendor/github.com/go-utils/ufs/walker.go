package ufs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//	Used for `DirWalker.DirVisitor` and `DirWalker.FileVisitor`.
//	Always return `keepWalking` as true unless you want to immediately terminate a `Walk` early.
type WalkerVisitor func(fullPath string) (keepWalking bool)

//	An empty `WalkerVisitor` used in place of a `nil` directory or file visitor during a `DirWalker.Walk`. Always returns `true`.
func walkerVisitorNoop(_ string) bool {
	return true
}

//	Provides recursive directory walking with a variety of options.
type DirWalker struct {
	//	`Walk` returns a slice of all `error`s encountered but keeps walking as indicated by
	//	`DirVisitor` and/or `FileVisitor` --- to abort walking upon the first `error`, set this to `true`.
	BreakOnError bool

	//	After invoking `DirVisitor` on the specified directory (if `VisitSelf`), by default
	//	its files get visited first before visiting its sub-directories.
	//	If `VisitDirsFirst` is `true`, then files get visited last, after
	//	having visited all sub-directories.
	VisitDirsFirst bool

	//	If `false`, only the items in the specified directory get visited
	//	(and the directory itself if `VisitSelf`), but no items inside its sub-directories.
	VisitSubDirs bool

	//	Defaults to `true` if initialized via `NewDirWalker`.
	VisitSelf bool

	//	Called for every directory being visited during a `Walk`.
	DirVisitor WalkerVisitor

	//	Called for every file being visited during a `Walk`.
	FileVisitor WalkerVisitor
}

//	Initializes and returns a new `DirWalker` with the specified (optional) `WalkerVisitor`s.
//	`deep` sets `VisitSubDirs`.
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: deep, VisitSelf: true}
	return
}

//	Initiates a walk starting at the specified `dirPath`.
func (me *DirWalker) Walk(dirPath string) (errs []error) {
	me.walk(me.VisitSelf, dirPath, &errs)
	return
}

func (me *DirWalker) walk(walkSelf bool, dirPath string, errs *[]error) {
	dirVisitor, fileVisitor := me.DirVisitor, me.FileVisitor
	if dirVisitor == nil {
		dirVisitor = walkerVisitorNoop
	}
	if fileVisitor == nil {
		fileVisitor = walkerVisitorNoop
	}
	if walkSelf {
		walkSelf = dirVisitor(dirPath)
	} else {
		walkSelf = true
	}
	if walkSelf {
		if fileInfos, err := ioutil.ReadDir(dirPath); err == nil {
			if me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
				}
			}
			if !me.walkInfos(dirPath, fileInfos, false, fileVisitor, errs) {
				return
			}
			if !me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
				}
			}
		} else if *errs = append(*errs, err); me.BreakOnError {
			return
		}
	}
}

func (me *DirWalker) walkInfos(dirPath string, fileInfos []os.FileInfo, isDir bool, visitor WalkerVisitor, errs *[]error) (keepWalking bool) {
	var fullPath string
	keepWalking = true
	for _, fi := range fileInfos {
		if fullPath = filepath.Join(dirPath, fi.Name()); fi.IsDir() == isDir {
			if keepWalking = visitor(fullPath); !keepWalking {
				break
			} else if isDir && me.VisitSubDirs {
				me.walk(false, fullPath, errs)
			}
		}
	}
	return
}
