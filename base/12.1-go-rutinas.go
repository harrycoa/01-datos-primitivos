package main

import (
	"fmt"
	"time"
)

func main() {

	// Ciclos repetitivos para go rutinas
	nombres := []string{"Juan", "Pedro", "Maria", "Ana"}
	go saludar(nombres)
	go despedir(nombres)

	teclado := ""
	fmt.Scanln(&teclado)

}

func saludar(nombres []string) {
	for _, nombre := range nombres {
		fmt.Println("Hola", nombre)
		time.Sleep(time.Second * 1)
	}
}

func despedir(nombres []string) {
	for _, nombre := range nombres {
		fmt.Println("adios", nombre)
		time.Sleep(time.Second * 1)
	}
}
