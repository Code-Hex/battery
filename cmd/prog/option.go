package main

import (
	"bytes"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type Options struct {
	Help           bool `short:"h" long:"help"`
	Version        bool `short:"v" long:"version"`
	Total          int  `short:"t" long:"total"`
	Width          int  `short:"w" long:"width"`
	DisablePercent bool `long:"disable-percent"`
	DisableCounter bool `long:"disable-counter"`
	Trace          bool `long:"trace"`
}

func (opts *Options) parse(argv []string) error {
	p := flags.NewParser(opts, flags.PrintErrors)
	_, err := p.ParseArgs(argv)

	if err != nil {
		os.Stderr.Write(opts.usage())
		return errors.Wrap(err, "invalid command line options")
	}

	return nil
}

func (opts Options) usage() []byte {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, msg+
		`Usage: 
  $ [commands] | prog [options]

  Options:
  -h,  --help                   print usage and exit
  -v,  --version                display the version of gpl and exit
  -t,  --total                  specify the total size
  -w,  --width                  specify width of progress bar
  --disable-percent             do not show percentage
  --disable-counter             do not show counter
  --trace                       display detail error messages
`)
	return buf.Bytes()
}
