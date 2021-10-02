/**
secara default golang itu passed by value bukan by reference
pointer hanya mengubah value dari field tidak bisa mengubah nilai keseluruhan dari sebuah value passed by reference
pointer juga dapat dibuat dengan menggunakan function "new" namun hanya dapat mengembalikan ke data kosong
*/

package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{
		city:     "Depok",
		province: "Jawa Barat",
		country:  "Indonesia",
	}
	address2 := address1
	address2.city = "Bogor"
	fmt.Println(address1, "<<< address1 before pointer") // value dari address1 meskipun address2 di-reassign karena golang passed by value bukan reference
	address3 := &address1                                // ini adalah pointer, agar pass by reference
	address3.city = "Bekasi"                             // ini yang berubah hanya city menjadi bekasi
	fmt.Println(address1, "<<< address1 after pointer")
	fmt.Println(address2, "<<< address2")
	fmt.Println(address3, "<<< address3")
	// address3 = &Address{"Malang", "Jawa Timur", "Indonesia"} // ini akan mengubah keseluruhan dari address3 tapi tidak akan mengubah nilai dari Address1
	*address3 = Address{"Semarang", "Jawa Tengah", "Indonesia"} // ini akan mengubah keseluruhan dari address3 termasuk dari Address1 (passed by reference)
	fmt.Println(address3, "<<< address3 after")
	fmt.Println(address1, "<<< address1 after")

	// contoh menggunakan function new
	address4 := new(Address) // this equals to *address4 := Address{}, only works for empty value
	address5 := address4
	address5.city = "Boyolali"
	address4.country = "Indonesia"
	fmt.Println(address4, "<<<<< address4")
	fmt.Println(address5, "<<<<< address5")
}
