package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/",fileServer))
	// Levantamos el servidor en el puerto 9000
	http.ListenAndServe(":9000",nil)

}

