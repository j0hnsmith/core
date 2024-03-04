// Code generated by "core generate"; DO NOT EDIT.

package raster

import (
	"errors"
	"log"
	"strconv"

	"cogentcore.org/core/enums"
)

var _PathCommandValues = []PathCommand{0, 1, 2, 3, 4}

// PathCommandN is the highest valid value for type PathCommand, plus one.
const PathCommandN PathCommand = 5

var _PathCommandNameToValueMap = map[string]PathCommand{`PathMoveTo`: 0, `PathLineTo`: 1, `PathQuadTo`: 2, `PathCubicTo`: 3, `PathClose`: 4}

var _PathCommandDescMap = map[PathCommand]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``}

var _PathCommandMap = map[PathCommand]string{0: `PathMoveTo`, 1: `PathLineTo`, 2: `PathQuadTo`, 3: `PathCubicTo`, 4: `PathClose`}

// String returns the string representation of this PathCommand value.
func (i PathCommand) String() string {
	if str, ok := _PathCommandMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the PathCommand value from its string representation,
// and returns an error if the string is invalid.
func (i *PathCommand) SetString(s string) error {
	if val, ok := _PathCommandNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type PathCommand")
}

// Int64 returns the PathCommand value as an int64.
func (i PathCommand) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the PathCommand value from an int64.
func (i *PathCommand) SetInt64(in int64) {
	*i = PathCommand(in)
}

// Desc returns the description of the PathCommand value.
func (i PathCommand) Desc() string {
	if str, ok := _PathCommandDescMap[i]; ok {
		return str
	}
	return i.String()
}

// PathCommandValues returns all possible values for the type PathCommand.
func PathCommandValues() []PathCommand {
	return _PathCommandValues
}

// Values returns all possible values for the type PathCommand.
func (i PathCommand) Values() []enums.Enum {
	res := make([]enums.Enum, len(_PathCommandValues))
	for i, d := range _PathCommandValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i PathCommand) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *PathCommand) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("PathCommand.UnmarshalText:", err)
	}
	return nil
}

var _JoinModeValues = []JoinMode{0, 1, 2, 3, 4, 5}

// JoinModeN is the highest valid value for type JoinMode, plus one.
const JoinModeN JoinMode = 6

var _JoinModeNameToValueMap = map[string]JoinMode{`Arc`: 0, `ArcClip`: 1, `Miter`: 2, `MiterClip`: 3, `Bevel`: 4, `Round`: 5}

var _JoinModeDescMap = map[JoinMode]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``}

var _JoinModeMap = map[JoinMode]string{0: `Arc`, 1: `ArcClip`, 2: `Miter`, 3: `MiterClip`, 4: `Bevel`, 5: `Round`}

// String returns the string representation of this JoinMode value.
func (i JoinMode) String() string {
	if str, ok := _JoinModeMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the JoinMode value from its string representation,
// and returns an error if the string is invalid.
func (i *JoinMode) SetString(s string) error {
	if val, ok := _JoinModeNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type JoinMode")
}

// Int64 returns the JoinMode value as an int64.
func (i JoinMode) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the JoinMode value from an int64.
func (i *JoinMode) SetInt64(in int64) {
	*i = JoinMode(in)
}

// Desc returns the description of the JoinMode value.
func (i JoinMode) Desc() string {
	if str, ok := _JoinModeDescMap[i]; ok {
		return str
	}
	return i.String()
}

// JoinModeValues returns all possible values for the type JoinMode.
func JoinModeValues() []JoinMode {
	return _JoinModeValues
}

// Values returns all possible values for the type JoinMode.
func (i JoinMode) Values() []enums.Enum {
	res := make([]enums.Enum, len(_JoinModeValues))
	for i, d := range _JoinModeValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i JoinMode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *JoinMode) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("JoinMode.UnmarshalText:", err)
	}
	return nil
}
