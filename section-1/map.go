package main

import "fmt"

func main() {
	// pada dasarnya map itu mirip dengan object pada javascript / bahasa lain
	person := map[string]string{
		"name": "hafidz",
		"age":  "",
	} // map[type key]type value{}
	person["age"] = "25"
	person["city"] = "depok"
	fmt.Println(person)
	fmt.Println(person["city"])
	fmt.Println(person["address"])

}
