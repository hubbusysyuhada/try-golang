package main

import "fmt"

type obj map[string]interface{}

func main() {
	greet()
	person := mapping("hafidz", 25)
	number := calculate(5, 3)
	fmt.Println("hasil dari 5 + 3 adalah", number)
	fmt.Println(person)

	// returning multiple values
	value1, value2 := math(3, 4)
	value3, _ := math(12, 100)
	fmt.Println(value1, value2, value3)
}

func math(num1 int, num2 int) (int, int) {
	return num1 * 2, num2 * 3
}

func greet() {
	fmt.Println("hello guys")
}

func calculate(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func mapping(name string, age int) obj {
	result := obj{
		"name": name,
		"age":  age,
	}
	return result
}
