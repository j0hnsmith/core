// Code generated by "core generate"; DO NOT EDIT.

package parse

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/ki"
)

var _ActionsValues = []Actions{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// ActionsN is the highest valid value for type Actions, plus one.
const ActionsN Actions = 10

var _ActionsNameToValueMap = map[string]Actions{`ChgToken`: 0, `AddSymbol`: 1, `PushScope`: 2, `PushNewScope`: 3, `PopScope`: 4, `PopScopeReg`: 5, `AddDetail`: 6, `AddType`: 7, `PushStack`: 8, `PopStack`: 9}

var _ActionsDescMap = map[Actions]string{0: `ChgToken changes the token to the Tok specified in the Act action`, 1: `AddSymbol means add name as a symbol, using current scoping and token type or the token specified in the Act action if != None`, 2: `PushScope means look for an existing symbol of given name to push onto current scope -- adding a new one if not found -- does not add new item to overall symbol list. This is useful for e.g., definitions of methods on a type, where this is not the definition of the type itself.`, 3: `PushNewScope means add a new symbol to the list and also push onto scope stack, using given token type or the token specified in the Act action if != None`, 4: `PopScope means remove the most recently-added scope item`, 5: `PopScopeReg means remove the most recently-added scope item, and also updates the source region for that item based on final SrcReg from corresponding Ast node -- for &#34;definitional&#34; scope`, 6: `AddDetail adds src at given path as detail info for the last-added symbol if there is already something there, a space is added for this new addition`, 7: `AddType Adds a type with the given name -- sets the Ast node for this rule and actual type is resolved later in a second language-specific pass`, 8: `PushStack adds name to stack -- provides context-sensitivity option for optimizing and ambiguity resolution`, 9: `PopStack pops the stack`}

var _ActionsMap = map[Actions]string{0: `ChgToken`, 1: `AddSymbol`, 2: `PushScope`, 3: `PushNewScope`, 4: `PopScope`, 5: `PopScopeReg`, 6: `AddDetail`, 7: `AddType`, 8: `PushStack`, 9: `PopStack`}

// String returns the string representation of this Actions value.
func (i Actions) String() string {
	if str, ok := _ActionsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Actions value from its string representation,
// and returns an error if the string is invalid.
func (i *Actions) SetString(s string) error {
	if val, ok := _ActionsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Actions")
}

// Int64 returns the Actions value as an int64.
func (i Actions) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Actions value from an int64.
func (i *Actions) SetInt64(in int64) {
	*i = Actions(in)
}

// Desc returns the description of the Actions value.
func (i Actions) Desc() string {
	if str, ok := _ActionsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ActionsValues returns all possible values for the type Actions.
func ActionsValues() []Actions {
	return _ActionsValues
}

// Values returns all possible values for the type Actions.
func (i Actions) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ActionsValues))
	for i, d := range _ActionsValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Actions) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Actions) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Actions.UnmarshalText:", err)
	}
	return nil
}

var _AstActsValues = []AstActs{0, 1, 2, 3, 4}

// AstActsN is the highest valid value for type AstActs, plus one.
const AstActsN AstActs = 5

var _AstActsNameToValueMap = map[string]AstActs{`NoAst`: 0, `AddAst`: 1, `SubAst`: 2, `AnchorAst`: 3, `AnchorFirstAst`: 4}

var _AstActsDescMap = map[AstActs]string{0: `NoAst means don&#39;t create an Ast node for this rule`, 1: `AddAst means create an Ast node for this rule, adding it to the current anchor Ast. Any sub-rules within this rule are *not* added as children of this node -- see SubAst and AnchorAst. This is good for token-only terminal nodes and list elements that should be added to a list.`, 2: `SubAst means create an Ast node and add all the elements of *this rule* as children of this new node (including sub-rules), *except* for the very last rule which is assumed to be a recursive rule -- that one goes back up to the parent node. This is good for adding more complex elements with sub-rules to a recursive list, without creating a new hierarchical depth level for every such element.`, 3: `AnchorAst means create an Ast node and set it as the anchor that subsequent sub-nodes are added into. This is for a new hierarchical depth level where everything under this rule gets organized.`, 4: `AnchorFirstAst means create an Ast node and set it as the anchor that subsequent sub-nodes are added into, *only* if this is the first time that this rule has matched within the current sequence (i.e., if the parent of this rule is the same rule then don&#39;t add a new Ast node). This is good for starting a new list of recursively-defined elements, without creating increasing depth levels.`}

var _AstActsMap = map[AstActs]string{0: `NoAst`, 1: `AddAst`, 2: `SubAst`, 3: `AnchorAst`, 4: `AnchorFirstAst`}

// String returns the string representation of this AstActs value.
func (i AstActs) String() string {
	if str, ok := _AstActsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the AstActs value from its string representation,
// and returns an error if the string is invalid.
func (i *AstActs) SetString(s string) error {
	if val, ok := _AstActsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type AstActs")
}

// Int64 returns the AstActs value as an int64.
func (i AstActs) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the AstActs value from an int64.
func (i *AstActs) SetInt64(in int64) {
	*i = AstActs(in)
}

// Desc returns the description of the AstActs value.
func (i AstActs) Desc() string {
	if str, ok := _AstActsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AstActsValues returns all possible values for the type AstActs.
func AstActsValues() []AstActs {
	return _AstActsValues
}

// Values returns all possible values for the type AstActs.
func (i AstActs) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AstActsValues))
	for i, d := range _AstActsValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i AstActs) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *AstActs) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("AstActs.UnmarshalText:", err)
	}
	return nil
}

