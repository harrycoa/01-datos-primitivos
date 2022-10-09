package main

import (
	"fmt"

)
// ClaseUsuario
type Usuario struct {
	// propiedades o atributos
	edad int
	nombre, apellido string
}

func (usuariop Usuario) nombreCompleto() string {
	return usuariop.nombre + " " + usuariop.apellido
}

// pasar como puntero * cuando seteamos

func (this *Usuario) set_name(n string) {
	this.nombre = n
}


func main() {
	// Estructura de Datos

	// Llamar a la estructura
	var rodol Usuario
	fmt.Println(rodol)

	// otra forma de declarar
	harri := Usuario{edad: 15, nombre: "ha",apellido:"coa"}
	fmt.Println(harri)

	// otra forma
	uriel := Usuario{20,"uriel", "hernandez"}
	fmt.Println(uriel)

	// otra forma, retorna un puntero
	max := new(Usuario)
	max.nombre = "asa"
	max.apellido="ff"
	fmt.Println(max.nombre,max.apellido)
	fmt.Println(max.nombreCompleto())
	max.set_name("nuevo nombre")
	fmt.Println(max.nombreCompleto())

}