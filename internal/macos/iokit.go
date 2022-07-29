//go:build darwin

package macOS

import "errors"

// IOKit linker flags for the external linker. See Issue 42459.

var IOPSCurrentCapacityKey = StringToCFString("Current Capacity")
var IOPSTimeToEmptyKey = StringToCFString("Time to Empty")
var IOPSPowerSourceStateKey = StringToCFString("Power Source State")

// Values for key IOPSPowerSourceStateKey.
var (
	// Power source is off-line or no longer connected.
	IOPSOffLineValue = "Off Line"
	// Power source is connected to external or AC power, and is not draining the internal battery.
	IOPSACPowerValue = "AC Power"
	// Power source is currently using the internal battery.
	IOPSBatteryPowerValue = "Battery Power"
)

//go:cgo_import_dynamic battery_IOPSCopyPowerSourcesInfo IOPSCopyPowerSourcesInfo "/System/Library/Frameworks/IOKit.framework/Versions/A/IOKit"

func IOPSCopyPowerSourcesInfo() (CFRef, error) {
	ret := syscall(FuncPC(battery_IOPSCopyPowerSourcesInfo_trampoline), 0, 0, 0, 0, 0, 0)
	// https://developer.apple.com/documentation/iokit/1523839-iopscopypowersourcesinfo
	// NULL if errors were encountered, a CFTypeRef otherwise.
	// Caller must CFRelease() the return value when done accessing it.
	if ret == 0 {
		return 0, errors.New("IOPSCopyPowerSourcesInfo: errors were encountered")
	}
	return CFRef(ret), nil
}
func battery_IOPSCopyPowerSourcesInfo_trampoline()

//go:cgo_import_dynamic battery_IOPSCopyPowerSourcesList IOPSCopyPowerSourcesList "/System/Library/Frameworks/IOKit.framework/Versions/A/IOKit"

func IOPSCopyPowerSourcesList(blob CFRef) (CFRef, error) {
	ret := syscall(FuncPC(battery_IOPSCopyPowerSourcesList_trampoline),
		uintptr(blob),
		0, 0, 0, 0, 0)
	// https://developer.apple.com/documentation/iokit/1523856-iopscopypowersourceslist
	// Returns NULL if errors were encountered, otherwise a CFArray of CFTypeRefs.
	// Caller must CFRelease() the returned CFArrayRef.
	if ret == 0 {
		return 0, errors.New("IOPSCopyPowerSourcesList: errors were encountered")
	}
	return CFRef(ret), nil
}
func battery_IOPSCopyPowerSourcesList_trampoline()

//go:cgo_import_dynamic battery_IOPSGetPowerSourceDescription IOPSGetPowerSourceDescription "/System/Library/Frameworks/IOKit.framework/Versions/A/IOKit"

func IOPSGetPowerSourceDescription(blob CFRef, ps CFRef) (CFRef, error) {
	ret := syscall(FuncPC(battery_IOPSGetPowerSourceDescription_trampoline),
		uintptr(blob), uintptr(ps),
		0, 0, 0, 0)
	// https://developer.apple.com/documentation/iokit/1523867-iopsgetpowersourcedescription
	// Returns NULL if an error was encountered, otherwise a CFDictionary.
	// Caller should NOT release the returned CFDictionary - it will be
	// released as part of the CFTypeRef returned by IOPSCopyPowerSourcesInfo().
	if ret == 0 {
		return 0, errors.New("IOPSGetPowerSourceDescription: errors were encountered")
	}
	return CFRef(ret), nil
}
func battery_IOPSGetPowerSourceDescription_trampoline()
