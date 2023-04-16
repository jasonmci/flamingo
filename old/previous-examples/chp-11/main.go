package main

import "log"

// structs can have functions assigned with them

type myStruct struct {
	FirstName string
}

// add a receiver in front of the func 
// you can add business logic here and that makes this useful
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Jason"

	myVar2 := myStruct{
		"Jennifer",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

}



