/**
same with it's name, it is a circular list of data
first value will become the last value and vice versa with no end
the length is not dinamic
example:
	1 => 2
	⬆	⬇
	4 <= 3

https://golang.org/pkg/container/ring/
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5)
	// to set the value of ring you must iterate it with for
	for i := 0; i < data.Len(); i++ {
		data.Value = fmt.Sprint("Value ", i+1)
		data = data.Next()
	}
	fmt.Println(data, "<<< data")
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
