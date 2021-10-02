package main

import "fmt"

func main() {
	checker := 5
	if checker == 1 {
		fmt.Println("checker 1")
	} else if checker == 2 {
		fmt.Println("checker 2")
	} else {
		fmt.Println("nilai baru checker", checker)
	}

	// short statement dalam if
	name := "hafidz"
	if len := len(name); len > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("hello", name)
	}
}
