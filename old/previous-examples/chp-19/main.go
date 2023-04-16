package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.3, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("the answer is:", result)
}


// go makes testing easy


func divide(a, b float32) (float32, error) {
	var result float32

	if b == 0 {
		return result, errors.New("division by zero error")
	}
	result = a / b
	return result, nil
}