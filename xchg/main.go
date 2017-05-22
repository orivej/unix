package main

import (
	"fmt"
	"os"

	"github.com/orivej/e"
	xchg "github.com/williamsandrew/go-xchg"
)

var help = `Arguments: PATH1 PATH2
Atomically exchange paths (if supported).

Supported in:
- ext4 and fuse since linux 3.15
- tmpfs since linux 3.17
`

func main() {
	if len(os.Args) < 3 {
		fmt.Print(help)
		os.Exit(1)
	}
	err := xchg.Exchange(os.Args[1], os.Args[2])
	e.Exit(err)
}
