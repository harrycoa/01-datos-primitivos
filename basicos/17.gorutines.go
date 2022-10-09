package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Sincrono
	mi_nombre_lentooo("harry")
	// Asincrono, Concurrente
	// go - routines, ventaja sobre los hilos del sistema
	go mi_nombre_lentooo("coa")
	var wait string
	fmt.Scanln(&wait)

	// otra forma dentro de una funcion
	go func(){
		var wait2 string
		fmt.Scanln(&wait2)
	}()

}

func mi_nombre_lentooo(name string) {
	letras := strings.Split(name, "")
	for _,letra := range(letras) {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}