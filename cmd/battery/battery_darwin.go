package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework IOKit

#import <Foundation/Foundation.h>
#import <IOKit/ps/IOPowerSources.h>
#import <IOKit/ps/IOPSKeys.h>

void setStrValue(char **dest, const char *src) {
	int len = strlen(src) + 1;
	*dest = (char*)calloc(len, sizeof(char));
	strncpy(*dest, src, len);
}

int battery(char **status, char **error) {

	CFTypeRef powerInfo = IOPSCopyPowerSourcesInfo();
	CFArrayRef powerSrcList = IOPSCopyPowerSourcesList(powerInfo);
	CFDictionaryRef powerSrcInfo;

	if (!powerSrcList) {
		CFRelease(powerInfo);
		setStrValue(error, "Failed to get value from IOPSCopyPowerSourcesList()");
		return -1;
	}

	int percentage;
	const void *powerSrcVal;
	const void *powerStatus;
	if (CFArrayGetCount(powerSrcList)) {
		powerSrcInfo = IOPSGetPowerSourceDescription(powerInfo, CFArrayGetValueAtIndex(powerSrcList, 0));
		powerSrcVal = CFDictionaryGetValue(powerSrcInfo, CFSTR(kIOPSCurrentCapacityKey));
		CFNumberGetValue((CFNumberRef)powerSrcVal, kCFNumberIntType, &percentage);

		powerSrcVal = CFDictionaryGetValue(powerSrcInfo, CFSTR(kIOPSPowerSourceStateKey));
		powerStatus = CFStringGetCStringPtr((CFStringRef)powerSrcVal, kCFStringEncodingUTF8);
		setStrValue(status, powerStatus);

		CFRelease(powerStatus);
		CFRelease(powerSrcVal);
	} else {
		setStrValue(error, "Could not get power resource infomation");
		return -1;
	}

    CFRelease(powerInfo);
    CFRelease(powerSrcList);
    CFRelease(powerSrcInfo);

    return percentage;
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func BatteryInfo() (int, bool, error) {
	var status *C.char
	var err *C.char
	defer C.free(unsafe.Pointer(status))
	defer C.free(unsafe.Pointer(err))

	percent := int(C.battery(&status, &err))
	if percent == -1 {
		return percent, false, errors.New(C.GoString(err))
	}

	isCharged := map[string]bool{
		"AC Power":      true,
		"Battery Power": false,
	}

	return percent, isCharged[C.GoString(status)], nil
}
