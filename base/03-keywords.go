package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ****************************** Condicionales ******************************")

	// Defer
	defer fmt.Println("A") // Se ejecutara antes de terminar la app
	fmt.Println("B")
	fmt.Println("C")

	// Continue
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 7 {
			fmt.Println("Es 7")
			break
		}
	}

	// Break

}
