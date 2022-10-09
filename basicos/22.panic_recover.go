package main
import (
	"os"
	"fmt"
	"bufio"
)
func main() {
	executeReadFile()
	fmt.Println("ejecucion despues del panic")

}
func executeReadFile(){
	ejecucion := readFile()
	fmt.Println(ejecucion)
}
func readFile() bool {
	archivo, err := os.Open("./holaa.txt")
	defer func() {
		archivo.Close()
		fmt.Println("defer")
		r := recover()
		fmt.Println(r)
	}()
	if err != nil {
		panic(err)
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
