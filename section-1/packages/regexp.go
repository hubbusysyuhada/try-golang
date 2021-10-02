/**
https://golang.org/pkg/regexp/
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	emailRegexp := regexp.MustCompile(`([a-z])@([a-z]).([a-z])`)
	isMatched := emailRegexp.MatchString("a@b.c")
	isEmail := emailRegexp.MatchString("john@example.com")
	fmt.Println(emailRegexp)
	fmt.Println(isMatched)
	fmt.Println(isEmail)
}
