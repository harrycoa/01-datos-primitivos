package main

import "fmt"

func main() {
	// slice
	// se pueden declarar slice sin tamaño
	// slices vacio con tamaño cero
	// cuando se inicializa el slice inicializa en nulo como un objeto
	// componentes de un slice : puntero ala rreglo, longitud, capacidad
	var slice [] int
	fmt.Println(slice)

	slice2 := []int{1,2,3,4}
	fmt.Println(slice2)

	// podemos obtener un arreglo a partir de un slice
	arreglo := [3]int{1,2,3}
	sliceA := arreglo[:2]
	fmt.Println(arreglo)
	fmt.Println(sliceA)
	// operacion slicin, particionar un slice



}