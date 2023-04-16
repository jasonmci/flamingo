package main

import "fmt"

// This one is about interfaces then we get a quiz
// interfaces have functions that are part fo the itnerface
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Cody",
		Breed: "German Shepard",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Bruce",
		Color:         "Silver",
		NumberOfTeeth: 40,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and had", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Grr"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

// the next thing the instructor is going to do is make pointers instead
// the functions are using pointers
