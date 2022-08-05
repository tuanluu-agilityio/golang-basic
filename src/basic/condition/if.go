package main

import (
	"fmt"
)

func main() {
	var input int
	if _, err := fmt.Scan(&input); input == 10 && err == nil {
		fmt.Println("Ping!!!")
	}
	fmt.Println("Done!!!")
}
