package main

import (
	"log"

	//"github.com/jasonmci/myniceprogram/helpers"
)

const numPool = 1000
func CalculateValue(intChan chan int) {
	//randomNumber := helpers.RandomNumber(numPool)
	//intChan <- randomNumber
}


func main() {
	intChan := make(chan int)
	// i am creating a channel to send information and it can only hold ints
	defer close(intChan) // do this because it is good practice

	// ah, a concurrent operation

	go CalculateValue(intChan)

	// next we have to listen

	myNumber := <- intChan
	log.Println(myNumber)

}

// i created a new go.mod file so this is no longer valid. 