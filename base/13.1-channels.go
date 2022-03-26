package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// channels
	chNumeros := make(chan int)
	chImprimir := make(chan string)

	chNumerosPares := make(chan int)
	chNumerosImpares := make(chan int)

	defer close(chNumeros)
	defer close(chImprimir)
	defer close(chNumerosPares)
	defer close(chNumerosImpares)

	go GenerarNumeros(chNumeros)
	go RecibirNumeros(chNumeros, chNumerosPares, chNumerosImpares)
	go NumeroPar(chNumerosPares, chImprimir)
	go NumeroImPar(chNumerosImpares, chImprimir)

	for {
		msgRecibido := <-chImprimir
		fmt.Println(msgRecibido)
	}

}

func NumeroPar(ch chan int, chSalida chan string) {
	for {
		numero := <-ch
		salida := "numero par"
		salida += strconv.Itoa(numero)
		chSalida <- salida
	}
}
func NumeroImPar(ch chan int, chSalida chan string) {
	for {
		numero := <-ch
		salida := "numero impar"
		salida += strconv.Itoa(numero)
		chSalida <- salida
	}
}

func GenerarNumeros(ch chan int) {
	numero := 0
	for {
		numero++
		ch <- numero
		time.Sleep(time.Second) // dormir 1 seg
	}

}

func RecibirNumeros(ch chan int, chPar chan int, chImpar chan int) {
	for {
		numero := <-ch
		if numero%2 == 0 {
			// canal de pares
			chPar <- numero
		} else {
			// canal de impares
			chImpar <- numero
		}
	}

}
