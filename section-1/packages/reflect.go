/**
package untuk membaca struktur code ketika sedang berjalan
https://golang.org/pkg/reflect/
*/

package main

import (
	"fmt"
	"reflect"
)

type Orang struct {
	name, city string `required:"true"`
	age        int
	isMarried  bool `tag:"true"` // this is a tag in struct using backtick (``), tag value must be string
	// tags can be used as a validation
}

func main() {
	person1 := Orang{"Hafidz", "depok", 20, false}
	sample := reflect.TypeOf(person1)
	field := sample.Field(0)
	isMarriedTag := sample.Field(3).Tag.Get("tag")
	invalidTag := sample.Field(0).Tag.Get("tag")
	isValid := validatePerson(person1)
	fmt.Println("reflect")
	fmt.Println(person1)
	fmt.Println(sample)
	fmt.Println(sample.NumField())
	fmt.Println(field)
	fmt.Println(field.Name)
	fmt.Println(isMarriedTag)
	fmt.Println(invalidTag)
	fmt.Println(isValid)
}

func validatePerson(person Orang) bool {
	sample := reflect.TypeOf(person)
	for i := 0; i < sample.NumField(); i++ {
		field := sample.Field(i)
		required := field.Tag.Get("required")
		value := reflect.ValueOf(person).Field(i).Interface()
		if required == "true" {
			return value != ""
		}

	}
	return true
}
