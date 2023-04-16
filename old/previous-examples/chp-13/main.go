package main

import "log"

// conditionals

func main() {

	// var isTrue bool
	
	// isTrue = false

	// if isTrue {
	// 	log.Println("isTrue", isTrue)

	// } else {
	// 	log.Println("isTrue", isTrue)
	// }

	myVar := "cajt"
	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")
	case "dog":
		log.Println("Cat is set to dog")
	case "mouse":
		log.Println("cat is set to mouse")
	default:
		log.Println("Cat is set to somethnig else")
	}
	
}