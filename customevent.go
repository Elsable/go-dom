// Code generated DO NOT EDIT
// customevent.go
package dom

import "syscall/js"

type CustomEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetDetail() Value
	GetEventPhase() int
	InitCustomEvent(args ...interface{})
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type CustomEvent struct {
	Value
	Event
}

func JSValueToCustomEvent(val js.Value) CustomEvent { return CustomEvent{Value: JSValueToValue(val)} }
func (v Value) AsCustomEvent() CustomEvent          { return CustomEvent{Value: v} }
func NewCustomEvent(args ...interface{}) CustomEvent {
	return CustomEvent{Value: JSValueToValue(js.Global().Get("CustomEvent").New(args...))}
}
func (c CustomEvent) GetDetail() Value {
	val := c.Get("detail")
	return val
}
func (c CustomEvent) InitCustomEvent(args ...interface{}) {
	c.Call("initCustomEvent", args...)
}
