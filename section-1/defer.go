// defer merupakan sebuah function yang dieksekusi setelah function lain dieksekusi walaupun error

package main

import "fmt"

func main() {
	runApp()
}

func runApp() {
	defer logging()
	fmt.Println("application run")
}

func logging() {
	fmt.Println("logging")
}
