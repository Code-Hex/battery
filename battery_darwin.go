package battery

import macOS "github.com/Code-Hex/battery/internal/macos"

func Info() (int, int, bool, error) {
	status, err := macOS.Battery()
	if err != nil {
		return 0, 0, false, err
	}
	return status.Percentage, status.Elapsed, macOS.IOPSACPowerValue == status.State, nil
}
