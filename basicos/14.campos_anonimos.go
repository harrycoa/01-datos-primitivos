package main

import "fmt"

// Definiendo herencia
type CPersona struct {
	nombre string
}

type CDocente struct {
	CPersona
}

func (this CPersona) hablar() string {
	return "bienvenido"
}

func main()  {

	docente := CDocente{CPersona{"harri"}}

	fmt.Println(docente.nombre)
	fmt.Println(docente.hablar())

}