/**
package for math operations
https://golang.org/pkg/math/
some new math func in go compared to js:
 - math.Max(float64, float64) => return max value of float64
 - math.Min(float64, float64) => return min value of float64
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(math.Max(100, 200))
	fmt.Println(math.Min(100, 200))
	fmt.Println(math.Round(1.499))
	fmt.Println(math.Floor(1.499))
	fmt.Println(math.Ceil(1.499))
	fmt.Println(rand.Float32())
}
