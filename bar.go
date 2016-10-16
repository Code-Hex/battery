package main

import (
	"fmt"
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

var DefaultRefreshRate = time.Millisecond * 200

type Bar struct {
	Gauge       []rune
	RefreshRate time.Duration
	width       int64
	currentVal  int64
	totalVal    int64
	charLen     int64
	prefix      rune
	suffix      rune
	isNotFinish bool
}

func main() {
	max := 100
	bar := New(max)
	bar.SetWidth(3)
	bar.Run()

	for i := 1; i <= max; i++ {
		bar.Set(i)
		time.Sleep(bar.RefreshRate / 4)
	}

	bar.Finish()
}

func New(total int) *Bar {
	return &Bar{
		RefreshRate: DefaultRefreshRate,
		totalVal:    int64(total),
		currentVal:  int64(-1),
		charLen:     int64(len(chars)),
		prefix:      '|',
		postfix:     '|',
		isNotFinish: true,
	}
}

func (bar *Bar) SetPrefix(char rune) {
	bar.prefix = char
}

func (bar *Bar) SetPostfix(char rune) {
	bar.postfix = char
}

func (bar *Bar) SetWidth(width int) {
	bar.width = int64(width)
	bar.Gauge = make([]rune, width, width)
}

func (bar *Bar) Set(set int) *Bar {
	return bar.Set64(int64(set))
}

func (bar *Bar) Set64(set int64) *Bar {
	atomic.StoreInt64(&bar.currentVal, set)
	return bar
}

func (bar *Bar) Increment() int {
	return bar.Add(1)
}

func (bar *Bar) Add(add int) int {
	return int(bar.Add64(int64(add)))
}

func (bar *Bar) Add64(add int64) int64 {
	return atomic.AddInt64(&bar.currentVal, add)
}

func (bar *Bar) Run() {
	go bar.writer()
}

// End print
func (bar *Bar) Finish() {
	bar.isNotFinish = false
	bar.print(atomic.LoadInt64(&bar.totalVal))
	fmt.Println()
}

func (bar *Bar) writer() {
	var c, oc int64
	for bar.isNotFinish {
		c = atomic.LoadInt64(&bar.currentVal)
		if c != oc {
			bar.print(c)
			oc = c
		}
		time.Sleep(bar.RefreshRate)
	}
}

func (bar *Bar) print(currentVal int64) {
	frac := float64(currentVal) / float64(bar.totalVal)
	barLen, fracBarLen := bar.divmod(frac)

	// append full block
	for i := 0; i < barLen; i++ {
		bar.Gauge[i] = chars[bar.charLen-1]
	}

	// append lower block
	bar.Gauge[barLen] = chars[fracBarLen]

	// padding with whitespace
	for i := barLen + 1; i < int(bar.width); i++ {
		bar.Gauge[i] = ' '
	}

	fmt.Printf("\r%3d%%%c%s%c", int(frac*100), bar.prefix, string(bar.Gauge), bar.postfix)
}

func (bar *Bar) divmod(frac float64) (int, int) {
	// Over 100%
	if frac >= 1.0 {
		return int(bar.width) - 1, int(bar.charLen) - 1
	}
	pre := int64(frac * float64(bar.width) * float64(bar.charLen))
	return int(pre / bar.charLen), int(pre % bar.charLen)
}
