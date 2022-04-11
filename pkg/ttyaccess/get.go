package ttyaccess

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-tty"
)

// Get returns io.Writer to TTY or an error.
//
// Writing directly to TTY can be useful when running your application as
// subprocess and the parent process does not pipe the subprocess stderr into
// end-user terminal.
//
// This happens often with various AWS tools when using credential_process.
func Get() (io.Writer, error) {
	var out io.Writer

	tty, err := tty.Open()

	if err != nil {
		return nil, err
	}

	defer tty.Close()
	out = colorable.NewColorable(tty.Output())

	return out, nil
}

// GetWithFallback returns either io.Writer to TTY or if it could not be
// accessed, process standard error stream (stderr) is returned.
//
// Writing directly to TTY can be useful when running your application as
// subprocess and the parent process does not pipe the subprocess stderr into
// end-user terminal.
//
// This happens often with various AWS tools when using credential_process.
func GetWithFallback() io.Writer {
	out, err := Get()
	if err != nil {
		return os.Stderr
	}

	return out
}
