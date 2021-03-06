// Code generated DO NOT EDIT
// promiserejectionevent.go
package dom

import "syscall/js"

type PromiseRejectionEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetEventPhase() int
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	PreventDefault(args ...interface{})
	GetPromise()
	GetReason() Value
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type PromiseRejectionEvent struct {
	Value
	Event
}

func JSValueToPromiseRejectionEvent(val js.Value) PromiseRejectionEvent {
	return PromiseRejectionEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPromiseRejectionEvent() PromiseRejectionEvent { return PromiseRejectionEvent{Value: v} }
func NewPromiseRejectionEvent(args ...interface{}) PromiseRejectionEvent {
	return PromiseRejectionEvent{Value: JSValueToValue(js.Global().Get("PromiseRejectionEvent").New(args...))}
}
func (p PromiseRejectionEvent) GetPromise() Value {
	val := p.Get("promise")
	return val
}
func (p PromiseRejectionEvent) GetReason() Value {
	val := p.Get("reason")
	return val
}
