package main

import "fmt"

func main() {
	fmt.Println(" ****************************** Repaso Golang ******************************")

	/*1.- Constantes */
	const pi float64 = 3.14
	const p2 float64 = 3.15
	fmt.Println(pi, p2)

	/*2.- Variables */
	base := 12          // forma 1 - variable directamente asignada
	var altura int = 14 // forma 2
	var area int        // forma 3 - Go no compila si las variables no son usadas

	fmt.Println(base, altura, area, base*altura)

	/*3.- Zero values */
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculo de area
	const bas = 17
	area = bas * bas
	fmt.Println(area)

	/*4.- Operadores aritmeticos */
	x := 37
	y := 9

	result := x + y
	fmt.Println("suma: ", result)

	result = x - y
	fmt.Println("resta: ", result)

	result = x * y
	fmt.Println("producto: ", result)

	result = x / y
	fmt.Println("division: ", result)

	result = x % y
	fmt.Println("residuo %: ", result)

	x++
	fmt.Println("incremental: ", x)

	x--
	fmt.Println("decremental: ", x)

	// Areas
	// Rectangulo
	const baseR int = 10
	const alturaR int = 17
	rectangulo := baseR * alturaR
	fmt.Println("rectangulo: ", rectangulo)

	// Triangulo
	const baseT int = 10
	const alturaT int = 17
	triangulo := baseT * alturaT / 2
	fmt.Println("triangulo: ", triangulo)

	// Circulo
	const radio = 64
	circulo := pi * radio * radio
	fmt.Println("circulo: ", circulo)

	/*5.- Datos primitivos */
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	/*6.- Paquete fmt */

	// variables
	helloMessge := "hello"
	worlMessage := "word"

	// print ln
	fmt.Println(helloMessge, worlMessage)

	// printf

	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/

	nombre := "abc"
	edad := 25
	fmt.Printf("%s tiene mas de %d años\n", nombre, edad)
	fmt.Printf("%v tiene mas de %v añoss\n", nombre, edad) // v no tienes certeza que dato ira

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d años\n", nombre, edad)
	fmt.Println(message)

	// Saber tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMessge)
	fmt.Printf("edad: %T\n", edad)

	escribir("mmmm")

	// Return
	value := returnValue(2)
	fmt.Println(value)

	// Return doble
	value1, value2 := dobleReturn(4)
	fmt.Println(value1, value2)
	fmt.Printf("valor1: %d,valor2 : %d", value1, value2)

	fmt.Println("")
	fmt.Println("______________________________________")
	fmt.Println("")
	// Arbol

	// Arbol 2

	alt := 5

	/*
		for y := 1; y <= (alt*2)-1; y++ {
			if y <= alt {
				fmt.Print(y)
			} else {
				fmt.Print(y - alt)
			}

		}
	*/

	/*
	   fmt.Println("")
	   	for x := 1; x <= alt; x++ {

	   		for e := 1; e <= alt-x; e++ {
	   			fmt.Print(e)
	   		}
	   		for n := 1; n <= (x*2)-1; n++ {
	   			fmt.Print("*")
	   		}

	   		fmt.Println("")

	   	}
	*/

	fmt.Println("")
	for x := 1; x <= alt; x++ {
		for e := 1; e <= alt-x; e++ {
			fmt.Print(e)
		}
		for n := 1; n <= (x*2)-1; n++ {
			fmt.Print("*")
			if n == (x*2)-1 {
				fmt.Println("l")
			}

		}

		fmt.Println("")

	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")

	alturaTriangulo := 5
	aux := alturaTriangulo
	inicial := 0

	for a := 1; a <= alturaTriangulo; a++ {

		// izq
		for b := 1; b <= aux; b++ {
			fmt.Print(b)
		}
		// mid
		for c := 1; c <= inicial-1; c++ {
			fmt.Print(" ")
		}
		// der
		for d := aux; d >= 1; d-- {

			fmt.Printf("%d", d)

		}

		aux = aux - 1
		inicial = inicial + 2
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")

	i := 5
	j := 5
	n := 5

	for i = 0; i <= n; i++ {
		fmt.Println("")

		// izq
		for j = 1; j <= n-i; j++ {
			fmt.Print(j)
		}
		// med
		for j = 0; j < 2*i-1; j++ {
			fmt.Print(" ")
		}
		// der
		for j = n - i; j > 0; j-- {
			if j != n {
				fmt.Printf("%d", j)
			}

		}

	}

	fmt.Println("****************************   Funciones    ******************************")

	val, _ := dobleReturn(8)

	fmt.Println(val)

}

func escribir(message string) {
	fmt.Println("funcion ext: ", message)
}

func returnValue(a int) int {
	return a * 2
}

// retorno de 2 valores
func dobleReturn(a int) (c, d int) {
	return a, a * 2
}
