# ufs
--
    import "github.com/go-utils/ufs"

Go programming helpers for common file-system needs.

## Usage

```go
var (
	//	The permission bits used in the `EnsureDirExists`, `WriteBinaryFile` and `WriteTextFile` functions.
	ModePerm = os.ModePerm
)
```

#### func  ClearDirectory

```go
func ClearDirectory(dirPath string, keepNamePatterns ...string) (err error)
```
Removes anything in `dirPath` (but not `dirPath` itself), except items whose
`os.FileInfo.Name` matches any of the specified `keepNamePatterns`.

#### func  ClearEmptyDirectories

```go
func ClearEmptyDirectories(dirPath string) (canDelete bool, err error)
```
Removes all directories inside `dirPath`, except those that contain files or
descendent directories that contain files.

#### func  CopyAll

```go
func CopyAll(srcDirPath, dstDirPath string, skipDirs *ustr.Matcher) (err error)
```
Copies all files and directories inside `srcDirPath` to `dstDirPath`. All
sub-directories whose `os.FileInfo.Name` is matched by `skipDirs` (optional) are
skipped.

#### func  CopyFile

```go
func CopyFile(srcFilePath, dstFilePath string) (err error)
```
Performs an `io.Copy` from the specified source file to the specified
destination file.

#### func  DirExists

```go
func DirExists(dirPath string) bool
```
Returns whether a directory (not a file) exists at the specified `dirPath`.

#### func  DirsOrFilesExistIn

```go
func DirsOrFilesExistIn(dirPath string, dirOrFileNames ...string) bool
```
Returns whether all of the specified `dirOrFileNames` exist in `dirPath`.

#### func  EnsureDirExists

```go
func EnsureDirExists(dirPath string) (err error)
```
If a directory does not exist at the specified `dirPath`, attempts to create it.

#### func  ExtractZipFile

```go
func ExtractZipFile(zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ...string) error
```
Extracts a ZIP archive to the local file system. zipFilePath: full file path to
the ZIP archive file. targetDirPath: directory path where un-zipped archive
contents are extracted to. deleteZipFile: deletes the ZIP archive file upon
successful extraction.

#### func  FileExists

```go
func FileExists(filePath string) (fileExists bool)
```
Returns whether a file (not a directory) exists at the specified `filePath`.

#### func  IsNewerThan

```go
func IsNewerThan(srcFilePath, dstFilePath string) (newer bool, err error)
```
Returns whether `srcFilePath` has been modified later than `dstFilePath`.

NOTE: be aware that `newer` will be returned as `true` if `err` is returned as
*not* `nil`, since that is often more convenient for many use-cases.

#### func  MatchesAny

```go
func MatchesAny(name string, patterns ...string) (matchingPattern string, err error)
```
Applies all specified `patterns` to `filepath.Match` and returns the first
successfully matching such pattern.

#### func  ReadBinaryFile

```go
func ReadBinaryFile(filePath string, panicOnError bool) []byte
```
Reads and returns the binary contents of a file with non-idiomatic error
handling, mostly for one-off `package main`s.

#### func  ReadTextFile

```go
func ReadTextFile(filePath string, panicOnError bool, defaultValue string) string
```
Reads and returns the contents of a text file with non-idiomatic error handling,
mostly for one-off `package main`s.

#### func  SaveToFile

```go
func SaveToFile(src io.Reader, dstFilePath string) (err error)
```
Performs an `io.Copy` from the specified `io.Reader` to the specified local
file.

#### func  WalkAllDirs

```go
func WalkAllDirs(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for `dirPath` and all descendent directories (but not files).

#### func  WalkAllFiles

```go
func WalkAllFiles(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all files (but not directories) directly or indirectly
descendent to `dirPath`.

#### func  WalkDirsIn

```go
func WalkDirsIn(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all directories (but not files) in `dirPath`, but not their
sub-directories and not `dirPath` itself.

#### func  WalkFilesIn

```go
func WalkFilesIn(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all files (but not directories) directly inside `dirPath`,
but not for any inside sub-directories.

#### func  WriteBinaryFile

```go
func WriteBinaryFile(filePath string, contents []byte) error
```
A short-hand for `ioutil.WriteFile` using `ModePerm`. Also ensures the target
file's directory exists.

#### func  WriteTextFile

```go
func WriteTextFile(filePath, contents string) error
```
A short-hand for `ioutil.WriteFile`, using `ModePerm`. Also ensures the target
file's directory exists.

#### type DirWalker

```go
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
```

Provides recursive directory walking with a variety of options.

#### func  NewDirWalker

```go
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker)
```
Initializes and returns a new `DirWalker` with the specified (optional)
`WalkerVisitor`s. `deep` sets `VisitSubDirs`.

#### func (*DirWalker) Walk

```go
func (me *DirWalker) Walk(dirPath string) (errs []error)
```
Initiates a walk starting at the specified `dirPath`.

#### type WalkerVisitor

```go
type WalkerVisitor func(fullPath string) (keepWalking bool)
```

Used for `DirWalker.DirVisitor` and `DirWalker.FileVisitor`. Always return
`keepWalking` as true unless you want to immediately terminate a `Walk` early.

#### type Watcher

```go
type Watcher struct {
}
```

A convenient wrapper around `go-forks/fsnotify.Watcher`.

**NOTE**: `godocdown` picked `watcher-sandboxed.go` shim instead of
`watcher-default.go`: Refer to http://godoc.org/github.com/go-utils/ufs#Watcher
for *actual* docs on `Watcher`.

#### func  NewWatcher

```go
func NewWatcher() (me *Watcher, err error)
```
Always returns a new `Watcher`, even if `err` is not `nil` (in which case,
however, `me.Watcher` might be `nil`).

#### func (*Watcher) Close

```go
func (me *Watcher) Close() (err error)
```
Closes the underlying `me.Watcher`.

#### func (*Watcher) Go

```go
func (me *Watcher) Go()
```
Starts watching. A loop designed to be called in a new go-routine, as in `go
myWatcher.Go`. This function returns when `me.Close()` is called.

#### func (*Watcher) WatchIn

```go
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error)
```
Watches dirs/files (whose `filepath.Base` names match the specified
`namePattern`) inside the specified `dirPath` for change event notifications.

`handler` is invoked whenever a change event is observed, providing the full
path.

`runHandlerNow` allows immediate one-off invokation of `handler`. This will
`DirWalker.Walk` the `dirPath`.

An empty `namePattern` is equivalent to `*`.

#### type WatcherHandler

```go
type WatcherHandler func(path string)
```

Handles a file-system notification originating in a `Watcher`.

--
**godocdown** http://github.com/robertkrimen/godocdown
