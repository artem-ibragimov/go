package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
Write a program which prompts the user to enter a string.
The program searches through the entered string
for the characters ‘i’, ‘a’, and ‘n’.

The program should print “Found!”
if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive,
so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Review criteria
This assignment is worth a total of 10 points:

3 points will be given if a program is written.

2 points will be given if the program compiles correctly.

3 points will be given for the first test execution,
if the program correctly prints "Found!" when a string
that contains the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is entered.

2 points will be given for the second test execution,
if the program correctly prints "Not Found!" when a string
that does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.
*/
const prefix = "i"
const content = "a"
const suffix = "n"

func main__() {
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	str := strings.TrimSpace(strings.ToLower(input))
	is_found := strings.HasPrefix(str, prefix) &&
		strings.Contains(str, content) &&
		strings.HasSuffix(str, suffix)

	if is_found {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")
}
