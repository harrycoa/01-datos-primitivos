package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	// usaremos un mutex que permite manipular el bloqueo y desbloqueo
	// en este caso como creamos una variable temporal en algun momento esa variable nombreTemporal puede sr asignada al muismo tiempo y bloquearse

	wg.Add(2)
	defer wg.Wait()

	nombreTemporal := ""

	// Ciclos repetitivos para go rutinas
	nombres := []string{"Juan", "Pedro", "Maria", "Ana"}
	go saludar(nombres, &wg, &nombreTemporal, &mx)
	go despedir(nombres, &wg, &nombreTemporal, &mx)

}

// simulando un problema en concurrencia

func saludar(nombres []string, wg *sync.WaitGroup, nombreTemporal *string, mx *sync.Mutex) {
	defer wg.Done()
	for _, nombre := range nombres {
		mx.Lock()
		nombreTemporal = &nombre
		fmt.Printf("Hola %s \n", *nombreTemporal)
		time.Sleep(time.Second * 1)
		mx.Unlock()
	}
}

func despedir(nombres []string, wg *sync.WaitGroup, nombreTemporal *string, mx *sync.Mutex) {
	defer wg.Done()
	for _, nombre := range nombres {
		mx.Lock()
		nombreTemporal = &nombre
		fmt.Printf("adios %s \n", *nombreTemporal)
		time.Sleep(time.Second / 2)
		mx.Unlock()
	}
}
