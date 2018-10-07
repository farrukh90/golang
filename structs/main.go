package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "anderson"
	fmt.Println(alex)       //Get the value only
	fmt.Printf("%+v", alex) //Get the value and the variable

}
