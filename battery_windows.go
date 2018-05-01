package battery

import (
	"math"
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

func Info() (int, int, bool, error) {
	var sps SYSTEM_POWER_STATUS
	_, r1, err := procGetSystemPowerStatus.Call(uintptr(unsafe.Pointer(&sps)))
	if r1 != 0 {
		if err != nil {
			return 0, 0, false, err
		}
	}
	percent := int(sps.BatteryLifePercent)
	var elapsed int
	// BatteryLifeTime has MaxUint32 (2^32-1) when it cannot be detected.
	if sps.BatteryLifeTime != math.MaxUint32 {
		elapsed = int(float64(sps.BatteryLifeTime) / 60)
	}
	return percent, elapsed, sps.ACLineStatus == 1, nil
}
