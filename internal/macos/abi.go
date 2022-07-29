//go:build darwin

package macOS

import (
	"unsafe"
	_ "unsafe"
)

// FuncPC returns the entry point for f. See comments in runtime/proc.go
// for the function of the same name.
//go:nosplit
func FuncPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}
