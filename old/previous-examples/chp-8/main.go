package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world!"
	fmt.Println(whatToSay)

	i = 42

	fmt.Println("i is set to", i)

	whatWasSaid, whatElseWasSaid := saySomething() 

	fmt.Println("The function returned", whatWasSaid, whatElseWasSaid)
}

func saySomething() (string, string) {
	return "Something", "Some one"
}