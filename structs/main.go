package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "anderson"
	alex.contact.email = "a@gmail.com"
	alex.contact.zip = 60659

	fmt.Println(alex)       //Get the value only
	fmt.Printf("%+v", alex) //Get the value and the variable
}
