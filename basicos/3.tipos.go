package main

import (
	"fmt" // paquete de escritura
	"strconv" // paquete de conversion
)

func main() {
	edad := 22

	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es : " +edad_str)

}