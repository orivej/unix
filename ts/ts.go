package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flPrec := flag.Int("p", 3, "precision")
	flUTC := flag.Bool("u", false, "print UTC time")
	flag.Parse()

	width := *flPrec + 4
	lineFormat := fmt.Sprintf("%%%d.%df %%%d.%df   %%s", width, *flPrec, width, *flPrec)
	startup := time.Now()
	previous := startup
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ScanLines)
	for scanner.Scan() {
		instant := time.Now()
		s := lineFormat
		if *flUTC {
			s = instant.UTC().Format(time.RFC3339Nano) + " " + s
		}
		fmt.Printf(
			s,
			instant.Sub(previous).Seconds(),
			instant.Sub(startup).Seconds(),
			scanner.Bytes(),
		)
		previous = instant
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
	}
}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[:i+1], nil
	}
	return 0, nil, nil
}
