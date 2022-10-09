package main

import "fmt"

func main() {
	/*
	1. Un puntero es una direccion de memoria
	2. en lugar de obtener el valor tenemos la direccion del valor
	3. w , z => #dirmemoria
	4. w=> 6
	5. z=> 6
	delarar un puntero *
	*/

	var puntero *int

	fmt.Println(puntero)

	var w,z *int
	entero := 5
	w = &entero
	z = &entero
	fmt.Println(w)
	fmt.Println(z)

	// acceso al valor del espacio de memoria
	*w = 7
	fmt.Println(*w)
	fmt.Println(*z)

}