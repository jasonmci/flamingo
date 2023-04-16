package main

import (
	"log"
	"time"
)


type User struct {
	FirstName string // Capital letters so we might use them as a type for a different package
	LastName string
	PhoneNumber string
	Age int
	Birthdata time.Time
}

func main() {
	user := User {
		FirstName: "Jason",
		LastName: "McInerney",
	}
	
	log.Println(user.FirstName, user.LastName)
}

