//go:build darwin

#include "textflag.h"

// The trampolines are ABIInternal as they are address-taken in
// Go code.

TEXT ·battery_CFArrayGetCount_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFArrayGetCount(SB)
TEXT ·battery_CFArrayGetValueAtIndex_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFArrayGetValueAtIndex(SB)
TEXT ·battery_CFDataGetBytePtr_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFDataGetBytePtr(SB)
TEXT ·battery_CFDataGetLength_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFDataGetLength(SB)
TEXT ·battery_CFRelease_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFRelease(SB)
TEXT ·battery_CFDictionaryGetValueIfPresent_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFDictionaryGetValueIfPresent(SB)
TEXT ·battery_CFNumberGetValue_trampoline(SB),NOSPLIT,$0-0
	JMP	battery_CFNumberGetValue(SB)
TEXT ·battery_CFStringCreateExternalRepresentation_trampoline(SB),NOSPLIT,$0-0
	JMP battery_CFStringCreateExternalRepresentation(SB)
TEXT ·battery_CFStringCreateWithBytes_trampoline(SB),NOSPLIT,$0-0
    JMP battery_CFStringCreateWithBytes(SB)
