// Code generated by "core generate"; DO NOT EDIT.

package config

import (
	"errors"
	"log"
	"strconv"

	"cogentcore.org/core/enums"
)

var _TypesValues = []Types{0, 1}

// TypesN is the highest valid value for type Types, plus one.
const TypesN Types = 2

var _TypesNameToValueMap = map[string]Types{`App`: 0, `Library`: 1}

var _TypesDescMap = map[Types]string{0: `TypeApp is an executable app`, 1: `TypeLibrary is an importable library`}

var _TypesMap = map[Types]string{0: `App`, 1: `Library`}

// String returns the string representation of this Types value.
func (i Types) String() string {
	if str, ok := _TypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Types value from its string representation,
// and returns an error if the string is invalid.
func (i *Types) SetString(s string) error {
	if val, ok := _TypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Types")
}

// Int64 returns the Types value as an int64.
func (i Types) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Types value from an int64.
func (i *Types) SetInt64(in int64) {
	*i = Types(in)
}

// Desc returns the description of the Types value.
func (i Types) Desc() string {
	if str, ok := _TypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TypesValues returns all possible values for the type Types.
func TypesValues() []Types {
	return _TypesValues
}

// Values returns all possible values for the type Types.
func (i Types) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TypesValues))
	for i, d := range _TypesValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Types) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Types) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Types.UnmarshalText:", err)
	}
	return nil
}
