// Code generated DO NOT EDIT
// document.go
package dom

import "syscall/js"

type DocumentIFace interface {
	AddEventListener(args ...interface{})
	AdoptNode(args ...interface{}) Node
	Append(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetCharacterSet() string
	GetCharset() string
	GetChildElementCount() float64
	GetChildNodes() NodeList
	GetChildren() HTMLCollection
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	GetCompatMode() string
	Contains(args ...interface{}) bool
	GetContentType() string
	CreateAttribute(args ...interface{}) Attr
	CreateAttributeNS(args ...interface{}) Attr
	CreateCDATASection(args ...interface{}) CDATASection
	CreateComment(args ...interface{}) Comment
	CreateDocumentFragment(args ...interface{}) DocumentFragment
	CreateElement(args ...interface{}) Element
	CreateElementNS(args ...interface{}) Element
	CreateEvent(args ...interface{}) Event
	CreateNodeIterator(args ...interface{}) NodeIterator
	CreateProcessingInstruction(args ...interface{}) ProcessingInstruction
	CreateRange(args ...interface{}) Range
	CreateTextNode(args ...interface{}) Text
	CreateTreeWalker(args ...interface{}) TreeWalker
	DispatchEvent(args ...interface{}) bool
	GetDoctype() DocumentType
	GetDocumentElement() Element
	GetDocumentURI() string
	GetFirstChild() Node
	GetFirstElementChild() Element
	GetElementById(args ...interface{}) Element
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	GetImplementation() DOMImplementation
	ImportNode(args ...interface{}) Node
	GetInputEncoding() string
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLastElementChild() Element
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	GetOnauxclick() EventHandler
	SetOnauxclick(EventHandler)
	GetOnblur() EventHandler
	SetOnblur(EventHandler)
	GetOncancel() EventHandler
	SetOncancel(EventHandler)
	GetOncanplay() EventHandler
	SetOncanplay(EventHandler)
	GetOncanplaythrough() EventHandler
	SetOncanplaythrough(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnclick() EventHandler
	SetOnclick(EventHandler)
	GetOnclose() EventHandler
	SetOnclose(EventHandler)
	GetOncontextmenu() EventHandler
	SetOncontextmenu(EventHandler)
	GetOncopy() EventHandler
	SetOncopy(EventHandler)
	GetOncuechange() EventHandler
	SetOncuechange(EventHandler)
	GetOncut() EventHandler
	SetOncut(EventHandler)
	GetOndblclick() EventHandler
	SetOndblclick(EventHandler)
	GetOndrag() EventHandler
	SetOndrag(EventHandler)
	GetOndragend() EventHandler
	SetOndragend(EventHandler)
	GetOndragenter() EventHandler
	SetOndragenter(EventHandler)
	GetOndragexit() EventHandler
	SetOndragexit(EventHandler)
	GetOndragleave() EventHandler
	SetOndragleave(EventHandler)
	GetOndragover() EventHandler
	SetOndragover(EventHandler)
	GetOndragstart() EventHandler
	SetOndragstart(EventHandler)
	GetOndrop() EventHandler
	SetOndrop(EventHandler)
	GetOndurationchange() EventHandler
	SetOndurationchange(EventHandler)
	GetOnemptied() EventHandler
	SetOnemptied(EventHandler)
	GetOnended() EventHandler
	SetOnended(EventHandler)
	GetOnerror() OnErrorEventHandler
	SetOnerror(OnErrorEventHandler)
	GetOnfocus() EventHandler
	SetOnfocus(EventHandler)
	GetOninput() EventHandler
	SetOninput(EventHandler)
	GetOninvalid() EventHandler
	SetOninvalid(EventHandler)
	GetOnkeydown() EventHandler
	SetOnkeydown(EventHandler)
	GetOnkeypress() EventHandler
	SetOnkeypress(EventHandler)
	GetOnkeyup() EventHandler
	SetOnkeyup(EventHandler)
	GetOnload() EventHandler
	SetOnload(EventHandler)
	GetOnloadeddata() EventHandler
	SetOnloadeddata(EventHandler)
	GetOnloadedmetadata() EventHandler
	SetOnloadedmetadata(EventHandler)
	GetOnloadend() EventHandler
	SetOnloadend(EventHandler)
	GetOnloadstart() EventHandler
	SetOnloadstart(EventHandler)
	GetOnmousedown() EventHandler
	SetOnmousedown(EventHandler)
	GetOnmouseenter() EventHandler
	SetOnmouseenter(EventHandler)
	GetOnmouseleave() EventHandler
	SetOnmouseleave(EventHandler)
	GetOnmousemove() EventHandler
	SetOnmousemove(EventHandler)
	GetOnmouseout() EventHandler
	SetOnmouseout(EventHandler)
	GetOnmouseover() EventHandler
	SetOnmouseover(EventHandler)
	GetOnmouseup() EventHandler
	SetOnmouseup(EventHandler)
	GetOnpaste() EventHandler
	SetOnpaste(EventHandler)
	GetOnpause() EventHandler
	SetOnpause(EventHandler)
	GetOnplay() EventHandler
	SetOnplay(EventHandler)
	GetOnplaying() EventHandler
	SetOnplaying(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOnratechange() EventHandler
	SetOnratechange(EventHandler)
	GetOnreset() EventHandler
	SetOnreset(EventHandler)
	GetOnresize() EventHandler
	SetOnresize(EventHandler)
	GetOnscroll() EventHandler
	SetOnscroll(EventHandler)
	GetOnsecuritypolicyviolation() EventHandler
	SetOnsecuritypolicyviolation(EventHandler)
	GetOnseeked() EventHandler
	SetOnseeked(EventHandler)
	GetOnseeking() EventHandler
	SetOnseeking(EventHandler)
	GetOnselect() EventHandler
	SetOnselect(EventHandler)
	GetOnstalled() EventHandler
	SetOnstalled(EventHandler)
	GetOnsubmit() EventHandler
	SetOnsubmit(EventHandler)
	GetOnsuspend() EventHandler
	SetOnsuspend(EventHandler)
	GetOntimeupdate() EventHandler
	SetOntimeupdate(EventHandler)
	GetOntoggle() EventHandler
	SetOntoggle(EventHandler)
	GetOnvolumechange() EventHandler
	SetOnvolumechange(EventHandler)
	GetOnwaiting() EventHandler
	SetOnwaiting(EventHandler)
	GetOnwheel() EventHandler
	SetOnwheel(EventHandler)
	GetOrigin() string
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	Prepend(args ...interface{})
	GetPreviousSibling() Node
	QuerySelector(args ...interface{}) Element
	QuerySelectorAll(args ...interface{}) NodeList
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetTextContent() string
	SetTextContent(string)
	GetURL() string
}
type Document struct {
	Value
	Node
	EventTarget
}

func jsValueToDocument(val js.Value) Document { return Document{Value: Value{Value: val}} }
func (v Value) AsDocument() Document          { return Document{Value: v} }
func (d Document) AdoptNode(args ...interface{}) Node {
	val := d.Call("adoptNode", args...)
	return jsValueToNode(val.JSValue())
}
func (d Document) Append(args ...interface{}) {
	d.Call("append", args...)
}
func (d Document) GetCharacterSet() string {
	val := d.Get("characterSet")
	return val.String()
}
func (d Document) GetCharset() string {
	val := d.Get("charset")
	return val.String()
}
func (d Document) GetChildElementCount() float64 {
	val := d.Get("childElementCount")
	return val.Float()
}
func (d Document) GetChildren() HTMLCollection {
	val := d.Get("children")
	return jsValueToHTMLCollection(val.JSValue())
}
func (d Document) GetCompatMode() string {
	val := d.Get("compatMode")
	return val.String()
}
func (d Document) GetContentType() string {
	val := d.Get("contentType")
	return val.String()
}
func (d Document) CreateAttribute(args ...interface{}) Attr {
	val := d.Call("createAttribute", args...)
	return jsValueToAttr(val.JSValue())
}
func (d Document) CreateAttributeNS(args ...interface{}) Attr {
	val := d.Call("createAttributeNS", args...)
	return jsValueToAttr(val.JSValue())
}
func (d Document) CreateCDATASection(args ...interface{}) CDATASection {
	val := d.Call("createCDATASection", args...)
	return jsValueToCDATASection(val.JSValue())
}
func (d Document) CreateComment(args ...interface{}) Comment {
	val := d.Call("createComment", args...)
	return jsValueToComment(val.JSValue())
}
func (d Document) CreateDocumentFragment(args ...interface{}) DocumentFragment {
	val := d.Call("createDocumentFragment", args...)
	return jsValueToDocumentFragment(val.JSValue())
}
func (d Document) CreateElement(args ...interface{}) Element {
	val := d.Call("createElement", args...)
	return jsValueToElement(val.JSValue())
}
func (d Document) CreateElementNS(args ...interface{}) Element {
	val := d.Call("createElementNS", args...)
	return jsValueToElement(val.JSValue())
}
func (d Document) CreateEvent(args ...interface{}) Event {
	val := d.Call("createEvent", args...)
	return jsValueToEvent(val.JSValue())
}
func (d Document) CreateNodeIterator(args ...interface{}) NodeIterator {
	val := d.Call("createNodeIterator", args...)
	return jsValueToNodeIterator(val.JSValue())
}
func (d Document) CreateProcessingInstruction(args ...interface{}) ProcessingInstruction {
	val := d.Call("createProcessingInstruction", args...)
	return jsValueToProcessingInstruction(val.JSValue())
}
func (d Document) CreateRange(args ...interface{}) Range {
	val := d.Call("createRange", args...)
	return jsValueToRange(val.JSValue())
}
func (d Document) CreateTextNode(args ...interface{}) Text {
	val := d.Call("createTextNode", args...)
	return jsValueToText(val.JSValue())
}
func (d Document) CreateTreeWalker(args ...interface{}) TreeWalker {
	val := d.Call("createTreeWalker", args...)
	return jsValueToTreeWalker(val.JSValue())
}
func (d Document) GetDoctype() DocumentType {
	val := d.Get("doctype")
	return jsValueToDocumentType(val.JSValue())
}
func (d Document) GetDocumentElement() Element {
	val := d.Get("documentElement")
	return jsValueToElement(val.JSValue())
}
func (d Document) GetDocumentURI() string {
	val := d.Get("documentURI")
	return val.String()
}
func (d Document) GetFirstElementChild() Element {
	val := d.Get("firstElementChild")
	return jsValueToElement(val.JSValue())
}
func (d Document) GetElementById(args ...interface{}) Element {
	val := d.Call("getElementById", args...)
	return jsValueToElement(val.JSValue())
}
func (d Document) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := d.Call("getElementsByClassName", args...)
	return jsValueToHTMLCollection(val.JSValue())
}
func (d Document) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := d.Call("getElementsByTagName", args...)
	return jsValueToHTMLCollection(val.JSValue())
}
func (d Document) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := d.Call("getElementsByTagNameNS", args...)
	return jsValueToHTMLCollection(val.JSValue())
}
func (d Document) GetImplementation() DOMImplementation {
	val := d.Get("implementation")
	return jsValueToDOMImplementation(val.JSValue())
}
func (d Document) ImportNode(args ...interface{}) Node {
	val := d.Call("importNode", args...)
	return jsValueToNode(val.JSValue())
}
func (d Document) GetInputEncoding() string {
	val := d.Get("inputEncoding")
	return val.String()
}
func (d Document) GetLastElementChild() Element {
	val := d.Get("lastElementChild")
	return jsValueToElement(val.JSValue())
}
func (d Document) GetOnabort() EventHandler {
	val := d.Get("onabort")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnabort(val EventHandler) {
	d.Set("onabort", val)
}
func (d Document) GetOnauxclick() EventHandler {
	val := d.Get("onauxclick")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnauxclick(val EventHandler) {
	d.Set("onauxclick", val)
}
func (d Document) GetOnblur() EventHandler {
	val := d.Get("onblur")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnblur(val EventHandler) {
	d.Set("onblur", val)
}
func (d Document) GetOncancel() EventHandler {
	val := d.Get("oncancel")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncancel(val EventHandler) {
	d.Set("oncancel", val)
}
func (d Document) GetOncanplay() EventHandler {
	val := d.Get("oncanplay")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncanplay(val EventHandler) {
	d.Set("oncanplay", val)
}
func (d Document) GetOncanplaythrough() EventHandler {
	val := d.Get("oncanplaythrough")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncanplaythrough(val EventHandler) {
	d.Set("oncanplaythrough", val)
}
func (d Document) GetOnchange() EventHandler {
	val := d.Get("onchange")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnchange(val EventHandler) {
	d.Set("onchange", val)
}
func (d Document) GetOnclick() EventHandler {
	val := d.Get("onclick")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnclick(val EventHandler) {
	d.Set("onclick", val)
}
func (d Document) GetOnclose() EventHandler {
	val := d.Get("onclose")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnclose(val EventHandler) {
	d.Set("onclose", val)
}
func (d Document) GetOncontextmenu() EventHandler {
	val := d.Get("oncontextmenu")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncontextmenu(val EventHandler) {
	d.Set("oncontextmenu", val)
}
func (d Document) GetOncopy() EventHandler {
	val := d.Get("oncopy")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncopy(val EventHandler) {
	d.Set("oncopy", val)
}
func (d Document) GetOncuechange() EventHandler {
	val := d.Get("oncuechange")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncuechange(val EventHandler) {
	d.Set("oncuechange", val)
}
func (d Document) GetOncut() EventHandler {
	val := d.Get("oncut")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOncut(val EventHandler) {
	d.Set("oncut", val)
}
func (d Document) GetOndblclick() EventHandler {
	val := d.Get("ondblclick")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndblclick(val EventHandler) {
	d.Set("ondblclick", val)
}
func (d Document) GetOndrag() EventHandler {
	val := d.Get("ondrag")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndrag(val EventHandler) {
	d.Set("ondrag", val)
}
func (d Document) GetOndragend() EventHandler {
	val := d.Get("ondragend")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragend(val EventHandler) {
	d.Set("ondragend", val)
}
func (d Document) GetOndragenter() EventHandler {
	val := d.Get("ondragenter")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragenter(val EventHandler) {
	d.Set("ondragenter", val)
}
func (d Document) GetOndragexit() EventHandler {
	val := d.Get("ondragexit")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragexit(val EventHandler) {
	d.Set("ondragexit", val)
}
func (d Document) GetOndragleave() EventHandler {
	val := d.Get("ondragleave")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragleave(val EventHandler) {
	d.Set("ondragleave", val)
}
func (d Document) GetOndragover() EventHandler {
	val := d.Get("ondragover")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragover(val EventHandler) {
	d.Set("ondragover", val)
}
func (d Document) GetOndragstart() EventHandler {
	val := d.Get("ondragstart")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndragstart(val EventHandler) {
	d.Set("ondragstart", val)
}
func (d Document) GetOndrop() EventHandler {
	val := d.Get("ondrop")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndrop(val EventHandler) {
	d.Set("ondrop", val)
}
func (d Document) GetOndurationchange() EventHandler {
	val := d.Get("ondurationchange")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOndurationchange(val EventHandler) {
	d.Set("ondurationchange", val)
}
func (d Document) GetOnemptied() EventHandler {
	val := d.Get("onemptied")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnemptied(val EventHandler) {
	d.Set("onemptied", val)
}
func (d Document) GetOnended() EventHandler {
	val := d.Get("onended")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnended(val EventHandler) {
	d.Set("onended", val)
}
func (d Document) GetOnerror() OnErrorEventHandler {
	val := d.Get("onerror")
	return jsValueToOnErrorEventHandler(val.JSValue())
}
func (d Document) SetOnerror(val OnErrorEventHandler) {
	d.Set("onerror", val)
}
func (d Document) GetOnfocus() EventHandler {
	val := d.Get("onfocus")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnfocus(val EventHandler) {
	d.Set("onfocus", val)
}
func (d Document) GetOninput() EventHandler {
	val := d.Get("oninput")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOninput(val EventHandler) {
	d.Set("oninput", val)
}
func (d Document) GetOninvalid() EventHandler {
	val := d.Get("oninvalid")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOninvalid(val EventHandler) {
	d.Set("oninvalid", val)
}
func (d Document) GetOnkeydown() EventHandler {
	val := d.Get("onkeydown")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnkeydown(val EventHandler) {
	d.Set("onkeydown", val)
}
func (d Document) GetOnkeypress() EventHandler {
	val := d.Get("onkeypress")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnkeypress(val EventHandler) {
	d.Set("onkeypress", val)
}
func (d Document) GetOnkeyup() EventHandler {
	val := d.Get("onkeyup")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnkeyup(val EventHandler) {
	d.Set("onkeyup", val)
}
func (d Document) GetOnload() EventHandler {
	val := d.Get("onload")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnload(val EventHandler) {
	d.Set("onload", val)
}
func (d Document) GetOnloadeddata() EventHandler {
	val := d.Get("onloadeddata")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnloadeddata(val EventHandler) {
	d.Set("onloadeddata", val)
}
func (d Document) GetOnloadedmetadata() EventHandler {
	val := d.Get("onloadedmetadata")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnloadedmetadata(val EventHandler) {
	d.Set("onloadedmetadata", val)
}
func (d Document) GetOnloadend() EventHandler {
	val := d.Get("onloadend")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnloadend(val EventHandler) {
	d.Set("onloadend", val)
}
func (d Document) GetOnloadstart() EventHandler {
	val := d.Get("onloadstart")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnloadstart(val EventHandler) {
	d.Set("onloadstart", val)
}
func (d Document) GetOnmousedown() EventHandler {
	val := d.Get("onmousedown")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmousedown(val EventHandler) {
	d.Set("onmousedown", val)
}
func (d Document) GetOnmouseenter() EventHandler {
	val := d.Get("onmouseenter")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmouseenter(val EventHandler) {
	d.Set("onmouseenter", val)
}
func (d Document) GetOnmouseleave() EventHandler {
	val := d.Get("onmouseleave")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmouseleave(val EventHandler) {
	d.Set("onmouseleave", val)
}
func (d Document) GetOnmousemove() EventHandler {
	val := d.Get("onmousemove")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmousemove(val EventHandler) {
	d.Set("onmousemove", val)
}
func (d Document) GetOnmouseout() EventHandler {
	val := d.Get("onmouseout")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmouseout(val EventHandler) {
	d.Set("onmouseout", val)
}
func (d Document) GetOnmouseover() EventHandler {
	val := d.Get("onmouseover")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmouseover(val EventHandler) {
	d.Set("onmouseover", val)
}
func (d Document) GetOnmouseup() EventHandler {
	val := d.Get("onmouseup")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnmouseup(val EventHandler) {
	d.Set("onmouseup", val)
}
func (d Document) GetOnpaste() EventHandler {
	val := d.Get("onpaste")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnpaste(val EventHandler) {
	d.Set("onpaste", val)
}
func (d Document) GetOnpause() EventHandler {
	val := d.Get("onpause")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnpause(val EventHandler) {
	d.Set("onpause", val)
}
func (d Document) GetOnplay() EventHandler {
	val := d.Get("onplay")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnplay(val EventHandler) {
	d.Set("onplay", val)
}
func (d Document) GetOnplaying() EventHandler {
	val := d.Get("onplaying")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnplaying(val EventHandler) {
	d.Set("onplaying", val)
}
func (d Document) GetOnprogress() EventHandler {
	val := d.Get("onprogress")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnprogress(val EventHandler) {
	d.Set("onprogress", val)
}
func (d Document) GetOnratechange() EventHandler {
	val := d.Get("onratechange")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnratechange(val EventHandler) {
	d.Set("onratechange", val)
}
func (d Document) GetOnreset() EventHandler {
	val := d.Get("onreset")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnreset(val EventHandler) {
	d.Set("onreset", val)
}
func (d Document) GetOnresize() EventHandler {
	val := d.Get("onresize")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnresize(val EventHandler) {
	d.Set("onresize", val)
}
func (d Document) GetOnscroll() EventHandler {
	val := d.Get("onscroll")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnscroll(val EventHandler) {
	d.Set("onscroll", val)
}
func (d Document) GetOnsecuritypolicyviolation() EventHandler {
	val := d.Get("onsecuritypolicyviolation")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnsecuritypolicyviolation(val EventHandler) {
	d.Set("onsecuritypolicyviolation", val)
}
func (d Document) GetOnseeked() EventHandler {
	val := d.Get("onseeked")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnseeked(val EventHandler) {
	d.Set("onseeked", val)
}
func (d Document) GetOnseeking() EventHandler {
	val := d.Get("onseeking")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnseeking(val EventHandler) {
	d.Set("onseeking", val)
}
func (d Document) GetOnselect() EventHandler {
	val := d.Get("onselect")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnselect(val EventHandler) {
	d.Set("onselect", val)
}
func (d Document) GetOnstalled() EventHandler {
	val := d.Get("onstalled")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnstalled(val EventHandler) {
	d.Set("onstalled", val)
}
func (d Document) GetOnsubmit() EventHandler {
	val := d.Get("onsubmit")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnsubmit(val EventHandler) {
	d.Set("onsubmit", val)
}
func (d Document) GetOnsuspend() EventHandler {
	val := d.Get("onsuspend")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnsuspend(val EventHandler) {
	d.Set("onsuspend", val)
}
func (d Document) GetOntimeupdate() EventHandler {
	val := d.Get("ontimeupdate")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOntimeupdate(val EventHandler) {
	d.Set("ontimeupdate", val)
}
func (d Document) GetOntoggle() EventHandler {
	val := d.Get("ontoggle")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOntoggle(val EventHandler) {
	d.Set("ontoggle", val)
}
func (d Document) GetOnvolumechange() EventHandler {
	val := d.Get("onvolumechange")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnvolumechange(val EventHandler) {
	d.Set("onvolumechange", val)
}
func (d Document) GetOnwaiting() EventHandler {
	val := d.Get("onwaiting")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnwaiting(val EventHandler) {
	d.Set("onwaiting", val)
}
func (d Document) GetOnwheel() EventHandler {
	val := d.Get("onwheel")
	return jsValueToEventHandler(val.JSValue())
}
func (d Document) SetOnwheel(val EventHandler) {
	d.Set("onwheel", val)
}
func (d Document) GetOrigin() string {
	val := d.Get("origin")
	return val.String()
}
func (d Document) Prepend(args ...interface{}) {
	d.Call("prepend", args...)
}
func (d Document) QuerySelector(args ...interface{}) Element {
	val := d.Call("querySelector", args...)
	return jsValueToElement(val.JSValue())
}
func (d Document) QuerySelectorAll(args ...interface{}) NodeList {
	val := d.Call("querySelectorAll", args...)
	return jsValueToNodeList(val.JSValue())
}
func (d Document) GetURL() string {
	val := d.Get("URL")
	return val.String()
}