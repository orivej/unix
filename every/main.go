package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/orivej/e"
)

var flVerbose = flag.Bool("v", false, "verbose")

func usage() {
	fmt.Println("Arguments: [flags] duration command args...")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		usage()
	}
	period, err := time.ParseDuration(args[0])
	if err != nil {
		usage()
	}
	loop(period, func() {
		cmd := exec.Command(args[1], args[2:]...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		e.Print(cmd.Run())
	})
}

func loop(period time.Duration, cb func()) {
	zero := time.Unix(0, 0)
	for {
		now := time.Since(zero)
		next := now.Round(period)
		for next < now {
			next += period
		}
		if *flVerbose {
			fmt.Println("next:", zero.Add(next))
		}
		time.Sleep(next - now)
		cb()
	}
}
