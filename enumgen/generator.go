// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on http://github.com/dmarkham/enumer and
// golang.org/x/tools/cmd/stringer:

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"strings"

	"goki.dev/enums/enumgen/config"
	"goki.dev/grease"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

// Generator holds the state of the generator.
// It is primarily used to buffer the output.
type Generator struct {
	Config     *config.Config // The configuration information
	Buf        bytes.Buffer   // The accumulated output.
	Pkg        *Package       // The package we are scanning.
	Types      []Type         // The enum types
	HasBitFlag bool           // Whether there is any bit flag enum type in the package (used for determining imports)
}

// NewGenerator returns a new generator with the
// given configuration information.
func NewGenerator(config *config.Config) *Generator {
	return &Generator{Config: config}
}

// ParsePackage parses the single package located in the configuration directory.
func (g *Generator) ParsePackage() error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, g.Config.Dir)
	if err != nil {
		return err
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("expected 1 package, but found %d packages", len(pkgs))
	}
	g.AddPackage(pkgs[0])
	return nil
}

// AddPackage adds a type-checked Package and its syntax files to the generator.
func (g *Generator) AddPackage(pkg *packages.Package) {
	// fmt.Println(pkg.TypesInfo.Defs)
	g.Pkg = &Package{
		Name:  pkg.Name,
		Defs:  pkg.TypesInfo.Defs,
		Files: make([]*File, 0),
	}

	for _, file := range pkg.Syntax {
		// ignore generated code
		if ast.IsGenerated(file) {
			continue
		}
		// need to use append and 0 initial length
		// because we don't know if it has generated code
		g.Pkg.Files = append(g.Pkg.Files, &File{
			File: file,
			Pkg:  g.Pkg,
		})
	}
}

// Printf prints the formatted string to the
// accumulated output in [Generator.Buf]
func (g *Generator) Printf(format string, args ...any) {
	fmt.Fprintf(&g.Buf, format, args...)
}

// PrintHeader prints the header and package clause
// to the accumulated output
func (g *Generator) PrintHeader() {
	cmdstr := filepath.Base(os.Args[0])
	if len(os.Args) > 1 {
		cmdstr += " " + strings.Join(os.Args[1:], " ")
	}
	g.Printf("// Code generated by \"%s\"; DO NOT EDIT.\n", cmdstr)
	g.Printf("\n")
	if g.Config.Comment != "" {
		g.Printf("// %s\n", g.Config.Comment)
	}
	g.Printf("package %s", g.Pkg.Name)
	g.Printf("\n")
	// these import statements don't serve
	// much purpose, as goimports should handle it
	// anyway, but if goimports fails because the
	// code is invalid, it may make it easier to debug.
	// also, they don't serve much harm and make it clear
	// what packages are being used.
	// TODO: maybe remove these import statements
	g.Printf("import (\n")
	if g.Config.SQL || g.Config.GQL { // sql and gql are the only ones that use fmt
		g.Printf("\t\"fmt\"\n")
	}
	g.Printf("\t\"strings\"\n")
	g.Printf("\t\"strconv\"\n")
	g.Printf("\t\"errors\"\n")
	if g.HasBitFlag {
		g.Printf("\t\"sync/atomic\"\n")
	}
	if g.Config.SQL {
		g.Printf("\t\"database/sql/driver\"\n")
	}
	if g.Config.JSON {
		g.Printf("\t\"encoding/json\"\n")
	}
	if g.Config.GQL {
		g.Printf("\t\"io\"\n")
	}
	g.Printf("\t\"goki.dev/enums\"\n")
	g.Printf(")\n")
}

// FindEnumTypes goes through all of the types in the package
// and finds all integer (signed or unsigned) types labeled with enums:enum
// or enums:bitflag. It stores the resulting types in [Generator.Types].
func (g *Generator) FindEnumTypes() error {
	g.Types = []Type{}
	for _, file := range g.Pkg.Files {
		var terr error
		ast.Inspect(file.File, func(n ast.Node) bool {
			if terr != nil {
				return false
			}
			cont, err := g.InspectForType(n)
			if err != nil {
				terr = err
			}
			return cont
		})
		if terr != nil {
			return fmt.Errorf("FindEnumTypes: error finding enum types: %w", terr)
		}
	}
	return nil
}

// AllowedEnumTypes are the types that can be used for enums
// that are not bit flags (bit flags can only be int64s).
// It is stored as a map for quick and convenient access.
var AllowedEnumTypes = map[string]bool{"int": true, "int64": true, "int32": true, "int16": true, "int8": true, "uint": true, "uint64": true, "uint32": true, "uint16": true, "uint8": true}

