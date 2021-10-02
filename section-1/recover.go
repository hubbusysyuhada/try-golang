/**
recover merupakan sebuah function untuk menangkap panic
dengan recover, proses panic akan berhenti dan program akan lanjut berjalan
recover disimpan di defer func agar terbaca
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
	message := recover()
	if message != nil {
		fmt.Println("error occurs, details", message)
	}
}
