package main

import "fmt"

func main() {
	countries := [...]string{
		"indonesia",
		"singapura",
		"amerika",
		"afrika",
	}

	sliced_array := countries[0:]
	new_sliced := append(sliced_array, "jerman")
	new_sliced = append(new_sliced, "venezuela")
	fmt.Println(new_sliced)

	new_array := make([]string, 2, 10)
	new_array[0] = "bola"
	new_array[1] = "kaos"
	fmt.Println(new_array, "<<< new array")
}
