package main

import (
	"fmt"
)

type figuras2D interface {
	area() float64
}
type cuadrado struct {
	lado float64
}

type rectangulo struct {
	base, altura float64
}

func (c *cuadrado) area() float64 {
	return c.lado * c.lado
}

func (r *rectangulo) area() float64 {
	return r.base * r.altura
}

func calculaArea(f figuras2D) {
	fmt.Println(f.area())
}

// Interfaces

func main() {

	myCuadrado := cuadrado{lado: 8}
	myRectangulo := rectangulo{base: 10, altura: 5}

	fmt.Println(myCuadrado)
	fmt.Println(myRectangulo)

	//calculaArea(&myCuadrado)
	//calculaArea(&myRectangulo)

	myInterface := []interface{}{"hola", 12, 8.66, true}
	fmt.Println(myInterface)
	fmt.Println(myInterface...)
}
