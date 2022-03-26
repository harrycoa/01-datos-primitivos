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

	// Stringers
	myPC := pc{ram: 16, disk: 1, memory: 4, brand: "Dell"}

	// Impresion personalizada
	fmt.Println(myPC)

}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de Disco, y es la marca %v", myPC.memory, myPC.disk, myPC.brand)
}
