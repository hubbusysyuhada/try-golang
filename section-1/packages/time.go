/**
https://golang.org/pkg/time/
time management package
there are so many but mainly .Date(), .Now(), and .Parse(layout string)
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Date(2002, 8, 27, 15, 49, 51, 2000, time.Local))
	parsedDate, _ := time.Parse("2006-01-02", "1945-08-17")
	fmt.Println(parsedDate)
}