// InspectForType looks at the given AST node and adds it
// to [Generator.Types] if it is marked with an appropriate
// comment directive. It returns whether the AST inspector should
// continue, and an error if there is one. It should only
// be called in [ast.Inspect].
func (g *Generator) InspectForType(n ast.Node) (bool, error) {
	ts, ok := n.(*ast.TypeSpec)
	if !ok {
		return true, nil
	}
	if ts.Comment == nil {
		return true, nil
	}
	for _, c := range ts.Comment.List {
		tool, directive, args, has, err := grease.ParseDirective(c.Text)
		if err != nil {
			return false, fmt.Errorf("error parsing comment directive %q: %w", c.Text, err)
		}
		if !has {
			continue
		}
		if tool != "enums" {
			continue
		}
		if directive != "enum" && directive != "bitflag" {
			return false, fmt.Errorf("unrecognized enums directive %q (from %q)", directive, c.Text)
		}

		ident, ok := ts.Type.(*ast.Ident)
		if !ok {
			return false, fmt.Errorf("type of enum type (%v) is %T, not *ast.Ident (try using a standard [un]signed integer type instead)", ts.Type, ts.Type)
		}
		cfg := &config.Config{}
		*cfg = *g.Config
		leftovers, err := grease.SetFromArgs(cfg, args)
		if err != nil {
			return false, fmt.Errorf("error setting config info from comment directive args: %w (from directive %q)", err, c.Text)
		}
		if len(leftovers) > 0 {
			return false, fmt.Errorf("expected 0 positional arguments but got %d (list: %v) (from directive %q)", len(leftovers), leftovers, c.Text)
		}

		typ := g.Pkg.Defs[ts.Name].Type()
		utyp := typ.Underlying()

		tt := Type{Name: ts.Name.Name, Type: ts, Config: cfg}
		if ident.String() != utyp.String() { // if our direct type isn't the same as our underlying type, we are extending our direct type
			tt.Extends = ident.String()
		}
		switch directive {
		case "enum":
			if !AllowedEnumTypes[utyp.String()] {
				return false, fmt.Errorf("enum type %s is not allowed; try using a standard [un]signed integer type instead", ident.Name)
			}
			tt.IsBitFlag = false
		case "bitflag":
			if utyp.String() != "int64" {
				return false, fmt.Errorf("bit flag enum type %s is not allowed; bit flag enums must be of type int64", ident.Name)
			}
			tt.IsBitFlag = true
			g.HasBitFlag = true
		}
		g.Types = append(g.Types, tt)

	}
	return true, nil
}

// Generate produces the enum methods for the types
// stored in [Generator.Types].
func (g *Generator) Generate() error {
	if len(g.Types) == 0 {
		return fmt.Errorf("no enum types found in package %q", g.Pkg.Name)
	}
	for _, typ := range g.Types {
		values := make([]Value, 0, 100)
		typeName := typ.Type.Name.String()
		for _, file := range g.Pkg.Files {
			file.Config = typ.Config
			// Set the state for this run of the walker.
			file.TypeName = typeName
			file.BitFlag = typ.IsBitFlag
			file.Values = nil
			if file.File != nil {
				var terr error
				ast.Inspect(file.File, func(n ast.Node) bool {
					if terr != nil {
						return false
					}
					cont, err := file.GenDecl(n)
					if err != nil {
						terr = err
					}
					return cont
				})
				if terr != nil {
					return fmt.Errorf("Generate: error parsing declaration clauses: %w", terr)
				}
				values = append(values, file.Values...)
			}
		}

		if len(values) == 0 {
			return errors.New("no values defined for type " + typeName)
		}

		g.TrimValueNames(values, typ.Config)

		err := g.TransformValueNames(values, typ.Config)
		if err != nil {
			return fmt.Errorf("error transforming value names: %w", err)
		}

		g.PrefixValueNames(values, typ.Config)

		runs := SplitIntoRuns(values)
		// The decision of which pattern to use depends on the number of
		// runs in the numbers. If there's only one, it's easy. For more than
		// one, there's a tradeoff between complexity and size of the data
		// and code vs. the simplicity of a map. A map takes more space,
		// but so does the code. The decision here (crossover at 10) is
		// arbitrary, but considers that for large numbers of runs the cost
		// of the linear scan in the switch might become important, and
		// rather than use yet another algorithm such as binary search,
		// we punt and use a map. In any case, the likelihood of a map
		// being necessary for any realistic example other than bitmasks
		// is very low. And bitmasks probably deserve their own analysis,
		// to be done some other day.
		const runsThreshold = 10
		typ.RunsThreshold = runsThreshold
		switch {
		case len(runs) == 1:
			g.BuildOneRun(runs, typ)
		case len(runs) <= runsThreshold:
			g.BuildMultipleRuns(runs, typeName, typ.IsBitFlag)
		default:
			g.BuildMap(runs, typeName, typ.IsBitFlag)
		}

		g.BuildNoOpOrderChangeDetect(runs, typeName)

		g.BuildBasicExtras(runs, typeName, typ.IsBitFlag, runsThreshold)
		if typ.IsBitFlag {
			g.BuildBitFlagMethods(runs, typeName)
		}

		if typ.Config.JSON {
			g.BuildJSONMethods(runs, typeName, runsThreshold)
		}
		if typ.Config.Text {
			g.BuildTextMethods(runs, typeName, runsThreshold)
		}
		if typ.Config.YAML {
			g.BuildYAMLMethods(runs, typeName, runsThreshold)
		}
		if typ.Config.SQL {
			g.AddValueAndScanMethod(typeName)
		}
		if typ.Config.GQL {
			g.BuildGQLMethods(runs, typeName)
		}
	}
	return nil
}

// Format returns the contents of the Generator's buffer
// ([Generator.Buf]) with goimports applied.
func (g *Generator) Format() ([]byte, error) {
	b, err := imports.Process(g.Config.Dir, g.Buf.Bytes(), nil)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		return g.Buf.Bytes(), errors.New("internal error: invalid Go generated: " + err.Error() + "; compile the package to analyze the error")
	}
	return b, nil
}

// Write formats the data in the the Generator's buffer
// ([Generator.Buf]) and writes it to the file specified by
// [Generator.Config.Output].
func (g *Generator) Write() error {
	b, ferr := g.Format()
	// we still write file even if formatting failed, as it is still useful
	// then we handle error later
	werr := os.WriteFile(g.Config.Output, b, 0666)
	if werr != nil {
		return fmt.Errorf("Generator.Write: error writing file: %w", werr)
	}
	if ferr != nil {
		return fmt.Errorf("Generator.Write: error formatting code: %w", ferr)
	}
	return nil
}
