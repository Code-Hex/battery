package main

import (
	"time"

	"../../../Progress"
)

func main() {
	max := 1000
	bar := Progress.New(max).SetWidth(5)
	bar.Run()

	for i := 1; i <= max; i++ {
		bar.Increment()
		time.Sleep(bar.RefreshRate / 4)
	}

	bar.Finish()
}
