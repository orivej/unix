package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/orivej/e"
)

var usage = `Arguments: [ACTION ARGS...]...

Actions:

- g PATTERN

  Print PATTERN regexp matches.

- s PATTERN REPLACEMENT

  Replace PATTERN regexp matches with REPLACEMENTs.

- gs PATTERN REPLACEMENT

  Print REPLACEMENTs of PATTERN regexp matches.

- r OLD NEW

  Replace literal OLD with NEW.
`

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	actions, err := parseActions(args)
	if err != nil {
		log.Println(usage)
		log.Fatal(err)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	e.Exit(err)

	output := performActions(actions, input)

	_, err = os.Stdout.Write(output)
	e.Exit(err)
}

func parseActions(args []string) (actions []Action, err error) {
	var action Action

	next := func(n int) []string {
		ret := args[:n]
		args = args[n:]
		return ret
	}

	for len(args) > 0 {
		command := next(1)[0]
		def, ok := ActionDefs[command]
		if !ok {
			err = fmt.Errorf("Unknown action %q", command)
			return
		}
		if len(args) < def.NArgs {
			err = fmt.Errorf("%q requires %d arguments", command, def.NArgs)
			return
		}

		action, err = def.Actioner(next(def.NArgs)...)
		if err != nil {
			return
		}

		actions = append(actions, action)
	}
	return
}

func performActions(actions []Action, data []byte) []byte {
	for _, action := range actions {
		data = action(data)
	}
	return data
}
