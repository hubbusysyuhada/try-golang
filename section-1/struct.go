// struct dalam javascript mirip dengan class

package main

import "fmt"

type Person struct {
	name, address, city string
	age                 int
	hobbies             []string
}

func main() {
	myself := Person{
		name:    "Hafidz",
		address: "Ratujaya",
		city:    "Depok",
		age:     25,
		hobbies: []string{"music", "movies"},
	}
	myself.hobbies = append(myself.hobbies, "games")
	fmt.Println(myself)
	fmt.Println(myself.name)
}
