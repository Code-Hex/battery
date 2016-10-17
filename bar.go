package main

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
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

var DefaultRefreshRate = time.Millisecond * 100

type Bar struct {
	Gauge       []rune
	GaugeWidth  int
	RefreshRate time.Duration
	Out         io.Writer
	width       int
	nowVal      int64
	totalVal    int64
	charLen     int64
	format      string
	prefix      rune
	postfix     rune
	isNotFinish bool
	showPercent bool
}

func main() {
	max := 100
	bar := New(max)
	bar.SetWidth(3)
	bar.Run()

	for i := 1; i <= max; i++ {
		bar.Increment()
		time.Sleep(bar.RefreshRate / 4)
	}

	bar.Finish()
}

func New(total int) *Bar {
	return &Bar{
		Out:         os.Stdout,
		RefreshRate: DefaultRefreshRate,
		totalVal:    int64(total),
		nowVal:      int64(-1),
		charLen:     int64(len(chars)),
		format:      "%s",
		prefix:      '|',
		postfix:     '|',
		isNotFinish: true,
		showPercent: true,
	}
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
	return bar.Set64(int64(set))
}

func (bar *Bar) Set64(set int64) *Bar {
	atomic.StoreInt64(&bar.nowVal, set)
	return bar
}

func (bar *Bar) Increment() int {
	return bar.Add(1)
}

func (bar *Bar) Add(add int) int {
	return int(bar.Add64(int64(add)))
}

func (bar *Bar) Add64(add int64) int64 {
	return atomic.AddInt64(&bar.nowVal, add)
}

func (bar *Bar) Run() {
	go bar.writer()
}

// End print
func (bar *Bar) Finish() {
	bar.isNotFinish = false
	bar.write(atomic.LoadInt64(&bar.totalVal))
	fmt.Println()
}

func (bar *Bar) writer() {
	if bar.showPercent {
		bar.format = "%3d%%" + bar.format
	}

	bar.format = "\r" + bar.format

	var load, oldload int64
	for bar.isNotFinish {
		load = atomic.LoadInt64(&bar.nowVal)
		if load != oldload {
			bar.write(load)
			oldload = load
		}
		time.Sleep(bar.RefreshRate)
	}
}

func (bar *Bar) write(nowVal int64) {
	frac := float64(nowVal) / float64(bar.totalVal)
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

	fmt.Fprintf(bar.Out, bar.format, int(frac*100), string(bar.Gauge))
}

func (bar *Bar) divmod(frac float64) (int, int) {
	// Over 100%
	if frac >= 1.0 {
		return bar.width, int(bar.charLen) - 1
	}
	pre := int64(frac * float64(bar.width) * float64(bar.charLen))
	return int(pre/bar.charLen) + 1, int(pre % bar.charLen)
}
