package main

import (
	"fmt"
	"sync"
	"time"
)

// A race condition arises when a program,
// to operate properly, depends on the sequence
// or timing of the program's processes or threads

// In case below, the final sequence will be different each time
// instead of the expected 1 2 3.
// This occurs because here the string concatenation operations
// are not mutually exclusive.
// Mutually exclusive operations are those that cannot be interrupted
// while accessing some resource such as a memory location.

func main6() {
	var wg sync.WaitGroup
	var data string
	var done = func() {
		wg.Done()
	}

	wg.Add(3)
	go log(&data, "1", &done)
	go log(&data, "2", &done)
	go log(&data, "3", &done)
	wg.Wait()
	defer fmt.Println("Expected sequence 1 2 3 becomes ", data)
}

func log(source *string, msg string, done *func()) {
	time.Sleep(time.Microsecond)
	*source = *source + msg + " "
	(*done)()
}
