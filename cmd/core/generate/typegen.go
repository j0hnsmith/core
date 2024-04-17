// Code generated by "core generate"; DO NOT EDIT.

package generate

import (
	"cogentcore.org/core/types"
)

var _ = types.AddFunc(&types.Func{Name: "cogentcore.org/core/cmd/core/generate.Generate", Doc: "Generate is the main entry point to code generation\nthat does all of the generation according to the\ngiven config info. It overrides the\n[config.Config.Generate.Typegen.InterfaceConfigs] info.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})