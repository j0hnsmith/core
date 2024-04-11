// Code generated by "core generate"; DO NOT EDIT.

package gradient

import (
	"cogentcore.org/core/enums"
)

var _SpreadsValues = []Spreads{0, 1, 2}

// SpreadsN is the highest valid value for type Spreads, plus one.
const SpreadsN Spreads = 3

var _SpreadsValueMap = map[string]Spreads{`pad`: 0, `reflect`: 1, `repeat`: 2}

var _SpreadsDescMap = map[Spreads]string{0: `Pad indicates to have the final color of the gradient fill the object beyond the end of the gradient.`, 1: `Reflect indicates to have a gradient repeat in reverse order (offset 1 to 0) to fully fill an object beyond the end of the gradient.`, 2: `Repeat indicates to have a gradient continue in its original order (offset 0 to 1) by jumping back to the start to fully fill an object beyond the end of the gradient.`}

var _SpreadsMap = map[Spreads]string{0: `pad`, 1: `reflect`, 2: `repeat`}

// String returns the string representation of this Spreads value.
func (i Spreads) String() string { return enums.String(i, _SpreadsMap) }

// SetString sets the Spreads value from its string representation,
// and returns an error if the string is invalid.
func (i *Spreads) SetString(s string) error {
	return enums.SetString(i, s, _SpreadsValueMap, "Spreads")
}

// Int64 returns the Spreads value as an int64.
func (i Spreads) Int64() int64 { return int64(i) }

// SetInt64 sets the Spreads value from an int64.
func (i *Spreads) SetInt64(in int64) { *i = Spreads(in) }

// Desc returns the description of the Spreads value.
func (i Spreads) Desc() string { return enums.Desc(i, _SpreadsDescMap) }

// SpreadsValues returns all possible values for the type Spreads.
func SpreadsValues() []Spreads { return _SpreadsValues }

// Values returns all possible values for the type Spreads.
func (i Spreads) Values() []enums.Enum { return enums.Values(_SpreadsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Spreads) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Spreads) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Spreads") }

var _UnitsValues = []Units{0, 1}

// UnitsN is the highest valid value for type Units, plus one.
const UnitsN Units = 2

var _UnitsValueMap = map[string]Units{`objectBoundingBox`: 0, `userSpaceOnUse`: 1}

var _UnitsDescMap = map[Units]string{0: `ObjectBoundingBox indicates that coordinate values are scaled relative to the size of the object and are specified in the normalized range of 0 to 1.`, 1: `UserSpaceOnUse indicates that coordinate values are specified in the current user coordinate system when the gradient is used (ie: actual SVG/core coordinates).`}

var _UnitsMap = map[Units]string{0: `objectBoundingBox`, 1: `userSpaceOnUse`}

// String returns the string representation of this Units value.
func (i Units) String() string { return enums.String(i, _UnitsMap) }

// SetString sets the Units value from its string representation,
// and returns an error if the string is invalid.
func (i *Units) SetString(s string) error { return enums.SetString(i, s, _UnitsValueMap, "Units") }

// Int64 returns the Units value as an int64.
func (i Units) Int64() int64 { return int64(i) }

// SetInt64 sets the Units value from an int64.
func (i *Units) SetInt64(in int64) { *i = Units(in) }

// Desc returns the description of the Units value.
func (i Units) Desc() string { return enums.Desc(i, _UnitsDescMap) }

// UnitsValues returns all possible values for the type Units.
func UnitsValues() []Units { return _UnitsValues }

// Values returns all possible values for the type Units.
func (i Units) Values() []enums.Enum { return enums.Values(_UnitsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Units) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Units) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Units") }
