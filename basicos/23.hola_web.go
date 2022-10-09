package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/holamundo",func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "holaaaa")
	})
	http.HandleFunc("/",handler)
	// Levantamos el servidor en el puerto 9000
	http.ListenAndServe(":9000",nil)


	fmt.Println("Hola web")
	// io.Closer()
}

func handler(w http.ResponseWriter, r *http.Request){
	// escribimos un hola web en la plantilla principal
	fmt.Println("hay una nueva peticion")
	io.WriteString(w, "Hola web")
}