package battery

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
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
	Out         io.Writer
	width       int
	nowVal      int
	totalVal    int
	charLen     int
	format      string
	prefix      rune
	postfix     rune
	ShowPercent bool
	ShowCounter bool
}

func digit(num int) string {
	return strconv.Itoa(int(math.Log10(float64(num))) + 1)
}

func New(total int) *Bar {
	if total <= 0 {
		panic(errors.New("Please specify total size that is greater than zero"))
	}
	bar := &Bar{
		Out:         os.Stdout,
		totalVal:    total,
		nowVal:      -1,
		charLen:     len(chars),
		format:      "%s",
		prefix:      '|',
		postfix:     '|',
		ShowPercent: true,
		ShowCounter: true,
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

	bar.format = "\r" + bar.format

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

	if bar.ShowPercent {
		args = append(args, int(frac*100))
	}

	args = append(args, string(bar.Gauge))

	if bar.ShowCounter {
		args = append(args, bar.nowVal)
		args = append(args, bar.totalVal)
	}
	fmt.Fprintf(bar.Out, bar.format, args...)
}

func (bar *Bar) divmod(frac float64) (int, int) {
	// Over 100%
	if frac >= 1.0 {
		return bar.width, bar.charLen - 1
	}
	pre := int(frac * float64(bar.width) * float64(bar.charLen))
	return pre/bar.charLen + 1, pre % bar.charLen
}
