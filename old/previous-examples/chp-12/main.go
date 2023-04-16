package main

import (
	"log"
	"sort"
)

// type User struct {
// 	FirstName string
// 	LastName string
// }

// func main() {

// 	// and now we are going to create a map...

// 	// we dont really do this, dont do this.
// 	//var myOtherMap map[string]string

// 	myMap := make(map[string]User)

// 	me := User {
// 		FirstName: "Jason",
// 		LastName: "McInerney",
// 	}

// 	myMap["me"] = me

// 	// myMap["Dog"] = "Cody"

// 	// myMap["Other Dog"] = "Cassie"

// 	// log.Println(myMap["Dog"])
// 	// log.Println(myMap["Other Dog"])
// 	// now lets break down the syntax of that.

// 	log.Println(myMap["me"].FirstName) // use double quotes otherwise its a rune.

// }


func main() {
	// now we are going to move on to slices
	// first a variable
	//var myString string
	// now a slice 
	var myOtherSlice []string

	myOtherSlice = append(myOtherSlice, "Jason")
	myOtherSlice = append(myOtherSlice, "Jennifer")
	myOtherSlice = append(myOtherSlice, "Gambit")

	sort.Strings(myOtherSlice)

	log.Println(myOtherSlice)


	numbers := []int{1,2,3,4,5,6,7,8,9} // not sure why they are curly braces but ok

	log.Println(numbers)

	log.Println(numbers[0:2])
	
}