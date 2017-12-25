package main

import (
	"bytes"
	"regexp"
)

type Action func([]byte) []byte

type Actioner func(...string) (Action, error)

type ActionDef struct {
	NArgs    int
	Actioner Actioner
}

var ActionDefs = map[string]ActionDef{
	"g":  {1, grep},
	"s":  {2, sed},
	"gs": {2, grepsed},
	"r":  {2, replace},
}

func withPattern(pattern string, f func(*regexp.Regexp, []byte) []byte) (Action, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return func(b []byte) []byte {
		return f(rx, b)
	}, nil
}

func grep(args ...string) (Action, error) {
	return withPattern(args[0], func(rx *regexp.Regexp, b []byte) []byte {
		matches := rx.FindAll(b, -1)
		if len(matches) > 0 {
			matches = append(matches, []byte{})
		}
		return bytes.Join(matches, []byte{'\n'})
	})
}

func sed(args ...string) (Action, error) {
	replacement := []byte(args[1])
	return withPattern(args[0], func(rx *regexp.Regexp, b []byte) []byte {
		return rx.ReplaceAll(b, replacement)
	})
}

func grepsed(args ...string) (Action, error) {
	replacement := []byte(args[1] + "\n")
	return withPattern(args[0], func(rx *regexp.Regexp, b []byte) []byte {
		var output []byte
		for _, match := range rx.FindAllSubmatchIndex(b, -1) {
			output = rx.Expand(output, replacement, b, match)
		}
		return output
	})
}

func replace(args ...string) (Action, error) {
	old, new := []byte(args[0]), []byte(args[1])
	return func(b []byte) []byte {
		return bytes.Replace(b, old, new, -1)
	}, nil
}
