package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to first enter a name,
// and then enter an address.
// Your program should create a map and add the name and address to the map
// using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object
// from the map, and then your program should print the JSON object.
// 3 points will be given if a program is written.

// 2 points will be given if the program compiles correctly.

// 5 points will be given if the program correctly prints
// a JSON object with keys ("name", "address") and they are associated
// with the name and address that was entered.

func mains() {
	data := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	data["name"] = read("name", reader)
	data["address"] = read("address", reader)
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(json))
}

func read(prop string, reader *bufio.Reader) string {
	fmt.Print("Enter your " + prop + ": ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(name)
}
