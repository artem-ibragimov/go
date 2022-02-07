package main

import (
	"testing"
)

func TestSpinWords(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"This is a test", "This is a test"},
		{"This is another test", "This is rehtona test"},
	}

	for _, tt := range tests {
		t.Run(tt.in+" -> "+tt.out, func(t *testing.T) {
			words := SpinWords(tt.in)
			if words != tt.out {
				t.Errorf("got %s, expexted %s", words, tt.out)
			}
		})
	}
}
