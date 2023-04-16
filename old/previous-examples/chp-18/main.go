package main

import (
	"encoding/json"
	"log"

)

// we use this to unmarshall



type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "John",
		"last_name": "McHenry",
		"hair_color": "Gray",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "Black",
		"has_dog": false
	}
]
	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json: ", err)
	}

	log.Printf("unmarshaled: %v", unmarshalled)


	// write json from a struct


	var mySlice []Person

	var m1 Person
	m1.FirstName = "Thor"
	m1.LastName = "Odenson"
	m1.HairColor = "blonde"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Tony"
	m2.LastName = "Stark"
	m2.HairColor = "Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshalling to json", err)
	}
	log.Println(string(newJson))
}