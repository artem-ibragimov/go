package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Write a program which prompts the user to enter integers
// and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create
// an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer
// to be added to the slice.
// The program adds the integer to the slice,
// sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number
// of integers which the user decides to enter.
// The program should only quit (exiting the loop)
// when the user enters the character ‘X’ instead of an integer.

// 3 points will be given if a program is written.

// 2 points will be given if the program compiles correctly.

// A maximum of 3 points will be given for the first test execution,
// if the program correctly prints the sorted slice after entering three distinct integers.
//  **Points are awarded incrementally each time that an integer is added
// and it correctly prints the sorted slice.

// A maximum of 2 points will be given for the second test execution,
// if the program correctly prints the sorted slice after entering
// four distinct integers.
// **Points are awarded if it correctly prints the sorted slice after adding the fourth integer.

func main1() {
	reader := bufio.NewReader(os.Stdin)
	numbers := make([]int, 3)
	for {
		fmt.Print("Enter integer: ")
		str, read_err := reader.ReadString('\n')
		input := strings.TrimSpace(str)
		if input == "X" {
			fmt.Println("Done!")
			return
		}
		if read_err != nil {
			fmt.Println(read_err)
			continue
		}
		n, parse_err := strconv.Atoi(input)
		if parse_err != nil {
			fmt.Println(parse_err)
			continue
		}
		numbers = append(numbers, n)
		sort.Ints(numbers)
		fmt.Println(numbers)
	}
}
