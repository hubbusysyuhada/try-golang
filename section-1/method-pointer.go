package main

import "fmt"

type Person struct {
	name string
}

func main() {
	hafidz := Person{"Hafidz"}
	fmt.Println(hafidz, "<<<< person di main")
	hafidz.gender("laki")
	fmt.Println(hafidz, "<<<< person di main after")

}

func (person *Person) gender(a string) { // this will make method passed by value
	if a == "laki" {
		person.name = "Mr. " + person.name
	} else {
		person.name = "Mrs. " + person.name
	}
	fmt.Println(person, "<<<< method")
}
