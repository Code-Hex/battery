package main

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32              = syscall.NewLazyDLL("kernel32")
	procGetSystemPowerStatus = modkernel32.NewProc("GetSystemPowerStatus")
)

type SYSTEM_POWER_STATUS struct {
	ACLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	Reserved1           byte
	BatteryLifeTime     uint32
	BatteryFullLifeTime uint32
}

func BatteryInfo() (int, bool, error) {
	var sps SYSTEM_POWER_STATUS
	_, r1, err := procGetSystemPowerStatus.Call(uintptr(unsafe.Pointer(&sps)))
	if r1 != 0 {
		if err != nil {
			return 0, false, err
		}
	}
	return int(sps.BatteryLifePercent), sps.ACLineStatus == 1, nil
}
