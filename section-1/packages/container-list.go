/**
package for double linked list
looks like an array
example:
	nil <=> 4 <=> 23 <=> 74 <=> 231 => nil
starts and ends with nil
works two ways
cannot accessed by index, must be iterated 1 by 1

list.PushBack => push
list.PushFront => unshift
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("muhammad")
	data.PushBack("hafidz")
	data.PushBack("hubbusysyuhada")
	data.PushBack("adalah")
	data.PushBack("saya")
	data.PushFront(1)
	fmt.Println(data.Front())
	fmt.Println(data.Back())
	fmt.Println(data)
	fmt.Println(data.Front().Next().Next().Next().Prev(), "<<<< next next terus")

	// to iterate list use for like below
	for v := data.Front(); v != nil; v = v.Next() {
		fmt.Println(v, "<<<< v")
	}
}
