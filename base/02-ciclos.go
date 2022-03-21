package main

import "fmt"

func main() {
	fmt.Println(" ****************************** Ciclos ******************************")
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")
	// For while
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}
	fmt.Printf("\n")
	// For forever
	/*
		cont2 := 0
		for {
			fmt.Println(count)
			cont2++
		}
	*/

	fmt.Printf("\n")
	for j := 10; j > 0; j-- {
		fmt.Println(j)
	}
	fmt.Printf("\n")

}
