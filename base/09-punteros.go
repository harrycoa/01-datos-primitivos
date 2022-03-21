package main

import (
	"fmt"
)

type pc struct {
	ram    int
	disk   int
	memory int
	brand  string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "ping")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	// Punteros y structs
	a := 50
	b := &a

	fmt.Println(a, b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	mypc := pc{ram: 4, disk: 500, memory: 8, brand: "Dell"}

	fmt.Println(mypc)

	mypc.ping()

	mypc.duplicateRAM()
	fmt.Println(mypc)

	mypc.duplicateRAM()
	fmt.Println(mypc)

	mypc.duplicateRAM()
	fmt.Println(mypc)

}
