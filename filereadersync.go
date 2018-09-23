// Code generated DO NOT EDIT
// filereadersync.go
package dom

import "syscall/js"

type FileReaderSyncIFace interface {
	ReadAsArrayBuffer(args ...interface{}) ArrayBuffer
	ReadAsBinaryString(args ...interface{}) string
	ReadAsDataURL(args ...interface{}) string
	ReadAsText(args ...interface{}) string
}
type FileReaderSync struct {
	Value
}

func jsValueToFileReaderSync(val js.Value) FileReaderSync {
	return FileReaderSync{Value: Value{Value: val}}
}
func (v Value) AsFileReaderSync() FileReaderSync { return FileReaderSync{Value: v} }
func (f FileReaderSync) ReadAsArrayBuffer(args ...interface{}) ArrayBuffer {
	val := f.Call("readAsArrayBuffer", args...)
	return jsValueToArrayBuffer(val.JSValue())
}
func (f FileReaderSync) ReadAsBinaryString(args ...interface{}) string {
	val := f.Call("readAsBinaryString", args...)
	return val.String()
}
func (f FileReaderSync) ReadAsDataURL(args ...interface{}) string {
	val := f.Call("readAsDataURL", args...)
	return val.String()
}
func (f FileReaderSync) ReadAsText(args ...interface{}) string {
	val := f.Call("readAsText", args...)
	return val.String()
}
