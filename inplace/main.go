package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dchest/uniuri"
	"github.com/orivej/e"
)

const usage = `inplace FILE CMD ARGS...

Filter FILE through CMD.  Replace FILE with output only when CMD succeeds.
`

func main() {
	if len(os.Args) < 3 {
		_, err := fmt.Fprintln(os.Stderr, usage)
		e.Exit(err)
		os.Exit(1)
	}

	var code int
	defer func() {
		os.Exit(code)
	}()

	path := os.Args[1]
	file, err := os.Open(path)
	e.Panic(err)
	defer e.CloseOrPrint(file)
	stat, err := file.Stat()
	e.Panic(err)

	tmppath := fmt.Sprintf("%s.%s~", path, uniuri.New())
	tmpfile, err := os.Create(tmppath)
	e.Panic(err)
	defer e.CloseOrPrint(tmpfile)

	done := false
	defer func() {
		if done {
			err = os.Rename(tmppath, path)
			e.Panic(err)
		} else {
			err = os.Remove(tmppath)
			e.Panic(err)
		}
	}()

	cmd := exec.Command(os.Args[2], os.Args[3:]...) // #nosec
	cmd.Stdin = file
	cmd.Stdout = tmpfile
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	switch err := err.(type) {
	case nil:
		code = 0
	case *exec.ExitError:
		code = 1
		return
	default:
		code = 2
		e.Panic(err)
	}
	err = os.Chmod(tmppath, stat.Mode())
	e.Print(err)
	err = os.Chtimes(tmppath, stat.ModTime(), stat.ModTime())
	e.Print(err)
	done = true
}
