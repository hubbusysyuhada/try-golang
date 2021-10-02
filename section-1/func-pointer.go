/**
menggunakan (*) pada parameter
*/

package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	Depok := Address{"Depok", "Jawa Barat", "Indonesia"}
	fmt.Println(Depok, "<<<< depok")
	changeCity(Depok, "Bandung") // won't change
	fmt.Println(Depok, "<<<< depok after")
	changeCityPointer(&Depok, "Bekasi") // will change
	fmt.Println(Depok, "<<<< depok after pointer change")

}

func changeCity(address Address, city string) {
	address.city = city
}

func changeCityPointer(address *Address, city string) {
	address.city = city
}
