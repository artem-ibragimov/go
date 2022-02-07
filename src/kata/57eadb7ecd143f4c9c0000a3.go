package main

import "strings"

// Write a function to convert a name into initials.
// This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:

// Sam Harris => S.H
// patrick feeney => P.F

func AbbrevName(name string) string {
	names := strings.Fields(name)
	return strings.ToUpper(names[0][:1]) + "." + strings.ToUpper(names[1][:1])
}
