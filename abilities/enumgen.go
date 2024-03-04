// Code generated by "core generate"; DO NOT EDIT.

package abilities

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"cogentcore.org/core/enums"
)

var _AbilitiesValues = []Abilities{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// AbilitiesN is the highest valid value for type Abilities, plus one.
const AbilitiesN Abilities = 16

var _AbilitiesNameToValueMap = map[string]Abilities{`Editable`: 0, `Selectable`: 1, `Activatable`: 2, `Clickable`: 3, `DoubleClickable`: 4, `TripleClickable`: 5, `RepeatClickable`: 6, `LongPressable`: 7, `Draggable`: 8, `Droppable`: 9, `Slideable`: 10, `Checkable`: 11, `Scrollable`: 12, `Focusable`: 13, `Hoverable`: 14, `LongHoverable`: 15}

var _AbilitiesDescMap = map[Abilities]string{0: `Editable means the element can be edited. Otherwise, it remains in ReadOnly mode.`, 1: `Selectable means it can be Selected`, 2: `Activatable means it can be made Active by pressing down on it, which gives it a visible state layer color change. This also implies Clickable, receiving Click events when the user executes a mouse down and up event on the same element.`, 3: `Clickable means it can be Clicked, receiving Click events when the user executes a mouse down and up event on the same element, but otherwise does not change its rendering when pressed (as Activatable does). Use this for items that are more passively clickable, such as frames or tables, whereas e.g., a Button is Activatable.`, 4: `DoubleClickable indicates that an element does something different when it is clicked on twice in a row.`, 5: `TripleClickable indicates that an element does something different when it is clicked on three times in a row.`, 6: `RepeatClickable indicates that an element should receive repeated click events when the pointer is held down on it.`, 7: `LongPressable indicates that an element can be LongPressed`, 8: `Draggable means it can be Dragged`, 9: `Droppable means it can receive DragEnter, DragLeave, and Drop events (not specific to current Drag item, just generally)`, 10: `Slideable means it has a slider element that can be dragged to change value. Cannot be both Draggable and Slideable.`, 11: `Checkable means it can be Checked`, 12: `Scrollable means it can be Scrolled`, 13: `Focusable means it can be Focused: capable of receiving and processing key events directly and typically changing the style when focused to indicate this property to the user.`, 14: `Hoverable means it can be Hovered`, 15: `LongHoverable means it can be LongHovered`}

var _AbilitiesMap = map[Abilities]string{0: `Editable`, 1: `Selectable`, 2: `Activatable`, 3: `Clickable`, 4: `DoubleClickable`, 5: `TripleClickable`, 6: `RepeatClickable`, 7: `LongPressable`, 8: `Draggable`, 9: `Droppable`, 10: `Slideable`, 11: `Checkable`, 12: `Scrollable`, 13: `Focusable`, 14: `Hoverable`, 15: `LongHoverable`}

// String returns the string representation of this Abilities value.
func (i Abilities) String() string {
	str := ""
	for _, ie := range _AbilitiesValues {
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

// BitIndexString returns the string representation of this Abilities value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i Abilities) BitIndexString() string {
	if str, ok := _AbilitiesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Abilities value from its string representation,
// and returns an error if the string is invalid.
func (i *Abilities) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Abilities value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *Abilities) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _AbilitiesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type Abilities", flg)
		}
	}
	return nil
}

// Int64 returns the Abilities value as an int64.
func (i Abilities) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Abilities value from an int64.
func (i *Abilities) SetInt64(in int64) {
	*i = Abilities(in)
}

// Desc returns the description of the Abilities value.
func (i Abilities) Desc() string {
	if str, ok := _AbilitiesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AbilitiesValues returns all possible values for the type Abilities.
func AbilitiesValues() []Abilities {
	return _AbilitiesValues
}

// Values returns all possible values for the type Abilities.
func (i Abilities) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AbilitiesValues))
	for i, d := range _AbilitiesValues {
		res[i] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i Abilities) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *Abilities) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i Abilities) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Abilities) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Abilities.UnmarshalText:", err)
	}
	return nil
}
