package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Students will receive five points if the program sorts the integers
// and five additional points if there are four goroutines
// that each prints out a set of array elements that it is storing.

func main7() {
	var wg sync.WaitGroup

	fmt.Print("Enter integers: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
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
	fth := len(numbers) / 4
	sorted := make([]int, 0)

	var done = func(numbers []int) {
		sorted = append(sorted, numbers...)
		sort.Ints(sorted)
		wg.Done()
	}

	wg.Add(4)
	go sort_ints(numbers[:fth], &done)
	go sort_ints(numbers[fth:fth*2], &done)
	go sort_ints(numbers[fth*2:fth*3], &done)
	go sort_ints(numbers[fth*3:], &done)
	wg.Wait()

	fmt.Println(sorted)
}

func sort_ints(numbers []int, done *func([]int)) {
	fmt.Println(numbers, " will be sorted")
	result := append(make([]int, 0), numbers...)
	sort.Ints(result)
	(*done)(result)
}
