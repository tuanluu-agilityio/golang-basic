package main

import(
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}

func main() {
  // power := 9000

  // name, power := getPower()
  // fmt.Printf("%s It's over %d\n", name, power)

  // goku := Saiyan{"Goku", 9000}

  // fmt.Printf("%s - %d\n", goku.Name, goku.Power)
  goku := &Saiyan{"Goku", 9000}

  Super(goku)
  fmt.Println(goku.Power)

  goku2 := &Saiyan{"Goku", 9000}
  goku2.SuperReceiver()
  fmt.Printf("goku2 Power: %d\n", goku2.Power)

  goku3 := NewSaiyan("goku3", 2000)
  fmt.Printf("%s - %d\n", goku3.Name, goku3.Power)

  fmt.Println(goku.Name)
}

func getPower() (string, int) {
  return "power", 9001
}

func Super(s *Saiyan) {
  s = &Saiyan{"Gohan", 9001}
}

func (s *Saiyan) SuperReceiver() {
  s.Power += 10000
}

func NewSaiyan(name string, power int) *Saiyan {
  return &Saiyan{
    Name: name,
    Power: power,
  }
}
