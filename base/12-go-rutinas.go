package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	println(s)
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello, World!")

	wg.Add(1)

	go say("Hello World2", &wg)

	wg.Wait()

	// funcion anonima
	go func(text string) {
		fmt.Println(text)
	}("adios")

	time.Sleep(time.Second * 1)
}
