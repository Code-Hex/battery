package battery

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Info() (int, int, bool, error) {
	f, err := os.Open("/sys/class/power_supply/BAT0/uevent")
	if err != nil {
		return 0, 0, false, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var full, now, powerNow float64
	var present bool
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
	var percent, elapsed int
	if full > 0 {
		percent = int(now / full * 100)
	}
	if powerNow > 0 {
		elapsed = int(now / powerNow * 60)
	}
	return percent, elapsed, present, nil
}
