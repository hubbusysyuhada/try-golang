// nil == null pada javascript

package main

import "fmt"

func main() {
	hafidz := newMap("hafidz")
	empty := newMap("")
	fmt.Println(hafidz)
	fmt.Println(empty)
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
