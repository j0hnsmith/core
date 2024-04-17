// Code generated by "core generate"; DO NOT EDIT.

package system

import (
	"cogentcore.org/core/enums"
)

var _PlatformsValues = []Platforms{0, 1, 2, 3, 4, 5, 6}

// PlatformsN is the highest valid value for type Platforms, plus one.
const PlatformsN Platforms = 7

var _PlatformsValueMap = map[string]Platforms{`MacOS`: 0, `Linux`: 1, `Windows`: 2, `IOS`: 3, `Android`: 4, `Web`: 5, `Offscreen`: 6}

var _PlatformsDescMap = map[Platforms]string{0: `MacOS is a Mac OS machine (aka Darwin)`, 1: `Linux is a Linux OS machine`, 2: `Windows is a Microsoft Windows machine`, 3: `IOS is an Apple iOS or iPadOS mobile phone or iPad`, 4: `Android is an Android mobile phone or tablet`, 5: `Web is a web browser running the app through WASM`, 6: `Offscreen is an offscreen driver typically used for testing, specified using the &#34;offscreen&#34; build tag`}

var _PlatformsMap = map[Platforms]string{0: `MacOS`, 1: `Linux`, 2: `Windows`, 3: `IOS`, 4: `Android`, 5: `Web`, 6: `Offscreen`}

// String returns the string representation of this Platforms value.
func (i Platforms) String() string { return enums.String(i, _PlatformsMap) }

// SetString sets the Platforms value from its string representation,
// and returns an error if the string is invalid.
func (i *Platforms) SetString(s string) error {
	return enums.SetString(i, s, _PlatformsValueMap, "Platforms")
}

// Int64 returns the Platforms value as an int64.
func (i Platforms) Int64() int64 { return int64(i) }

// SetInt64 sets the Platforms value from an int64.
func (i *Platforms) SetInt64(in int64) { *i = Platforms(in) }

// Desc returns the description of the Platforms value.
func (i Platforms) Desc() string { return enums.Desc(i, _PlatformsDescMap) }

// PlatformsValues returns all possible values for the type Platforms.
func PlatformsValues() []Platforms { return _PlatformsValues }

// Values returns all possible values for the type Platforms.
func (i Platforms) Values() []enums.Enum { return enums.Values(_PlatformsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Platforms) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Platforms) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "Platforms")
}

var _ScreenOrientationValues = []ScreenOrientation{0, 1, 2}

// ScreenOrientationN is the highest valid value for type ScreenOrientation, plus one.
const ScreenOrientationN ScreenOrientation = 3

var _ScreenOrientationValueMap = map[string]ScreenOrientation{`OrientationUnknown`: 0, `Portrait`: 1, `Landscape`: 2}

var _ScreenOrientationDescMap = map[ScreenOrientation]string{0: `OrientationUnknown means device orientation cannot be determined. Equivalent on Android to Configuration.ORIENTATION_UNKNOWN and on iOS to: UIDeviceOrientationUnknown UIDeviceOrientationFaceUp UIDeviceOrientationFaceDown`, 1: `Portrait is a device oriented so it is tall and thin. Equivalent on Android to Configuration.ORIENTATION_PORTRAIT and on iOS to: UIDeviceOrientationPortrait UIDeviceOrientationPortraitUpsideDown`, 2: `Landscape is a device oriented so it is short and wide. Equivalent on Android to Configuration.ORIENTATION_LANDSCAPE and on iOS to: UIDeviceOrientationLandscapeLeft UIDeviceOrientationLandscapeRight`}

var _ScreenOrientationMap = map[ScreenOrientation]string{0: `OrientationUnknown`, 1: `Portrait`, 2: `Landscape`}

// String returns the string representation of this ScreenOrientation value.
func (i ScreenOrientation) String() string { return enums.String(i, _ScreenOrientationMap) }

// SetString sets the ScreenOrientation value from its string representation,
// and returns an error if the string is invalid.
func (i *ScreenOrientation) SetString(s string) error {
	return enums.SetString(i, s, _ScreenOrientationValueMap, "ScreenOrientation")
}

// Int64 returns the ScreenOrientation value as an int64.
func (i ScreenOrientation) Int64() int64 { return int64(i) }

// SetInt64 sets the ScreenOrientation value from an int64.
func (i *ScreenOrientation) SetInt64(in int64) { *i = ScreenOrientation(in) }

// Desc returns the description of the ScreenOrientation value.
func (i ScreenOrientation) Desc() string { return enums.Desc(i, _ScreenOrientationDescMap) }

// ScreenOrientationValues returns all possible values for the type ScreenOrientation.
func ScreenOrientationValues() []ScreenOrientation { return _ScreenOrientationValues }

// Values returns all possible values for the type ScreenOrientation.
func (i ScreenOrientation) Values() []enums.Enum { return enums.Values(_ScreenOrientationValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ScreenOrientation) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ScreenOrientation) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "ScreenOrientation")
}

var _WindowFlagsValues = []WindowFlags{0, 1, 2, 3, 4, 5}

// WindowFlagsN is the highest valid value for type WindowFlags, plus one.
const WindowFlagsN WindowFlags = 6

var _WindowFlagsValueMap = map[string]WindowFlags{`Dialog`: 0, `Modal`: 1, `Tool`: 2, `Fullscreen`: 3, `Minimized`: 4, `Focused`: 5}

var _WindowFlagsDescMap = map[WindowFlags]string{0: `Dialog indicates that this is a temporary, pop-up window.`, 1: `Modal indicates that this dialog window blocks events going to other windows until it is closed.`, 2: `Tool indicates that this is a floating tool window that has minimized window decoration.`, 3: `Fullscreen indicates a window that occupies the entire screen.`, 4: `Minimized indicates a window reduced to an icon, or otherwise no longer visible or active. Otherwise, the window should be assumed to be visible.`, 5: `Focused indicates that the window has the focus.`}

var _WindowFlagsMap = map[WindowFlags]string{0: `Dialog`, 1: `Modal`, 2: `Tool`, 3: `Fullscreen`, 4: `Minimized`, 5: `Focused`}

// String returns the string representation of this WindowFlags value.
func (i WindowFlags) String() string { return enums.BitFlagString(i, _WindowFlagsValues) }

// BitIndexString returns the string representation of this WindowFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i WindowFlags) BitIndexString() string { return enums.String(i, _WindowFlagsMap) }

// SetString sets the WindowFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *WindowFlags) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the WindowFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *WindowFlags) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _WindowFlagsValueMap, "WindowFlags")
}

// Int64 returns the WindowFlags value as an int64.
func (i WindowFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the WindowFlags value from an int64.
func (i *WindowFlags) SetInt64(in int64) { *i = WindowFlags(in) }

// Desc returns the description of the WindowFlags value.
func (i WindowFlags) Desc() string { return enums.Desc(i, _WindowFlagsDescMap) }

// WindowFlagsValues returns all possible values for the type WindowFlags.
func WindowFlagsValues() []WindowFlags { return _WindowFlagsValues }

// Values returns all possible values for the type WindowFlags.
func (i WindowFlags) Values() []enums.Enum { return enums.Values(_WindowFlagsValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i WindowFlags) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(&i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *WindowFlags) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i WindowFlags) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *WindowFlags) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "WindowFlags")
}