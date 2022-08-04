package main

import (
	"fmt"
	"math/rand"
)

func main() {
  lookup := make(map[string]int)
  lookup["goku"] = 9001
  power, exists := lookup["vegeta"]

  fmt.Println(power, exists);

  // get len of maps
  fmt.Println(len(lookup));

  // delete a key
  delete(lookup, "goku")
  delete(lookup, "vegeta")

  fmt.Println(len(lookup));

  // init a maps with make
  lookup2 := make(map[int]int32, 100)
  for i := 0; i < 200; i++ {
    lookup2[i] = rand.Int31n(100)
  }

  fmt.Println(lookup2);
  delete(lookup2, 10)
  fmt.Println(lookup2);

}
