// package untuk memanipulasi tipe data string
// .Trim(), .ToLower(), .ToUpper(), .Split(), .Contains(), .ReplaceAll(), dll
// https://golang.org/pkg/strings/

package main

import (
	"fmt"
	"strings"
)

func main() {
	trimmed := strings.Trim("   aduh  ", " ")
	lowered := strings.ToLower("HAFIDZ")
	uppered := strings.ToUpper("hubbusysyuhada")
	splitted := strings.Split("muhammad-hafidz-hubbusysyuhada", "-")
	containsTrue := strings.Contains("muhammad hafidz hubbusysyuhada", "hafidz")
	containsFalse := strings.Contains("muhammad hafidz hubbusysyuhada", "budi")
	replacedAll := strings.ReplaceAll("muhammad hafidz hubbusysyuhada", "hubbusysyuhada", "ganteng")
	replaced := strings.Replace("Muhammad Hafidz Hubbusysyuhada", "u", "i", 2)
	fmt.Println(trimmed)
	fmt.Println(lowered)
	fmt.Println(uppered)
	fmt.Println(splitted)
	fmt.Println(containsTrue)
	fmt.Println(containsFalse)
	fmt.Println(replacedAll)
	fmt.Println(replaced)
}
