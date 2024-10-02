package main

import (
	"fmt"
	"sync"
)

type request struct {
	num1      float64
	operation string
	num2      float64
}

type response struct {
	result float64
	error  error
}

var requestChan = make(chan request)
var responseChan = make(chan response)

func Calculator() {
	for req := range requestChan {
		var result float64
		var err error

		switch req.operation {
		case "+":
			result = req.num1 + req.num2
		case "-":
			result = req.num1 - req.num2
		case "*":
			result = req.num1 * req.num2
		case "/":
			if req.num2 != 0 {
				result = req.num1 / req.num2
			} else {
				err = fmt.Errorf("Ошибка: деление на ноль")
			}
		default:
			err = fmt.Errorf("Ошибка: неизвестная операция: %s", req.operation)
		}

		responseChan <- response{result: result, error: err}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Calculator()
	}()

	requests := []request{
		{num1: 21, operation: "+", num2: 5},
		{num1: 13, operation: "-", num2: 8},
		{num1: 37, operation: "*", num2: 9},
		{num1: 57, operation: "/", num2: 14},
	}

	for _, req := range requests {
		requestChan <- req

		response := <-responseChan
		if response.error != nil {
			fmt.Println(response.error)
		} else {
			fmt.Printf("Результат для %.2f %s %.2f = %.2f\n", req.num1, req.operation, req.num2, response.result)
		}
	}

	close(requestChan)
	wg.Wait()
	close(responseChan)
}
