package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
	array := [...]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for _, value := range array {
		fmt.Println("value =", value)
	}
}
