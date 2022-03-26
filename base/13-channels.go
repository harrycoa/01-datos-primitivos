package main

import (
	"fmt"
)

func say(s string, c chan string) {
	c <- s
	// println(s)
}

// dentro de los canales podemos indicar si el canal sera de entrada o si sera de salida
func say2(s string, c <-chan string) {
	s = <-c
	// este canal sera solo de salida

}

func say3(s string, c chan<- string) {
	c <- s
	// este canal solo sera de entrada
}

func main() {

	// channels
	c := make(chan string, 1)
	fmt.Println("Channel:", c)

	go say3("Hello", c)

	fmt.Println(<-c)
}
