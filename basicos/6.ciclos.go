package main

import "fmt"

func main() {

	// no hay while
	for i:=0; i< 10; i=i+2 {
		fmt.Println("=>", i)
	}

	// simular el ciclo while

	j:=0
	for j< 10 {
		fmt.Println("=>", j)
	}


	k := 0
	for {
		fmt.Println(k)
		k++
		if k > 10 {
			break
		}

	}

}