/**
package to convert string type
https://golang.org/pkg/strconv/
parseBool => to boolean, parseFloat => to float, parseInt => to int
formatBool => from boolean, formatFloat => from float, formatInt => from int
almost all func in strconv could return 2 values: value itself and error on second value
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	boole, _ := strconv.ParseBool("true")
	floated, _ := strconv.ParseFloat("2.3124", 64)
	integ, err := strconv.ParseInt("12142", 10, 32)
	fmt.Println(boole)
	fmt.Println(floated)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(integ)

	Atoi, _ := strconv.Atoi("123123") // automatically use decimal as base number type
	Itoa := strconv.Itoa(12312414)    // same as atoi
	formatBool := strconv.FormatBool(true)
	formatFloat := strconv.FormatFloat(32.415, 'e', 'e', 64)
	formatInt := strconv.FormatInt(24124123, 10)
	fmt.Println(formatBool)
	fmt.Println(formatFloat)
	fmt.Println(formatInt)
	fmt.Println(Atoi)
	fmt.Println(Itoa)
}
