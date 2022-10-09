package main 

import ( "fmt"
		"bufio" // direcciones de la memoria
		"os"
		)

func main() {
	edad := 22
	fmt.Println("Coloca tu edad ")	
	// leer valor
	fmt.Scanf("%d", &edad)
	fmt.Printf("mi edad es %d\n", edad)


	// v2
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese nombre: ")
	nombre, err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}