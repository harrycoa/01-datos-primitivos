package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" ****************************** Condicionales ******************************")

	// Condicional simple
	valor := 1
	if valor > 0 && valor == 1 { // operadores logicos
		println("=>")
	} else {
		if valor == 3 || valor == 4 {
			println("<=")
		}

	}

	// ejemplo real
	value, err := strconv.Atoi("48")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	// Par impar
	fmt.Printf("\n")
	parImpar(1854)

	// 2 password
	fmt.Printf("\n")
	password("m,", "M,")

	// Par impar SC
	fmt.Printf("\n")
	parImparSC(1854)

}

// par o impar
func parImpar(a int) {
	if a%2 == 0 {
		fmt.Printf("el numero %d es par", a)
	} else {
		fmt.Printf("el numero %d es impar", a)
	}
}

// Switch case
func parImparSC(a int) {

	switch modulo := a % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Impar")
	}
}

// password correctos
func password(cadena string, cadenaValidar string) {
	if strings.ToUpper(cadena) == strings.ToUpper(cadenaValidar) {
		fmt.Printf("El password %s es igual a %s", cadena, cadenaValidar)
	} else {
		fmt.Printf("El password %s es diferente a %s", cadena, cadenaValidar)
	}
}
