package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (this *Animal) eat() {
	fmt.Println(this.food)
}
func (this *Animal) move() {
	fmt.Println(this.locomotion)
}
func (this *Animal) speak() {
	fmt.Println(this.noise)
}

func main5() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	names := map[string]string{}
	var command string
	for {
		fmt.Print("> ")
		fmt.Scan(&command)
		switch command {

		case "newanimal":
			var name, kind string
			fmt.Scan(&name, &kind)
			names[name] = kind
			fmt.Println("Created it!")

		case "query":
			var name, method string
			fmt.Scan(&name, &method)
			animal := animals[names[name]]
			switch method {
			case "eat":
				animal.eat()
			case "move":
				animal.move()
			case "speak":
				animal.speak()
			default:
				fmt.Println("Error: wrong request, no method", method)
			}
		default:
			fmt.Println("Unknown command", command)
		}
	}
}
