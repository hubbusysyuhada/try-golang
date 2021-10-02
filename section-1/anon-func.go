package main

import "fmt"

type Block func(string) bool

func main() {
	fmt.Println(register("hafidz"))
	fmt.Println(register("budi"))
}

func register(name string) string {
	blocked := func(person string) bool {
		if person == "budi" {
			return true
		}
		return false
	} // this is anonymous function

	if blocked(name) {
		return "You're blocked " + name
	}
	return "Welcome " + name
}
