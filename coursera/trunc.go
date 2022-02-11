package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
Write a program which prompts the user to enter
a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.

runcation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.

This assignment is worth a total of 10 points:

3 points will be given if a program is written.

2 points will be given if the program compiles correctly.

3 points will be given for the first test execution,
if the program correctly prints the truncated integer when a floating point number is entered.

2 points will be given for the second test execution,
if the program correctly prints the truncated integer when another floating point number is entered.
*/
func main_() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	n := strings.Split(text, " ")[0]
	fmt.Println(strings.Split(n, ".")[0])
}
