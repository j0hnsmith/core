// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gti

import (
	"reflect"
)

// Type represents a type
type Type struct {
	// Name is the fully-package-path-qualified name of the type (eg: goki.dev/gi/v2/gi.Button)
	Name string

	// Doc has all of the comment documentation
	// info as one string with directives removed.
	Doc string

	// Directives has the parsed comment directives
	Directives Directives

	// unique type ID number
	ID uint64 `desc:"unique type ID number"`

	// Methods are available for all types
	Methods *Methods

	// Embeded fields for struct types
	Embeds *Fields

	// Fields for struct types
	Fields *Fields

	// instance of the type
	Instance any `desc:"instance of the type"`
}

/*
// NewType creates a new Type for given instance. This call is auto-generated for each Ki type.
func NewType(nm string, inst Ki) *Type {
	inst.InitName(inst, nm)
	tp := &Type{Name: nm, Instance: inst}
	tp.ID = atomic.AddUint64(&TypeIDCounter, 1)
	TypeRegistry[nm] = tp
	return tp
}

// NewInstance returns a new instance of given type
// Note: otherwise impossible to generate new instance generically, unless using reflection
func (tp *Type) NewInstance() Ki {
	return tp.Instance.NewInstance()
}
*/

// ReflectType returns the [reflect.Type] of a given Ki Type
func (tp *Type) ReflectType() reflect.Type {
	return reflect.TypeOf(tp.Instance).Elem()
}
