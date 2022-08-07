package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}

func main() {
  // scores := [4]int{11, 12, 13, 14}

  // for index, value := range scores {
  //   fmt.Printf("%d - %d\n", index, value)
  // }

  // scores2 := make([]int, 0, 10)
  // for _, value := range scores {
  //   scores2 = append(scores2, value)
  // }

  // fmt.Printf("%d\n", scores2)

  // scores3 := make([]int, 5, 10)
  // scores3 = scores3[0:8]
  // scores3[7] = 999

  // fmt.Printf("%d\n", scores3)

  // scores4 := make([]int, 5)
  // scores4 = append(scores4, 9243)
  // fmt.Printf("%d\n", scores4)
  // saiyans := []*Saiyan {&Saiyan{"ABC", 123}, &Saiyan{"XYZ", 456}}
  // powers := extractPowers(saiyans)
  // for _, value := range powers {
  //   fmt.Printf("powers: %d\n", value)
  // }

  // // slice
  // scores := []int{1,2,3,4,5}
  // slice := scores[2:4]
  // slice[0] = 999
  // fmt.Println(scores)

  scores := []int{1, 2, 3, 4, 5}
  scores = removeAtIndex(scores, 2)
  fmt.Println(scores)
}

func removeAtIndex(source []int, index int) []int {
  lastIndex := len(source) - 1
  source[index], source[lastIndex] = source[lastIndex], source[index]
  return source[:lastIndex]
}

func extractPowers(saiyans []*Saiyan) []int {
  powers := make([]int, len(saiyans))

  for index, saiyan := range saiyans {
    powers[index] = saiyan.Power
  }

  return powers
}
