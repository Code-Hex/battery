//go:build darwin

#include "textflag.h"

// The trampolines are ABIInternal as they are address-taken in
// Go code.

TEXT ·battery_IOPSCopyPowerSourcesInfo_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_IOPSCopyPowerSourcesInfo(SB)
TEXT ·battery_IOPSCopyPowerSourcesList_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_IOPSCopyPowerSourcesList(SB)
TEXT ·battery_IOPSGetPowerSourceDescription_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_IOPSGetPowerSourceDescription(SB)
