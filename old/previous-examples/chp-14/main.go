package main

import "log"

func main() {
	// all loops are for loops
	// for i:= 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// this is good to know
	animals := []string{"dog", "cat", "mouse", "horse", "cow"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// or if you dont care about the counter
	for _, animal := range animals {
		log.Println(animal)
	}

	moreAnimals := make(map[string]string)

	moreAnimals["dog"] = "penny"
	for aType, animal := range moreAnimals {
		log.Println(aType, animal)
	}

	// in go it is a slice of bytes
	// and a rune is a byte

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	// create a variable user that is an array with a type of User datatype
	var users []User

	users = append(users, User{"Jason", "McInerney", "jason@example.com", 44})
	users = append(users, User{"Jennifer", "McInerney", "jennifer@example.com", 43})

	for _, u := range users {
		log.Println(u.FirstName, u.LastName)
	}
}
