package main

import (
	"fmt"
	"testing"
)

func TestAbbrevName(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"Sam Harris", "S.H"},
		{"patrick feeney", "P.F"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf(tt.in, tt.out)
		t.Run(testname, func(t *testing.T) {
			name := AbbrevName(tt.in)
			if name != tt.out {
				t.Errorf("got %s, expext %s", name, tt.out)
			}
		})
	}
}
