package main

func main() {
	// you can't re-assign a constant even with same data types
	// if you make a constant but never used, golang will not throw error
	// when declaring constant you must declaring the value as well (with const in front)
	const A = "A"
	const (
		B string = "B"
		C        = "C"
	)
}
