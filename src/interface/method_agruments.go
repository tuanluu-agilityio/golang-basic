package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

// Human implements fmt.Stringer
func (h Human) String() string {
	return "Name:" + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

func main() {
	Bob := Human{"Bob", 37, "222-222-XXX"}
	fmt.Println("This Human is: ", Bob)
}
