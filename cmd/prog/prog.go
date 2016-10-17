package main

import (
	"fmt"
	"io"
	"os"

	"../../../Progress"
	"github.com/pkg/errors"
)

const (
	version = "0.0.1"
	msg     = "prog v" + version + ", Progress bar\n"
)

type Prog struct {
	Trace  bool
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func New() *Prog {
	return &Prog{
		Stdin:  os.Stdin,
		Stdout: os.Stderr,
		Stderr: os.Stderr,
	}
}

func (prog *Prog) Run() int {
	if err := prog.Ress(); err != nil {
		if prog.Trace {
			fmt.Fprintf(prog.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(prog.Stderr, "Error:\n  %v\n", errmsg(err))
		}
		return 1
	}
	return 0

}

func (prog *Prog) Ress() error {
	var opts Options
	if err := prog.ParseOptions(&opts, os.Args[1:]); err != nil {
		return errmsg(err)
	}

	bar := Progress.New(opts.Total)
	bar.Out = prog.Stdout

	if opts.DisablePercent {
		bar.ShowPercent = false
	}

	if opts.DisableCounter {
		bar.ShowCounter = false
	}

	if opts.Width > 0 {
		bar.SetWidth(opts.Width)
	}

	bar.Run()

	r := bar.NewProxyReader(prog.Stdin)
	io.Copy(os.Stdout, r)

	bar.Finish()

	return nil
}

// ParseOption for prog command line arguments.
func (prog *Prog) ParseOptions(opts *Options, argv []string) error {

	if err := opts.parse(argv); err != nil {
		return errors.Wrap(err, "failed to parse command line options")
	}

	if opts.Help {
		prog.Stdout.Write(opts.usage())
		return makeIgnoreErr()
	}

	if opts.Version {
		prog.Stdout.Write([]byte(msg))
		return makeIgnoreErr()
	}

	if opts.Trace {
		prog.Trace = opts.Trace
	}

	return nil
}
