package main

import (
  "fmt"
  "math/rand"
  "math"
  // "time"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

var container = []string{"zero", "one", "two"}
var a, b, c = 3, 4, 7

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return y, x
}

func main() {
  // container := map[int]string{0: "zero", 1: "one", 2: "two"}
  // rand.Seed(time.Now())
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println(split(17))
  fmt.Println(a, b, c)

  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
  
  e := 42
  o := float64(e)
  u := uint(o)

  fmt.Println(e)
  fmt.Println(o)
  fmt.Println(u)

  var x, y int = 3, 4
  var c float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(c)
  fmt.Println(x, y, z)

  var w int = int(math.Sqrt(float64(x*x + y*y)))
  fmt.Printf("%T, %T\n", z, w)

  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  // error:  constant 1267650600228229401496703205376 overflows int
	// fmt.Println(Big) 
  fmt.Println(needFloat(Big))
  
  sum := 1
	for sum < 100 {
		sum += sum
	}
  fmt.Println(sum)
  fmt.Println(sqrt(2), sqrt(-4))
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}