// error interface khusus untuk memberikan error type

package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil1, err1 := divideNumber(2, 0)
	if err1 == nil {
		fmt.Println("hasilnya adalah", hasil1)
	} else {
		fmt.Println("error with:", err1.Error())
	}
}

func divideNumber(num1 int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		return num1 / divider, nil
	}
}
