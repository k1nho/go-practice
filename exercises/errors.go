package main

import (
	"fmt"
)

// implementing Error method from the error interface for this new type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", int(e))
}

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for n := 0; n < 7; n++ {
		z -= (z*z - x) /(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

