package main

import (
	"fmt"
	"os"

	"github.com/Code-Hex/battery"
)

func main() {
	percent, state, err := BatteryInfo()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bar := battery.New(100)
	bar.ShowCounter = false
	bar.EnableColor = false
	bar.Showthunder = state

	bar.Set(percent).Run()
}
