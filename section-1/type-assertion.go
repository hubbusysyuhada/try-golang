/**
mengubah datatype menjadi datatype lain yang diinginkan
umumnya digunakan pada interface kosong
*/

package main

import "fmt"

func main() {
	result := random()

	/**
	resultString := result.(string) // this is type assertion
	fmt.Println(result)
	fmt.Println(resultString)

	this will resulting error
	resultInt := result.(int) // panic
	fmt.Println(resultInt)
	to do so use switch case
	*/

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func random() interface{} {
	return "OK"
}
