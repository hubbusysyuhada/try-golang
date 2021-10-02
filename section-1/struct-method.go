// mirip dengan class method

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	myself := Person{
		name: "Hafidz",
		age:  25,
	}
	myself.greet()
}

func (person Person) greet() {
	fmt.Println(fmt.Sprint("halo nama saya "+person.name+". Umur saya ", person.age, " tahun."))
}
