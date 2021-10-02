/**
https://golang.org/pkg/sort/
sesuai namanya, untuk sorting
to use sort pkg, Len - Less - Swap functions must be defined
*/

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type PersonSlice []Person

func main() {
	arr := []Person{}
	person1 := Person{
		name: "budi",
		age:  35,
	}
	person2 := Person{
		name: "dwi",
		age:  27,
	}
	person3 := Person{
		name: "gugun",
		age:  52,
	}

	a := append(arr, person1)
	b := append(a, person2)
	c := append(b, person3)
	fmt.Println(c, "<<<< before")
	sort.Sort(PersonSlice(c)) // must be an interface invoked with the value
	fmt.Println(c, "<<<< after")
}

func (value PersonSlice) Len() int {
	return len(value)
}

func (value PersonSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

func (value PersonSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
