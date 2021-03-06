// Code generated DO NOT EDIT
// htmltrackelement.go
package dom

import "syscall/js"

type HTMLTrackElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetDefault() bool
	SetDefault(bool)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetAttribute(args ...interface{}) string
	GetAttributeNS(args ...interface{}) string
	GetAttributeNames(args ...interface{})
	GetAttributeNode(args ...interface{}) Attr
	GetAttributeNodeNS(args ...interface{}) Attr
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHidden() bool
	SetHidden(bool)
	GetId() string
	SetId(string)
	GetInnerText() string
	SetInnerText(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetKind() string
	SetKind(string)
	GetLabel() string
	SetLabel(string)
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	GetReadyState() int
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetSrclang() string
	SetSrclang(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTrack() TextTrack
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLTrackElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLTrackElement(val js.Value) HTMLTrackElement {
	return HTMLTrackElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTrackElement() HTMLTrackElement { return HTMLTrackElement{Value: v} }
func NewHTMLTrackElement(args ...interface{}) HTMLTrackElement {
	return HTMLTrackElement{Value: JSValueToValue(js.Global().Get("HTMLTrackElement").New(args...))}
}
func (h HTMLTrackElement) GetDefault() bool {
	val := h.Get("default")
	return val.Bool()
}
func (h HTMLTrackElement) SetDefault(val bool) {
	h.Set("default", val)
}
func (h HTMLTrackElement) GetKind() string {
	val := h.Get("kind")
	return val.String()
}
func (h HTMLTrackElement) SetKind(val string) {
	h.Set("kind", val)
}
func (h HTMLTrackElement) GetLabel() string {
	val := h.Get("label")
	return val.String()
}
func (h HTMLTrackElement) SetLabel(val string) {
	h.Set("label", val)
}
func (h HTMLTrackElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLTrackElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLTrackElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLTrackElement) GetSrclang() string {
	val := h.Get("srclang")
	return val.String()
}
func (h HTMLTrackElement) SetSrclang(val string) {
	h.Set("srclang", val)
}
func (h HTMLTrackElement) GetTrack() TextTrack {
	val := h.Get("track")
	return JSValueToTextTrack(val.JSValue())
}
