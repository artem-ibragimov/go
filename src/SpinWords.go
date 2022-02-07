package main

import "strings"

// Stop gninnipS My sdroW!
// https://www.codewars.com/kata/5264d2b162488dc400000001/train/go

// Write a function that takes in a string of one or more words,
// and returns the same string,
// but with all five or more letter words reversed
// (Just like the name of this Kata).

// Strings passed in will consist of only letters and spaces.
// Spaces will be included only when more than one word is present.

// Examples:
// spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
// spinWords( "This is a test") => returns "This is a test"
// spinWords( "This is another test" )=> returns "This is rehtona test"

func SpinWords(str string) string {
	var result []string
	for _, w := range strings.Fields(str) {
		if len(w) > 5 {
			result = append(result, reverse(w))
		} else {
			result = append(result, w)
		}
	}
	return strings.Join(result, " ")
} // SpinWords

func reverse(s string) string {
	runs := []rune(s)
	len := len(runs)
	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]
	}
	return string(runs)
}
