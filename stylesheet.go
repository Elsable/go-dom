// Code generated DO NOT EDIT
// stylesheet.go
package dom

import "syscall/js"

type StyleSheetIFace interface {
	GetDisabled() bool
	SetDisabled(bool)
	GetHref() string
	GetMedia() MediaList
	GetOwnerNode()
	GetParentStyleSheet() StyleSheet
	GetTitle() string
	GetType() string
}
type StyleSheet struct {
	Value
}

func JSValueToStyleSheet(val js.Value) StyleSheet { return StyleSheet{Value: JSValueToValue(val)} }
func (v Value) AsStyleSheet() StyleSheet          { return StyleSheet{Value: v} }
func NewStyleSheet(args ...interface{}) StyleSheet {
	return StyleSheet{Value: JSValueToValue(js.Global().Get("StyleSheet").New(args...))}
}
func (s StyleSheet) GetDisabled() bool {
	val := s.Get("disabled")
	return val.Bool()
}
func (s StyleSheet) SetDisabled(val bool) {
	s.Set("disabled", val)
}
func (s StyleSheet) GetHref() string {
	val := s.Get("href")
	return val.String()
}
func (s StyleSheet) GetMedia() MediaList {
	val := s.Get("media")
	return JSValueToMediaList(val.JSValue())
}
func (s StyleSheet) GetOwnerNode() Value {
	val := s.Get("ownerNode")
	return val
}
func (s StyleSheet) GetParentStyleSheet() StyleSheet {
	val := s.Get("parentStyleSheet")
	return JSValueToStyleSheet(val.JSValue())
}
func (s StyleSheet) GetTitle() string {
	val := s.Get("title")
	return val.String()
}
func (s StyleSheet) GetType() string {
	val := s.Get("type")
	return val.String()
}
