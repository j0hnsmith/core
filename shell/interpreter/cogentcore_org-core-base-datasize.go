// Code generated by 'yaegi extract cogentcore.org/core/base/datasize'. DO NOT EDIT.

package interpreter

import (
	"cogentcore.org/core/base/datasize"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/base/datasize/datasize"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"B":               reflect.ValueOf(datasize.B),
		"EB":              reflect.ValueOf(datasize.EB),
		"ErrBits":         reflect.ValueOf(&datasize.ErrBits).Elem(),
		"GB":              reflect.ValueOf(datasize.GB),
		"KB":              reflect.ValueOf(datasize.KB),
		"MB":              reflect.ValueOf(datasize.MB),
		"MustParse":       reflect.ValueOf(datasize.MustParse),
		"MustParseString": reflect.ValueOf(datasize.MustParseString),
		"PB":              reflect.ValueOf(datasize.PB),
		"Parse":           reflect.ValueOf(datasize.Parse),
		"ParseString":     reflect.ValueOf(datasize.ParseString),
		"TB":              reflect.ValueOf(datasize.TB),

		// type definitions
		"Size": reflect.ValueOf((*datasize.Size)(nil)),
	}
}