// Code generated by "goki generate"; DO NOT EDIT.

package lms

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _OpponentsValues = []Opponents{0, 1, 2}

// OpponentsN is the highest valid value
// for type Opponents, plus one.
const OpponentsN Opponents = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _OpponentsNoOp() {
	var x [1]struct{}
	_ = x[WhiteBlack-(0)]
	_ = x[RedGreen-(1)]
	_ = x[BlueYellow-(2)]
}

var _OpponentsNameToValueMap = map[string]Opponents{
	`WhiteBlack`: 0,
	`whiteblack`: 0,
	`RedGreen`:   1,
	`redgreen`:   1,
	`BlueYellow`: 2,
	`blueyellow`: 2,
}

var _OpponentsDescMap = map[Opponents]string{
	0: `White vs. Black greyscale`,
	1: `Red vs. Green`,
	2: `Blue vs. Yellow`,
}

var _OpponentsMap = map[Opponents]string{
	0: `WhiteBlack`,
	1: `RedGreen`,
	2: `BlueYellow`,
}

// String returns the string representation
// of this Opponents value.
func (i Opponents) String() string {
	if str, ok := _OpponentsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Opponents value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Opponents) SetString(s string) error {
	if val, ok := _OpponentsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _OpponentsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Opponents")
}

// Int64 returns the Opponents value as an int64.
func (i Opponents) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Opponents value from an int64.
func (i *Opponents) SetInt64(in int64) {
	*i = Opponents(in)
}

// Desc returns the description of the Opponents value.
func (i Opponents) Desc() string {
	if str, ok := _OpponentsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// OpponentsValues returns all possible values
// for the type Opponents.
func OpponentsValues() []Opponents {
	return _OpponentsValues
}

// Values returns all possible values
// for the type Opponents.
func (i Opponents) Values() []enums.Enum {
	res := make([]enums.Enum, len(_OpponentsValues))
	for i, d := range _OpponentsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Opponents.
func (i Opponents) IsValid() bool {
	_, ok := _OpponentsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Opponents) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Opponents) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
