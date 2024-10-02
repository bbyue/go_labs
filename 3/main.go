package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generator(c chan int) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		c <- num
		time.Sleep(time.Second)
	}
	close(c)
}

func Checker(c chan int, results chan string) {
	for num := range c {
		if num%2 == 0 {
			results <- fmt.Sprintf("%d четное", num)
		} else {
			results <- fmt.Sprintf("%d нечетное", num)
		}
	}
}

func main() {
	numChannel := make(chan int)
	resultChannel := make(chan string)

	go Generator(numChannel)
	go Checker(numChannel, resultChannel)

	for {
		select {
		case result := <-resultChannel:
			fmt.Println(result)
		case <-time.After(3 * time.Second):
			fmt.Println("Завершение работы программы.")
			return
		}
	}
}
