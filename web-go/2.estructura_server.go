package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
func Hola2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo dos")
}

func main() {
	// server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola2)

	// ruta
	http.HandleFunc("/", Hola) //defaultServerMux

	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: nil, // Si es el nil utilizamos

	}

	log.Fatal(server.ListenAndServe())

	// log para el control de errores
	log.Fatal(http.ListenAndServe(":9000", nil))

}
