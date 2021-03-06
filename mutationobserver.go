// Code generated DO NOT EDIT
// mutationobserver.go
package dom

import "syscall/js"

type MutationObserverIFace interface {
	Disconnect(args ...interface{})
	Observe(args ...interface{})
	TakeRecords(args ...interface{})
}
type MutationObserver struct {
	Value
}

func JSValueToMutationObserver(val js.Value) MutationObserver {
	return MutationObserver{Value: JSValueToValue(val)}
}
func (v Value) AsMutationObserver() MutationObserver { return MutationObserver{Value: v} }
func NewMutationObserver(args ...interface{}) MutationObserver {
	return MutationObserver{Value: JSValueToValue(js.Global().Get("MutationObserver").New(args...))}
}
func (m MutationObserver) Disconnect(args ...interface{}) {
	m.Call("disconnect", args...)
}
func (m MutationObserver) Observe(args ...interface{}) {
	m.Call("observe", args...)
}
func (m MutationObserver) TakeRecords(args ...interface{}) {
	m.Call("takeRecords", args...)
}
