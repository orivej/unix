package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	tests := [][4]string{
		{".+", "\n", ""},
		{".*", "\n", "\n\n"},
		{".+", "\na\n\nb", "a\nb\n"},
	}
	for _, test := range tests {
		action, err := grep(test[0])
		if !assert.NoError(t, err) {
			return
		}
		output := action([]byte(test[1]))
		assert.Equal(t, test[2], string(output))
	}

}

func TestSed(t *testing.T) {
	tests := [][4]string{
		{".+", ".", "a\nb\n", ".\n.\n"},
	}
	for _, test := range tests {
		action, err := sed(test[0], test[1])
		if !assert.NoError(t, err) {
			return
		}
		output := action([]byte(test[2]))
		assert.Equal(t, test[3], string(output))
	}

}

func TestGrepSed(t *testing.T) {
	tests := [][4]string{
		{"(?P<first>.).*(?P<last>.)", "$last", "a\nab\nabc\n", "b\nc\n"},
	}
	for _, test := range tests {
		action, err := grepsed(test[0], test[1])
		if !assert.NoError(t, err) {
			return
		}
		output := action([]byte(test[2]))
		assert.Equal(t, test[3], string(output))
	}

}
