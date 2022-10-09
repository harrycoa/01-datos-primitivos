package main

import "fmt"

func main() {
	slice := []int{1,2,3,4}
	copia := make([]int,len(slice),cap(slice)*2)
	fmt.Println(slice)
	fmt.Println(copia)

	// copy(destino,fuente)
	// copia el minimo de elementos en ambos arreglos
	copy(copia,slice)
	fmt.Println(slice)
	fmt.Println(copia)
}