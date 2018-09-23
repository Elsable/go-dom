// Code generated DO NOT EDIT
// cssmarginrule.go
package dom

import "syscall/js"

type CSSMarginRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetName() string
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetStyle() CSSStyleDeclaration
	GetType() int
}
type CSSMarginRule struct {
	Value
	CSSRule
}

func jsValueToCSSMarginRule(val js.Value) CSSMarginRule {
	return CSSMarginRule{Value: Value{Value: val}}
}
func (v Value) AsCSSMarginRule() CSSMarginRule { return CSSMarginRule{Value: v} }
func (c CSSMarginRule) GetName() string {
	val := c.Get("name")
	return val.String()
}
func (c CSSMarginRule) GetStyle() CSSStyleDeclaration {
	val := c.Get("style")
	return jsValueToCSSStyleDeclaration(val.JSValue())
}