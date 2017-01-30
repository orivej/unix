package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	startup := time.Now()
	previous := startup
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ScanLines)
	for scanner.Scan() {
		instant := time.Now()
		fmt.Printf(
			"%7.3f %7.3f   %s",
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
