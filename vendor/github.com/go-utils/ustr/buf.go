package ustr

import (
	"bytes"
	"fmt"
)

//	A convenient wrapper for `bytes.Buffer`.
type Buffer struct {
	bytes.Buffer
}

//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format, args...))`
func (me *Buffer) Write(format string, args ...interface{}) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	me.Buffer.WriteString(format)
}

//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format+"\n", args...))`
func (me *Buffer) Writeln(format string, args ...interface{}) {
	me.Write(format, args...)
	me.Buffer.WriteString("\n")
}
