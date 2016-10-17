package main

import "github.com/pkg/errors"

type ignore struct {
	err error
}

type cause interface {
	Cause() error
}

// errmsg method will get important message from wrapped error message
func errmsg(err error) error {
	for e := err; e != nil; {
		switch e.(type) {
		case ignore:
			return nil
		case cause:
			e = e.(cause).Cause()
		default:
			return e
		}
	}

	return nil
}

func makeIgnoreErr() ignore {
	return ignore{
		err: errors.New(""),
	}
}

// Error due to options: version, usage
func (i ignore) Error() string {
	return i.err.Error()
}
