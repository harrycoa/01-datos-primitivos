package main

import (
	"fmt"

	mypackage "./package"
)

func main() {
	// Modificadores de acceso los tructs deben estar la primera letra en Mayuscula para ser publicos
	var myCar mypackage.CarPublic
	myCar.Brand = "lexus"
	myCar.Year = 1998

	fmt.Println(myCar)

	// Los structs privados no se pueden acceder desde fuera del paquete
	/*
		var myOther mypackage.CarPrivate
		fmt.Println(myOther)
	*/

	mypackage.PrintMessage()
	// mypackage.printMessage() // No se puede acceder desde fuera del paquete
}
