package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

// Test the program by running it twice and
// testing it with a different sequence of integers each time. The first test
// sequence of integers should be all positive numbers and the second test should
// have at least one negative number. Give 3 points if the program works correctly
// for one test sequence, and give 3 more points if the program works correctly
// for the second test sequence.

// Examine the code. If the code contains a function
// called BubbleSort() which takes a slice of integers as an argument, then
// give another 2 points.If the code
// contains a function called Swap() function which takes two arguments, a slice of
// integers and an index value i, then give another 2 points.
func main3() {
	fmt.Print("Enter 10 intergers: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	txt := strings.Fields(strings.TrimSpace(input))
	numbers := make([]int, len(txt))
	for i, s := range txt {
		numbers[i], err = strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}

func BubbleSort(nums []int) {
	for i := range nums {
		for j := range nums[:len(nums)-i-1] {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