var _RuleFlagsValues = []RuleFlags{7, 8, 9, 10, 11, 12, 13}

// RuleFlagsN is the highest valid value for type RuleFlags, plus one.
const RuleFlagsN RuleFlags = 14

var _RuleFlagsNameToValueMap = map[string]RuleFlags{`SetsScope`: 7, `Reverse`: 8, `NoToks`: 9, `OnlyToks`: 10, `MatchEOS`: 11, `MultiEOS`: 12, `TokMatchGroup`: 13}

var _RuleFlagsDescMap = map[RuleFlags]string{7: `SetsScope means that this rule sets its own scope, because it ends with EOS`, 8: `Reverse means that this rule runs in reverse (starts with - sign) -- for arithmetic binary expressions only: this is needed to produce proper associativity result for mathematical expressions in the recursive descent parser. Only for rules of form: Expr &#39;+&#39; Expr -- two sub-rules with a token operator in the middle.`, 9: `NoToks means that this rule doesn&#39;t have any explicit tokens -- only refers to other rules`, 10: `OnlyToks means that this rule only has explicit tokens for matching -- can be optimized`, 11: `MatchEOS means that the rule ends with a *matched* EOS with StInc = 1. SetsScope applies for optional and matching EOS rules alike.`, 12: `MultiEOS means that the rule has multiple EOS tokens within it -- changes some of the logic`, 13: `TokMatchGroup is a group node that also has a single token match, so it can be used in a FirstTokMap to optimize lookup of rules`}

var _RuleFlagsMap = map[RuleFlags]string{7: `SetsScope`, 8: `Reverse`, 9: `NoToks`, 10: `OnlyToks`, 11: `MatchEOS`, 12: `MultiEOS`, 13: `TokMatchGroup`}

// String returns the string representation of this RuleFlags value.
func (i RuleFlags) String() string {
	str := ""
	for _, ie := range ki.FlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _RuleFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string representation of this RuleFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i RuleFlags) BitIndexString() string {
	if str, ok := _RuleFlagsMap[i]; ok {
		return str
	}
	return ki.Flags(i).BitIndexString()
}

// SetString sets the RuleFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *RuleFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the RuleFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *RuleFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _RuleFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*ki.Flags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the RuleFlags value as an int64.
func (i RuleFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the RuleFlags value from an int64.
func (i *RuleFlags) SetInt64(in int64) {
	*i = RuleFlags(in)
}

// Desc returns the description of the RuleFlags value.
func (i RuleFlags) Desc() string {
	if str, ok := _RuleFlagsDescMap[i]; ok {
		return str
	}
	return ki.Flags(i).Desc()
}

// RuleFlagsValues returns all possible values for the type RuleFlags.
func RuleFlagsValues() []RuleFlags {
	es := ki.FlagsValues()
	res := make([]RuleFlags, len(es))
	for i, e := range es {
		res[i] = RuleFlags(e)
	}
	res = append(res, _RuleFlagsValues...)
	return res
}

// Values returns all possible values for the type RuleFlags.
func (i RuleFlags) Values() []enums.Enum {
	es := ki.FlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_RuleFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _RuleFlagsValues {
		res[i+les] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i RuleFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *RuleFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i RuleFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *RuleFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("RuleFlags.UnmarshalText:", err)
	}
	return nil
}

var _StepsValues = []Steps{0, 1, 2, 3, 4}

// StepsN is the highest valid value for type Steps, plus one.
const StepsN Steps = 5

var _StepsNameToValueMap = map[string]Steps{`Match`: 0, `SubMatch`: 1, `NoMatch`: 2, `Run`: 3, `RunAct`: 4}

var _StepsDescMap = map[Steps]string{0: `Match happens when a rule matches`, 1: `SubMatch is when a sub-rule within a rule matches`, 2: `NoMatch is when the rule fails to match (recorded at first non-match, which terminates matching process`, 3: `Run is when the rule is running and iterating through its sub-rules`, 4: `RunAct is when the rule is running and performing actions`}

var _StepsMap = map[Steps]string{0: `Match`, 1: `SubMatch`, 2: `NoMatch`, 3: `Run`, 4: `RunAct`}

// String returns the string representation of this Steps value.
func (i Steps) String() string {
	if str, ok := _StepsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Steps value from its string representation,
// and returns an error if the string is invalid.
func (i *Steps) SetString(s string) error {
	if val, ok := _StepsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Steps")
}

// Int64 returns the Steps value as an int64.
func (i Steps) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Steps value from an int64.
func (i *Steps) SetInt64(in int64) {
	*i = Steps(in)
}

// Desc returns the description of the Steps value.
func (i Steps) Desc() string {
	if str, ok := _StepsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StepsValues returns all possible values for the type Steps.
func StepsValues() []Steps {
	return _StepsValues
}

// Values returns all possible values for the type Steps.
func (i Steps) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StepsValues))
	for i, d := range _StepsValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Steps) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Steps) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Steps.UnmarshalText:", err)
	}
	return nil
}
