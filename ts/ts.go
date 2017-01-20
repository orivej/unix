package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startup := time.Now()
	previous := startup
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instant := time.Now()
		fmt.Printf(
			"%7.3f %7.3f   %v\n",
			instant.Sub(previous).Seconds(),
			instant.Sub(startup).Seconds(),
			scanner.Text(),
		)
		previous = instant
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
	}
}
