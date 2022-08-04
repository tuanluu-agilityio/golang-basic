package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	counter = 0
	lock sync.Mutex
)

func main() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Microsecond * 10)
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
