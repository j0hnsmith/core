// Code generated by "goki generate"; DO NOT EDIT.

package vci

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _FileStatusValues = []FileStatus{0, 1, 2, 3, 4, 5, 6}

// FileStatusN is the highest valid value
// for type FileStatus, plus one.
const FileStatusN FileStatus = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FileStatusNoOp() {
	var x [1]struct{}
	_ = x[Untracked-(0)]
	_ = x[Stored-(1)]
	_ = x[Modified-(2)]
	_ = x[Added-(3)]
	_ = x[Deleted-(4)]
	_ = x[Conflicted-(5)]
	_ = x[Updated-(6)]
}

var _FileStatusNameToValueMap = map[string]FileStatus{
	`Untracked`:  0,
	`untracked`:  0,
	`Stored`:     1,
	`stored`:     1,
	`Modified`:   2,
	`modified`:   2,
	`Added`:      3,
	`added`:      3,
	`Deleted`:    4,
	`deleted`:    4,
	`Conflicted`: 5,
	`conflicted`: 5,
	`Updated`:    6,
	`updated`:    6,
}

var _FileStatusDescMap = map[FileStatus]string{
	0: `Untracked means file is not under VCS control`,
	1: `Stored means file is stored under VCS control, and has not been modified in working copy`,
	2: `Modified means file is under VCS control, and has been modified in working copy`,
	3: `Added means file has just been added to VCS but is not yet committed`,
	4: `Deleted means file has been deleted from VCS`,
	5: `Conflicted means file is in conflict -- has not been merged`,
	6: `Updated means file has been updated in the remote but not locally`,
}

var _FileStatusMap = map[FileStatus]string{
	0: `Untracked`,
	1: `Stored`,
	2: `Modified`,
	3: `Added`,
	4: `Deleted`,
	5: `Conflicted`,
	6: `Updated`,
}

// String returns the string representation
// of this FileStatus value.
func (i FileStatus) String() string {
	if str, ok := _FileStatusMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the FileStatus value from its
// string representation, and returns an
// error if the string is invalid.
func (i *FileStatus) SetString(s string) error {
	if val, ok := _FileStatusNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _FileStatusNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type FileStatus")
}

// Int64 returns the FileStatus value as an int64.
func (i FileStatus) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the FileStatus value from an int64.
func (i *FileStatus) SetInt64(in int64) {
	*i = FileStatus(in)
}

// Desc returns the description of the FileStatus value.
func (i FileStatus) Desc() string {
	if str, ok := _FileStatusDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FileStatusValues returns all possible values
// for the type FileStatus.
func FileStatusValues() []FileStatus {
	return _FileStatusValues
}

// Values returns all possible values
// for the type FileStatus.
func (i FileStatus) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FileStatusValues))
	for i, d := range _FileStatusValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type FileStatus.
func (i FileStatus) IsValid() bool {
	_, ok := _FileStatusMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i FileStatus) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *FileStatus) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
