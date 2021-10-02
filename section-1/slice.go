package main

import "fmt"

func main() {
	countries := [5]string{
		"indonesia",
		"singapura",
		"vietnam",
		"malaysia",
		"thailand",
	}

	// use [...]string{} instead if you still dont know how much the desired capacity/len is

	newarr := countries[2:]   // vietnam untill end
	newarr2 := countries[:2]  // start untill before vietnam (2)
	newarr3 := countries[1:4] // singapura untill before thailand (4)
	fmt.Println(newarr)
	fmt.Println(newarr2)
	fmt.Println(newarr3)
	fmt.Println(len(newarr2))                              // get the slice length
	fmt.Println(cap(newarr2))                              // get the slice capacity (from start of slice to the end of original array)
	appendedCountries := append(countries[:], "australia") // append countries, could be used even if higher than the original capacity (create new array)
	fmt.Println(append(appendedCountries[:], "jepang"))
	createdSlice := make([]string, 2, 5) // make new slice with length = 2 and capacity = 5
	createdSlice[0] = "koria"
	createdSlice[1] = "china"
	fmt.Println(createdSlice)
	copiedSlice := make([]string, len(createdSlice), cap(createdSlice))
	fmt.Println(copiedSlice, "<<< before")
	copy(copiedSlice, createdSlice) // make a copy of an array/slice into destination
	fmt.Println(copiedSlice, "<<< after")

}
