/**
panic function merupakan sebuah function untuk menghentikan program, biasanya diletakkan di error
defer function akan tetap dieksekusi
panic seperti throw new Error pada javascript
*/

package main

import "fmt"

func main() {
	runApp(false)
	runApp(true)
}

func runApp(error bool) {
	defer logging()
	fmt.Println("application run")
	if error {
		panic("error!!!")
	}
	fmt.Println("application end")
}

func logging() {
	fmt.Println("defer logging")
}
