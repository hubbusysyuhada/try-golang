package main

import "fmt"

func main() {
	countries := [2]string{"indonesia", "singapura"} // bisa juga singapura tidak dimasukan
	// use [...]string{} instead if you still dont know how much the desired capacity/len is
	fmt.Println(countries)

	var names [3]string
	fmt.Println(names) // still empty
	fmt.Println(len(names), "<<<< names len")
	names[0] = "muhammad"
	names[1] = "hafidz"
	fmt.Println(names)
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], i, "<<< for i")
	}
}
