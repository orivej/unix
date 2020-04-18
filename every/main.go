package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/orivej/e"
)

var (
	flRound   = flag.Bool("r", false, "round time")
	flVerbose = flag.Bool("v", false, "verbose")
)

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
	origin := time.Now()
	if *flRound {
		origin = time.Unix(0, 0)
	}
	for {
		now := time.Since(origin)
		next := now.Round(period)
		for next < now {
			next += period
		}
		if *flVerbose {
			fmt.Println("next:", origin.Add(next))
		}
		time.Sleep(next - now)
		cb()
	}
}
