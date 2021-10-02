package main

import "fmt"

func main() {
	greet("hafidz")
	// countries := [...]string{
	// 	"indonesia",
	// 	"singapura",
	// 	"amerika",
	// 	"afrika",
	// }

	// sliced_array := countries[0:]
	// new_sliced := append(sliced_array, "jerman")
	// new_sliced = append(new_sliced, "venezuela")
	// fmt.Println(new_sliced)

	// new_array := make([]string, 2, 10)
	// new_array[0] = "bola"
	// new_array[1] = "kaos"
	// fmt.Println(new_array, "<<< new array")

	// person := make(map[string]string)
	// person["name"] = "hafidz"
	// person["job"] = "software engineer"
	// fmt.Println(person)
	// delete(person, "job")
	// fmt.Println(person)

	// person = map[string]string{
	// 	"name": "hafidz",
	// 	"age":  "25",
	// }

	// person["name"] = "budi"

	// if person["name"] == "hafidz" {
	// 	fmt.Println("helo hafidz")
	// } else if person["name"] == "budi" {
	// 	fmt.Println("helo budi")
	// } else {
	// 	fmt.Println("helo guys")
	// }

	// traffic_light := "grey"

	// switch traffic_light {
	// case "red":
	// 	fmt.Println("stop")
	// case "yellow":
	// 	fmt.Println("get ready")
	// case "green":
	// 	fmt.Println("go")
	// default:
	// 	fmt.Println("cannot recognize the traffic light")
	// }

	// students := []string{
	// 	"budi",
	// 	"sujatmiko",
	// 	"bagus",
	// }
	// for index, name := range students {
	// 	fmt.Println(index, name)
	// }

	// biodata := map[string]string{
	// 	"name":       "hafidz",
	// 	"age":        "25",
	// 	"occupation": "programmer",
	// }

	// for key, value := range biodata {
	// 	fmt.Println(key, value)
	// }
}

func greet(name string) {
	fmt.Println("hello", name)
}

func tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}
