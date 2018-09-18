package libs

import ( 
	"fmt"
  "math/rand"
  "time"
)

func Random() {
  rand.Seed(time.Now())
  fmt.Printf("Random: %v\n", rand.Intn(10))
}