package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Ler archivos completos en memoria
	dato_leido,err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Println("Hubo errores  " )
	}
	// convertimos lo leido a string --> string(datoaconvertir)
	fmt.Println(string(dato_leido))

}