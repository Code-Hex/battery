package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func BatteryInfo() (int, bool, error) {
	f, err := os.Open("/sys/class/power_supply/BAT0/uevent")
	if err != nil {
		return 0, false, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var full, now float64
	var present bool
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), "=", 2)
		if len(tokens) != 2 {
			continue
		}
		switch tokens[0] {
		case "POWER_SUPPLY_ENERGY_FULL_DESIGN":
			full, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_ENERGY_NOW":
			now, _ = strconv.ParseFloat(tokens[1], 64)
		case "POWER_SUPPLY_PRESENT":
			present = tokens[1] == "1"
		}
	}
	return int(now/full) * 100, present, nil
}
