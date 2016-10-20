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
	version = "0.0.1"
	msg     = "battery v" + version + "\n"
)

// Options struct for parse command line arguments
type Options struct {
	Help bool `short:"h" long:"help"`
	Tmux bool `short:"t" long:"tmux"`
}

func main() {
	var opts Options
	parseOptions(&opts, os.Args[1:])
	percent, state, err := BatteryInfo()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bar := battery.New(100)
	bar.EnableTmux = opts.Tmux
	bar.ShowCounter = false
	bar.EnableColor = false
	bar.Showthunder = state

	bar.Set(percent).Run()
}

func parseOptions(opts *Options, argv []string) {

	if len(argv) == 0 {
		os.Stdout.Write(opts.usage())
		os.Exit(0)
	}

	_, err := opts.parse(argv)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if opts.Help {
		os.Stdout.Write(opts.usage())
		os.Exit(0)
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
  -v,  --version     display the version of pget and exit
  -t,  --tmux        display battery ascii art on tmux           
`)
	return buf.Bytes()
}
