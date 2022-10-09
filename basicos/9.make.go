package main

import "fmt"

func main() {
	slice := make([]int,3,5)
	// tamaño del slice
	fmt.Println(len(slice))
	// capacidad
	fmt.Println(cap(slice))
	 // agregar elementos al slice, inicialmente tiene 3 ceros
	fmt.Println(slice)
	slice = append(slice,2)
	fmt.Println(slice)

}