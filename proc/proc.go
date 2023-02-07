package proc

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/pkg/errors"

	"github.com/cwdot/go-stdlib/wood"
)

type RunOpts struct {
	Dir string
}

func Run(process string, opts RunOpts, args ...string) (string, string, error) {
	cmd := exec.Command(process, args...)

	wood.Debugf("Executing proc %s: %s with %s", process, cmd, args)

	var outs bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &outs
	cmd.Stderr = &errs
	if opts.Dir != "" {
		cmd.Dir = opts.Dir
	}

	err := cmd.Run()
	if err != nil {
		_, err2 := os.Stderr.Write(outs.Bytes())
		if err2 != nil {
			wood.Fatal("Error writing to stderr", err2)
		}

		_, err2 = os.Stderr.Write(errs.Bytes())
		if err2 != nil {
			wood.Fatal("Error writing to stderr", err2)
		}
		return "", "", errors.Wrap(err, "proc call failed")
	}

	wood.Debugf("Executed proc: %s", outs.String())

	return outs.String(), errs.String(), nil
}
