package main

import (
"fmt"
"os"
"bufio"
)

func main() {
	// leer archivos linea por linea
	archivo, err := os.Open("./hola.txt")
	if err != nil {
		fmt.Println("Hubo errores  " )
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i,linea)
	}
	archivo.Close()
}