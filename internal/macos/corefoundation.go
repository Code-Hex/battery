//go:build darwin

package macOS

import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
)

// Core Foundation linker flags for the external linker. See Issue 42459 in golang/go.

// CFRef is an opaque reference to a Core Foundation object. It is a pointer,
// but to memory not owned by Go, so not an unsafe.Pointer.
type CFRef uintptr

// syscall is implemented in the runtime package (runtime/sys_darwin.go)

//go:linkname syscall crypto/x509/internal/macos.syscall
func syscall(fn, a1, a2, a3, a4, a5 uintptr, f1 float64) uintptr

// CFStringToString returns a Go string representation of the passed
// in CFString, or an empty string if it's invalid.
func CFStringToString(ref CFRef) string {
	data, err := CFStringCreateExternalRepresentation(ref)
	if err != nil {
		return ""
	}
	b := CFDataToSlice(data)
	CFRelease(data)
	return string(b)
}

const kCFNumberIntType = 9

type CFString CFRef

const kCFAllocatorDefault = 0
const kCFStringEncodingUTF8 = 0x08000100

//go:cgo_import_dynamic battery_CFStringCreateExternalRepresentation CFStringCreateExternalRepresentation "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFStringCreateExternalRepresentation(strRef CFRef) (CFRef, error) {
	ret := syscall(FuncPC(battery_CFStringCreateExternalRepresentation_trampoline), kCFAllocatorDefault, uintptr(strRef), kCFStringEncodingUTF8, 0, 0, 0)
	if ret == 0 {
		return 0, errors.New("string can't be represented as UTF-8")
	}
	return CFRef(ret), nil
}
func battery_CFStringCreateExternalRepresentation_trampoline()

func CFDataToSlice(data CFRef) []byte {
	length := CFDataGetLength(data)
	ptr := CFDataGetBytePtr(data)
	src := (*[1 << 20]byte)(unsafe.Pointer(ptr))[:length:length]
	out := make([]byte, length)
	copy(out, src)
	return out
}

//go:cgo_import_dynamic battery_CFDataGetLength CFDataGetLength "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDataGetLength(data CFRef) int {
	ret := syscall(FuncPC(battery_CFDataGetLength_trampoline), uintptr(data), 0, 0, 0, 0, 0)
	return int(ret)
}
func battery_CFDataGetLength_trampoline()

//go:cgo_import_dynamic battery_CFDataGetBytePtr CFDataGetBytePtr "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDataGetBytePtr(data CFRef) uintptr {
	ret := syscall(FuncPC(battery_CFDataGetBytePtr_trampoline), uintptr(data), 0, 0, 0, 0, 0)
	return ret
}
func battery_CFDataGetBytePtr_trampoline()

//go:cgo_import_dynamic battery_CFRelease CFRelease "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFRelease(ref CFRef) {
	syscall(FuncPC(battery_CFRelease_trampoline), uintptr(ref), 0, 0, 0, 0, 0)
}
func battery_CFRelease_trampoline()

//go:cgo_import_dynamic battery_CFArrayGetCount CFArrayGetCount "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayGetCount(array CFRef) int {
	ret := syscall(FuncPC(battery_CFArrayGetCount_trampoline), uintptr(array), 0, 0, 0, 0, 0)
	return int(ret)
}
func battery_CFArrayGetCount_trampoline()

const kCFNumberSInt32Type = 3

//go:cgo_import_dynamic battery_CFNumberGetValue CFNumberGetValue "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFNumberGetValue(num CFRef) (int32, error) {
	var value int32
	ret := syscall(FuncPC(battery_CFNumberGetValue_trampoline), uintptr(num), uintptr(kCFNumberSInt32Type),
		uintptr(unsafe.Pointer(&value)), 0, 0, 0)
	if ret == 0 {
		return 0, errors.New("CFNumberGetValue call failed")
	}
	return value, nil
}
func battery_CFNumberGetValue_trampoline()

//go:cgo_import_dynamic battery_CFDictionaryGetValueIfPresent CFDictionaryGetValueIfPresent "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDictionaryGetValueIfPresent(dict CFRef, key CFString) (value CFRef, ok bool) {
	ret := syscall(FuncPC(battery_CFDictionaryGetValueIfPresent_trampoline), uintptr(dict), uintptr(key),
		uintptr(unsafe.Pointer(&value)), 0, 0, 0)
	if ret == 0 {
		return 0, false
	}
	return value, true
}
func battery_CFDictionaryGetValueIfPresent_trampoline()

//go:cgo_import_dynamic battery_CFStringCreateWithBytes CFStringCreateWithBytes "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

// StringToCFString returns a copy of the UTF-8 contents of s as a new CFString.
func StringToCFString(s string) CFString {
	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	ret := syscall(FuncPC(battery_CFStringCreateWithBytes_trampoline), kCFAllocatorDefault, uintptr(p),
		uintptr(len(s)), uintptr(kCFStringEncodingUTF8), 0 /* isExternalRepresentation */, 0)
	runtime.KeepAlive(p)
	return CFString(ret)
}
func battery_CFStringCreateWithBytes_trampoline()

//go:cgo_import_dynamic battery_CFArrayGetValueAtIndex CFArrayGetValueAtIndex "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayGetValueAtIndex(array CFRef, index int) CFRef {
	ret := syscall(FuncPC(battery_CFArrayGetValueAtIndex_trampoline), uintptr(array), uintptr(index), 0, 0, 0, 0)
	return CFRef(ret)
}
func battery_CFArrayGetValueAtIndex_trampoline()

// ReleaseCFArray iterates through an array, releasing its contents, and then
// releases the array itself. This is necessary because we cannot, easily, set the
// CFArrayCallBacks argument when creating CFArrays.
func ReleaseCFArray(array CFRef) {
	for i := 0; i < CFArrayGetCount(array); i++ {
		ref := CFArrayGetValueAtIndex(array, i)
		CFRelease(ref)
	}
	CFRelease(array)
}
