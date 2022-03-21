package main

import "fmt"

func main() {
	// Diccionarios / llaves / maps
	m := make(map[string]int)

	m["jose"] = 14
	m["pepe"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {

		fmt.Println(i, v)
	}

	// encontrar valor
	value, ok := m["jose"] // ok indica si esa llave existe en el diccionario
	fmt.Println(value, ok)

	value2, ok := m["josess"] // ok indica si esa llave existe en el diccionario
	fmt.Println(value2, ok)
}
