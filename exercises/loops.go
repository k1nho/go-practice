package exercises

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for n := 0; n < 7; n++ {
		z -= (z*z - x) /(2*z)
	}
	return z
}

func main() {
	for i:= 1; i < 100; i++ {
		fmt.Printf("for %d\n", i)
		fmt.Println(Sqrt(float64(i)))
		fmt.Println(math.Sqrt(float64(i)))
		fmt.Println();
	}

}
