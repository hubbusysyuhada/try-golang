package example

import "fmt"

func init() {
	fmt.Println("init from test-package.go")
}

func Greeting() {
	fmt.Println("greeting, peasant!")
}
