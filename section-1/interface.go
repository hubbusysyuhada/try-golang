/**
interface merupakan sebuah type declaration yang dapat diistilahkan sebagai sebuah kontrak
interface dapat mendeklarasikan function/method ataupun datatype primitif biasa
*/

package main

import "fmt"

type MyInterface interface {
	getAge() string
	getName() string
}

type Person struct {
	name string
	age  int
}

func main() {
	hafidz := Person{
		name: "hafidz",
		age:  25,
	}
	greet(hafidz)

}

func (p Person) getAge() string {
	return fmt.Sprintln(p.age, "years old")
}

func (p Person) getName() string {
	return fmt.Sprint("my name is ", p.name, ".")
}

func greet(in MyInterface) {
	fmt.Println("hello", in.getName(), "I am", in.getAge())
}
