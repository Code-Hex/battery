package battery

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// LENGTH: 8
var chars = []rune{
	'▏',
	'▎',
	'▍',
	'▌',
	'▋',
	'▊',
	'▉',
	'█',
}

type Bar struct {
	Gauge       []rune
	GaugeWidth  int
	width       int
	nowVal      int
	totalVal    int
	charLen     int
	format      string
	prefix      rune
	postfix     rune
	charge      string
	ShowPercent bool
	ShowCounter bool
	Showthunder bool
	EnableColor bool
	EnableTmux  bool
}

func digit(num int) string {
	return strconv.Itoa(int(math.Log10(float64(num))) + 1)
}

func New(total int) *Bar {
	if total <= 0 {
		panic(errors.New("Please specify total size that is greater than zero"))
	}
	bar := &Bar{

		totalVal:    total,
		nowVal:      -1,
		charLen:     len(chars),
		format:      "%s%s",
		prefix:      '|',
		postfix:     '|',
		charge:      "⚡️",
		ShowPercent: true,
		ShowCounter: true,
		Showthunder: false,
		EnableColor: false,
		EnableTmux:  true,
	}
	return bar.SetWidth(3)
}

func (bar *Bar) SetPrefix(char rune) *Bar {
	bar.prefix = char
	return bar
}

func (bar *Bar) SetPostfix(char rune) *Bar {
	bar.postfix = char
	return bar
}

func (bar *Bar) SetWidth(width int) *Bar {
	bar.width = width
	// +1 for postfix
	bar.GaugeWidth = width + 1
	// +1 for prefix
	bar.Gauge = make([]rune, bar.GaugeWidth+1, bar.GaugeWidth+1)
	return bar
}

func (bar *Bar) Set(set int) *Bar {
	bar.nowVal = set
	return bar
}

func (bar *Bar) Run() {
	bar.writer()
}

func (bar *Bar) writer() {
	if bar.ShowPercent {
		bar.format = "%3d%%" + bar.format
	}

	if bar.ShowCounter {
		digit := digit(bar.totalVal)
		bar.format += " %" + digit + "d/%" + digit + "d"
	}

	bar.format = "\r" + bar.format + "\n"

	if bar.nowVal <= bar.totalVal {
		bar.print()
	}

}

func (bar *Bar) print() {
	frac := float64(bar.nowVal) / float64(bar.totalVal)
	barLen, fracBarLen := bar.divmod(frac)

	// append prefix
	bar.Gauge[0] = bar.prefix

	// append full block
	for i := 1; i < barLen; i++ {
		bar.Gauge[i] = chars[bar.charLen-1]
	}

	// append lower block
	bar.Gauge[barLen] = chars[fracBarLen]

	// padding with whitespace
	for i := barLen + 1; i < bar.GaugeWidth; i++ {
		bar.Gauge[i] = ' '
	}

	// append postfix
	bar.Gauge[bar.GaugeWidth] = bar.postfix

	bar.write(frac)
}

func (bar *Bar) write(frac float64) {
	var args []interface{}
	percent := int(frac * 100)

	if bar.ShowPercent {
		args = append(args, percent)
	}

	args = append(args, string(bar.Gauge))

	if bar.Showthunder {
		args = append(args, bar.charge)
	} else {
		args = append(args, "  ")
	}

	if bar.ShowCounter {
		args = append(args, bar.nowVal)
		args = append(args, bar.totalVal)
	}

	if bar.EnableColor {
		if bar.EnableTmux {
			colorTmuxPrint(percent, args...)
		} else {
			colorPrint(percent, args...)
		}
	} else {
		fmt.Fprintf(os.Stderr, bar.format, args...)
	}
}

func (bar *Bar) colorTmuxPrint(percent int, args ...interface{}) {
	if percent >= 60 {
		bar.format = "#[fg=1;32]" + bar.format + "#[default]"
	} else if 20 <= percent && percent < 60 {
		bar.format = "#[fg=1;33]" + bar.format + "#[default]"
	} else {
		bar.format = "#[fg=0;31]" + bar.format + "#[default]"
	}
	fmt.Fprintf(os.Stderr, bar.format, args...)
}

func (bar *Bar) colorPrint(percent int, args ...interface{}) {
	if percent >= 60 {
		fmt.Fprintf(os.Stderr, color.GreenString(bar.format, args...))
	} else if 20 <= percent && percent < 60 {
		fmt.Fprintf(os.Stderr, color.YellowString(bar.format, args...))
	} else {
		fmt.Fprintf(os.Stderr, color.RedString(bar.format, args...))
	}
}

func (bar *Bar) divmod(frac float64) (int, int) {
	// Over 100%
	if frac >= 1.0 {
		return bar.width, bar.charLen - 1
	}
	pre := int(frac * float64(bar.width) * float64(bar.charLen))
	return pre/bar.charLen + 1, pre % bar.charLen
}
