package main

import "fmt"

func main() {
	// arreglos
	var arreglo [4] int
	fmt.Println(arreglo)

	arreglo2 := [3]int{1,2,3}
	fmt.Println(arreglo2)

	for i:=0; i < len(arreglo2);i++{
		fmt.Println(arreglo2[i])
	}

	// Matrices
	var matriz [2][2] int
	fmt.Println(matriz)

}