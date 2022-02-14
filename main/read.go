package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which reads information
// from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order,
// separated by a single space on the line.

// Your program will define a name struct which has two fields,
// fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file
// and create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice,
// and after all lines have been read from the file,
// your program will have a slice containing one struct for each line in the file.
// After reading all lines from the file,
// your program should iterate through your slice of structs
// and print the first and last names found in each struct.
// 3 points will be given if a program is written.

// 2 points will be given if the program compiles correctly.

// 5 points will be given if test execution is successful and your program:
// 1. Opens a named text file.
// 2. Prints all first name/ last name pairs.

type ID struct {
	fname string
	lname string
}

func newID(names []string) ID {
	return ID{names[0], names[1]}
}

func main2() {
	fmt.Print("Enter the relative file path: ")
	reader := bufio.NewReader(os.Stdin)
	filepath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	data, err := os.ReadFile(strings.TrimSpace(filepath))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	rows := strings.Split(string(data), "\n")
	ids := make([]ID, len(rows))
	for i, name := range rows {
		ids[i] = newID(strings.Split(name, " "))
	}

	fmt.Println(ids)
}
