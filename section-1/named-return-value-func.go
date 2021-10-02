package main

import "fmt"

func main() {
	a, b := details("hafidz", 25)
	fmt.Println(a, b)
}

func details(name string, age int) (fullName string, umur string) {
	fullName = name
	umur = fmt.Sprint(age, " tahun")
	return
}
