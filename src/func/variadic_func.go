package main

import "fmt"

func main() {
	myFunc(1, 2, 3, 4)
}

func myFunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number: %d\n", n)
	}
}
