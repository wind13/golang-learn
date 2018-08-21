package libs

import ( 
	"fmt"
  "math"
)

func Sqrt(n float64) {
  var s = math.Sqrt(n)
  fmt.Printf("sqrt: %v\n", s)
  // fmt.Printf("sqrt: %g\n", s)
}