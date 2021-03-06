// Code generated DO NOT EDIT
// progressevent.go
package dom

import "syscall/js"

type ProgressEventIFace interface {
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
	GetLengthComputable() bool
	GetLoaded() int
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetTotal() int
	GetType() string
}
type ProgressEvent struct {
	Value
	Event
}

func JSValueToProgressEvent(val js.Value) ProgressEvent {
	return ProgressEvent{Value: JSValueToValue(val)}
}
func (v Value) AsProgressEvent() ProgressEvent { return ProgressEvent{Value: v} }
func NewProgressEvent(args ...interface{}) ProgressEvent {
	return ProgressEvent{Value: JSValueToValue(js.Global().Get("ProgressEvent").New(args...))}
}
func (p ProgressEvent) GetLengthComputable() bool {
	val := p.Get("lengthComputable")
	return val.Bool()
}
func (p ProgressEvent) GetLoaded() int {
	val := p.Get("loaded")
	return val.Int()
}
func (p ProgressEvent) GetTotal() int {
	val := p.Get("total")
	return val.Int()
}
