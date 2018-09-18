package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", e)
}

func Sqrt(x float64) (float64, error) {
	z:=1.0;
	e:=ErrNegativeSqrt(x);
	if x < 0 {
		// e = ErrNegativeSqrt(x)
		return x, e;
	}
	for i:=1;i<10;i++ {
		fmt.Printf("z: %g\n", z)
		z -= (z*z - x) / (2*z)
	}
	return z, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
