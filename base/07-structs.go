package main

import "fmt"

type car struct {
	brand string
	year  int
	color string
	model string
}

func main() {
	// Structs
	// instanciar structs

	// v1
	myCar := car{brand: "Toyota", year: 2022}
	fmt.Println(myCar)

	// v2
	var myCar2 car
	myCar2.brand = "Ferrari"

	fmt.Println(myCar2)

}
