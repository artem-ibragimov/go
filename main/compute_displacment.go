package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Let us assume the following formula for
// displacement s as a function of time t, acceleration a, initial velocity vo,
// and initial displacement so.

// s = ½ a t2 + vot + so

// Write a program which first prompts the user
// to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the
// program should compute the displacement after the entered time.

// You will need to define and use a function
// called GenDisplaceFn() which takes three float64
// arguments, acceleration a, initial velocity vo, and initial
// displacement so. GenDisplaceFn()
// should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial
// displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
// float64 argument which is the displacement travelled after time t.

// For example, let’s say that I want to assume
// the following values for acceleration, initial velocity, and initial
// displacement: a = 10, vo = 2, so = 1. I can use the
// following statement to call GenDisplaceFn() to
// generate a function fn which will compute displacement as a function of time.

// fn := GenDisplaceFn(10, 2, 1)

// Then I can use the following statement to
// print the displacement after 3 seconds.

// fmt.Println(fn(3))

// And I can use the following statement to print
// the displacement after 5 seconds.

// fmt.Println(fn(5))

// Test the program by running it twice and
// testing it with two different sets of values for acceleration, initial velocity,
// initial displacement, and time. Give 3 points if the program works correctly
// for one set of values, and give 3 more points if the program works correctly
// for the other set of values.

// Examine the code. If the code contains a
// function called GenDisplaceFn()
// which takes three float64 arguments, acceleration a, initial velocity vo,
// and initial displacement so and returns a function, then give
// another 2 points. If the function returned by GenDisplaceFn() is used to compute the time, give another 2 points
type Params struct {
	a, v0, s0, t float64
}

func main4() {
	params, err := read_params()
	if err != nil {
		fmt.Println(err)
		return
	}
	get_displace := GenDisplaceFn(params.a, params.v0, params.s0)
	fmt.Println("The displacement after the entered time: ", get_displace(params.t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func read_params() (Params, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter values for acceleration, initial velocity, and initial displacement: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return Params{}, err
	}
	params := make([]float64, 3)
	for i, s := range strings.Fields(strings.TrimSpace(input)) {
		params[i], err = strconv.ParseFloat(s, 64)
		if err != nil {
			return Params{}, err
		}
	}
	fmt.Print("Enter value for time: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		return Params{}, err
	}
	time, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return Params{}, err
	}
	return Params{a: params[0], v0: params[1], s0: params[2], t: time}, nil
}
