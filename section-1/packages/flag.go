// parsing command line argument

package main

import (
	"flag"
	"fmt"
)

func main() {
	types := flag.String("types", "file", "please enter type")
	username := flag.String("username", "hubbusysyuhada", "please enter username")
	flag.Parse()                   // the command should be `go run flag.go -types=doc -username=anyone`
	fmt.Println(*types, *username) // need to use '*' because the outcome is a pointer
	// more in https://golang.org/pkg/flag/
}
