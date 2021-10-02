/**
function yang akan selalu diakses ketika package tersebut diakses
cukup membuat func dengan nama init
contoh: open db
*/

package main

import (
	"Try-Golang/section-1/example"
	// _ "Try-Golang/section-1/example" // blank identifier, just run the init func
)

func main() {
	example.GetDB() // will first run the init func in example package
	example.Greeting()
}
