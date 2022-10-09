package main
import (
	"os"
	"fmt"
	"bufio"
)
func main() {
	// uso de defer para control de errores
	ejecucion := readFile()
	fmt.Println(ejecucion)
}
func readFile() bool {
	archivo, err := os.Open("./hola.txt")
	defer func() {
		archivo.Close()
		fmt.Println("defer")
	}()
	if err != nil {
		fmt.Println("Hubo errores  " )
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i,linea)
	}
	archivo.Close()
	return true
}
