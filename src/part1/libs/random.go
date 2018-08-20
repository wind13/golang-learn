package libs

import ( 
	"fmt"
  "math/rand"
)

func Random() {
  rand.Seed(9)
  fmt.Printf("Random: %s!\n", rand.Intn(10))
}