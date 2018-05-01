package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Code-Hex/battery"
	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

const (
	version = "0.2.0"
	msg     = "battery v" + version + "\n"
)

// Options struct for parse command line arguments
type Options struct {
	Help    bool `short:"h" long:"help"`
	Tmux    bool `short:"t" long:"tmux"`
	Has     bool `long:"has"`
	Elapsed bool `short:"e" long:"elapsed"`
	Version bool `short:"v" long:"version"`
}

func main() {
	var opts Options
	parseOptions(&opts, os.Args[1:])
	percent, elapsed, state, err := battery.Info()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bar := battery.New(100)
	bar.EnableTmux = opts.Tmux
	bar.ShowCounter = false
	bar.EnableColor = true
	bar.Showthunder = state
	bar.ShowElapsed = opts.Elapsed

	bar.Set(percent, elapsed).Run()
}

func hasBattery() int {
	if _, _, _, err := battery.Info(); err != nil {
		return 1
	}
	return 0
}

func parseOptions(opts *Options, argv []string) {

	if _, err := opts.parse(argv); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if opts.Help {
		os.Stdout.Write(opts.usage())
		os.Exit(0)
	}

	if opts.Version {
		fmt.Fprintf(os.Stdout, msg)
		os.Exit(0)
	}

	if opts.Has {
		// If your device have the battery, exit code is 0
		os.Exit(hasBattery())
	}
}

func (opts *Options) parse(argv []string) ([]string, error) {
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.ParseArgs(argv)

	if err != nil {
		os.Stderr.Write(opts.usage())
		return nil, errors.Wrap(err, "invalid command line options")
	}

	return args, nil
}

func (opts Options) usage() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintf(&buf, msg+
		`Usage: battery [options]
  Options:
  -h,  --help        print usage and exit
  -v,  --version     display the version of battery and exit
  -t,  --tmux        display battery ascii art on tmux
  -e,  --elapsed     display the elapsed time to charge / discharge
       --has         check to see if your device have the battery
`)
	return buf.Bytes()
}
