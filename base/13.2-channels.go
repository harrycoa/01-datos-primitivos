package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {

	// channels
	c := make(chan string, 2) // 2 es la capacitdad maxima del channel
	c <- "hola"
	c <- "mundo"

	fmt.Println(len(c), cap(c))

	// Close .- cierra el canal y si agregars un valor no se puede leer
	close(c)

	// Range
	for msg := range c {
		fmt.Println(msg)
	}

	// select
	email := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-email:
			fmt.Println("email", msg)
		case msg := <-email2:
			fmt.Println("email2", msg)
		}
	}

}
