package main

import (
	"fmt"
)

func main() {
	// los go channels permiten comunicar gorutines
	// ......
	// definir un canal
	canal := make(chan string)
	// crear una gorutine
	go func(canal chan string){
		for {
			var nombre string
			fmt.Scanln(&nombre)
			// lado derecho dato que queremos enviar
			// lado izquierdo canal en el cual vamos a enviar
			canal <- nombre
		}
	}(canal)
	for {
		msg := <-canal
		fmt.Println("Estoy imprimiendo lo que recibi del canal:  " + msg)
	}


}