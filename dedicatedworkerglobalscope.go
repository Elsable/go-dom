// Code generated DO NOT EDIT
// dedicatedworkerglobalscope.go
package dom

import "syscall/js"

type DedicatedWorkerGlobalScopeIFace interface {
	AddEventListener(args ...interface{})
	CancelAnimationFrame(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	ImportScripts(args ...interface{})
	GetLocation() WorkerLocation
	GetName() string
	GetNavigator() WorkerNavigator
	GetOnerror() OnErrorEventHandler
	SetOnerror(OnErrorEventHandler)
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
	RequestAnimationFrame(args ...interface{}) int
	GetSelf() WorkerGlobalScope
}
type DedicatedWorkerGlobalScope struct {
	Value
	WorkerGlobalScope
	EventTarget
}

func JSValueToDedicatedWorkerGlobalScope(val js.Value) DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: JSValueToValue(val)}
}
func (v Value) AsDedicatedWorkerGlobalScope() DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: v}
}
func NewDedicatedWorkerGlobalScope(args ...interface{}) DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: JSValueToValue(js.Global().Get("DedicatedWorkerGlobalScope").New(args...))}
}
func (d DedicatedWorkerGlobalScope) CancelAnimationFrame(args ...interface{}) {
	d.Call("cancelAnimationFrame", args...)
}
func (d DedicatedWorkerGlobalScope) Close(args ...interface{}) {
	d.Call("close", args...)
}
func (d DedicatedWorkerGlobalScope) GetName() string {
	val := d.Get("name")
	return val.String()
}
func (d DedicatedWorkerGlobalScope) GetOnmessage() EventHandler {
	val := d.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnmessage(val EventHandler) {
	d.Set("onmessage", val)
}
func (d DedicatedWorkerGlobalScope) GetOnmessageerror() EventHandler {
	val := d.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnmessageerror(val EventHandler) {
	d.Set("onmessageerror", val)
}
func (d DedicatedWorkerGlobalScope) PostMessage(args ...interface{}) {
	d.Call("postMessage", args...)
}
func (d DedicatedWorkerGlobalScope) RequestAnimationFrame(args ...interface{}) int {
	val := d.Call("requestAnimationFrame", args...)
	return val.Int()
}
