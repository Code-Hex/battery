package battery

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Info() (percent int, elapsed int, present bool, err error) {
	var uevents []string
	uevents, err = filepath.Glob("/sys/class/power_supply/BAT*/uevent")
	if err != nil {
		return
	}
	if len(uevents) == 0 {
		return
	}
	var f *os.File
	for _, u := range uevents {
		f, err = os.Open(u)
		if err == nil {
			break
		}
	}
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var full, now, powerNow float64
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), "=", 2)
		if len(tokens) != 2 {
			continue
		}
		switch tokens[0] {
		case "POWER_SUPPLY_ENERGY_FULL_DESIGN":
			full, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_CHARGE_FULL":
			full, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_ENERGY_NOW":
			now, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_CHARGE_NOW":
			now, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_STATUS":
			present = tokens[1] == "Charging"
		case "POWER_SUPPLY_POWER_NOW":
			powerNow, _ = strconv.ParseFloat(tokens[1], 64)
		}
	}
	if full > 0 {
		percent = int(now / full * 100)
	}
	if powerNow > 0 {
		elapsed = int(now / powerNow * 60)
	}
	return
}
