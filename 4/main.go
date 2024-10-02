package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex

func Increment() {
	for i := 0; i < 1000; i++ {
		//mutex.Lock()
		count++
		//mutex.Unlock()
	}
}

func main() {
	go Increment()
	go Increment()
	go Increment()
	go Increment()
	go Increment()
	time.Sleep(1 * time.Second)
	fmt.Println("Счетчик:", count)
}
