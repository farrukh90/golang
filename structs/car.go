package main

import (
	"fmt"
)

type car struct {
	model string
	make  string
	year  int
}

func main() {
	var car car
	car.model = "BMW"
	car.make = "m3"
	car.year = 2019
	fmt.Printf("%+v", car)
}
