// Code generated by 'yaegi extract cogentcore.org/core/htmlcore'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/htmlcore"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/htmlcore/htmlcore"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BindTextEditor":   reflect.ValueOf(&htmlcore.BindTextEditor).Elem(),
		"ElementHandlers":  reflect.ValueOf(&htmlcore.ElementHandlers).Elem(),
		"ExtractText":      reflect.ValueOf(htmlcore.ExtractText),
		"Get":              reflect.ValueOf(htmlcore.Get),
		"GetAttr":          reflect.ValueOf(htmlcore.GetAttr),
		"HandleElement":    reflect.ValueOf(htmlcore.HandleElement),
		"HandleText":       reflect.ValueOf(htmlcore.HandleText),
		"HandleTextTag":    reflect.ValueOf(htmlcore.HandleTextTag),
		"HasAttr":          reflect.ValueOf(htmlcore.HasAttr),
		"IsText":           reflect.ValueOf(htmlcore.IsText),
		"IsURL":            reflect.ValueOf(htmlcore.IsURL),
		"NewContext":       reflect.ValueOf(htmlcore.NewContext),
		"NodeString":       reflect.ValueOf(htmlcore.NodeString),
		"NormalizeURL":     reflect.ValueOf(htmlcore.NormalizeURL),
		"ParseRelativeURL": reflect.ValueOf(htmlcore.ParseRelativeURL),
		"ParseURL":         reflect.ValueOf(htmlcore.ParseURL),
		"PkgGoDevWikilink": reflect.ValueOf(htmlcore.PkgGoDevWikilink),
		"ReadHTML":         reflect.ValueOf(htmlcore.ReadHTML),
		"ReadHTMLNode":     reflect.ValueOf(htmlcore.ReadHTMLNode),
		"ReadHTMLString":   reflect.ValueOf(htmlcore.ReadHTMLString),
		"ReadMD":           reflect.ValueOf(htmlcore.ReadMD),
		"ReadMDString":     reflect.ValueOf(htmlcore.ReadMDString),
		"RootNode":         reflect.ValueOf(htmlcore.RootNode),
		"TextTags":         reflect.ValueOf(&htmlcore.TextTags).Elem(),
		"UserAgentStyles":  reflect.ValueOf(&htmlcore.UserAgentStyles).Elem(),

		// type definitions
		"Context": reflect.ValueOf((*htmlcore.Context)(nil)),
	}
}