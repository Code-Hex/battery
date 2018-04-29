package battery

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

void battery(int *percentage, int *elapsed, char **status, char **error) {

	CFTypeRef powerInfo = IOPSCopyPowerSourcesInfo();
	CFArrayRef powerSrcList = IOPSCopyPowerSourcesList(powerInfo);
	CFDictionaryRef powerSrcInfo = NULL;

	if (!powerSrcList) {
		if (powerInfo) CFRelease(powerInfo);
		setStrValue(error, "Failed to get value from IOPSCopyPowerSourcesList()");
		return;
	}

	const void *powerSrcVal = NULL;
	const char *powerStatus = NULL;
	if (CFArrayGetCount(powerSrcList)) {
		powerSrcInfo = IOPSGetPowerSourceDescription(powerInfo, CFArrayGetValueAtIndex(powerSrcList, 0));
		powerSrcVal = CFDictionaryGetValue(powerSrcInfo, CFSTR(kIOPSCurrentCapacityKey));
		CFNumberGetValue((CFNumberRef)powerSrcVal, kCFNumberIntType, percentage);

		powerSrcVal = CFDictionaryGetValue(powerSrcInfo, CFSTR(kIOPSTimeToEmptyKey));
		CFNumberGetValue((CFNumberRef)powerSrcVal, kCFNumberIntType, elapsed);

		powerSrcVal = CFDictionaryGetValue(powerSrcInfo, CFSTR(kIOPSPowerSourceStateKey));
		powerStatus = CFStringGetCStringPtr((CFStringRef)powerSrcVal, kCFStringEncodingUTF8);
		setStrValue(status, powerStatus);
	} else {
		setStrValue(error, "Could not get power resource infomation");
		return;
	}

    if (powerInfo) CFRelease(powerInfo);
    if (powerSrcList) CFRelease(powerSrcList);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Info() (int, int, bool, error) {
	var percent C.int
	var elapsed C.int
	var status *C.char
	var err *C.char
	defer C.free(unsafe.Pointer(status))
	defer C.free(unsafe.Pointer(err))

	C.battery(&percent, &elapsed, &status, &err)
	p := int(percent)
	e := int(elapsed)
	if p == -1 {
		return p, e, false, errors.New(C.GoString(err))
	}

	return p, e, "AC Power" == C.GoString(status), nil
}
