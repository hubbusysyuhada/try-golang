package main

import (
	"fmt"
	"os" //
)

func main() {
	args := os.Args[1:]            // this is arguments in command (slice of strings) same with process.argv in javascript
	hostName, err := os.Hostname() // name of this computer
	env := os.Getenv("GO_USER")    // same with process.env in javascript
	fmt.Println("Args:", args)     // more in https://golang.org/pkg/os/
	fmt.Println("env GO_USER:", env)
	if err == nil {
		fmt.Println(hostName)
	}
}
