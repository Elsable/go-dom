// Code generated DO NOT EDIT
// worker.go
package dom

import "syscall/js"

type WorkerIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
	Terminate(args ...interface{})
}
type Worker struct {
	Value
	EventTarget
}

func jsValueToWorker(val js.Value) Worker { return Worker{Value: Value{Value: val}} }
func (v Value) AsWorker() Worker          { return Worker{Value: v} }
func (w Worker) GetOnerror() EventHandler {
	val := w.Get("onerror")
	return jsValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnerror(val EventHandler) {
	w.Set("onerror", val)
}
func (w Worker) GetOnmessage() EventHandler {
	val := w.Get("onmessage")
	return jsValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnmessage(val EventHandler) {
	w.Set("onmessage", val)
}
func (w Worker) GetOnmessageerror() EventHandler {
	val := w.Get("onmessageerror")
	return jsValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnmessageerror(val EventHandler) {
	w.Set("onmessageerror", val)
}
func (w Worker) PostMessage(args ...interface{}) {
	w.Call("postMessage", args...)
}
func (w Worker) Terminate(args ...interface{}) {
	w.Call("terminate", args...)
}
