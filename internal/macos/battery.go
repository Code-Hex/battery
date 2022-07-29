//go:build darwin

package macOS

import (
	"errors"
)

type Status struct {
	Percentage int
	Elapsed    int
	State      string
}

func Battery() (*Status, error) {
	powerInfo, err := IOPSCopyPowerSourcesInfo()
	if err != nil {
		return nil, err
	}
	defer CFRelease(powerInfo)

	powerSrcList, err := IOPSCopyPowerSourcesList(powerInfo)
	if err != nil {
		return nil, err
	}
	defer CFRelease(powerSrcList)

	powerSrcNum := CFArrayGetCount(powerSrcList)
	if powerSrcNum == 0 {
		return nil, errors.New("could not get power resource infomation")
	}

	powerSrcInfo, err := IOPSGetPowerSourceDescription(
		powerInfo,
		CFArrayGetValueAtIndex(powerSrcList, 0),
	)
	if err != nil {
		return nil, err
	}

	percentageRef, ok := CFDictionaryGetValueIfPresent(powerSrcInfo, IOPSCurrentCapacityKey)
	if !ok {
		return nil, errors.New("not found IOPSCurrentCapacityKey")
	}
	percentage, err := CFNumberGetValue(percentageRef)
	if err != nil {
		return nil, err
	}

	elapsedRef, ok := CFDictionaryGetValueIfPresent(powerSrcInfo, IOPSTimeToEmptyKey)
	if !ok {
		return nil, errors.New("not found IOPSTimeToEmptyKey")
	}
	elapsed, err := CFNumberGetValue(elapsedRef)
	if err != nil {
		return nil, err
	}

	powerStateRef, ok := CFDictionaryGetValueIfPresent(powerSrcInfo, IOPSPowerSourceStateKey)
	if !ok {
		return nil, errors.New("not found IOPSPowerSourceStateKey")
	}
	powerState := CFStringToString(powerStateRef)

	return &Status{
		Percentage: int(percentage),
		Elapsed:    int(elapsed),
		State:      powerState,
	}, nil
}
