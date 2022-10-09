package main

import (
	"net/http"
	"log"
	"fmt"
)
// Metodos http
// get(obtener), post(agregar), put(editar), delete(eliminar)

func main() {

	http.HandleFunc("/params",func(w http.ResponseWriter, r *http.Request){
		// separar la uri del query
		fmt.Println(r.URL)
		fmt.Println(r.URL.RawQuery)
		fmt.Println(r.URL.Query()) // mapa

		// manipular
		name:= r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println(name)
		}

		// obtener parametro
		parametro:= r.URL.Query().Get("parametro")
		if len(parametro) != 0 {
			fmt.Println(parametro)
		}

		fmt.Fprintf(w,"Hola mundo 3")
	})

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("El metodo es+" + r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola mundo con el metodo GET")
		case "POST":
			fmt.Fprintf(w, "Hola mundo con el metodo GET")
		case "PUT":
			fmt.Fprintf(w, "Hola mundo con el metodo GET")
		case "DELETE":
			fmt.Fprintf(w, "Hola mundo con el metodo GET")
		default:
			http.Error(w, "Metodo no valido", http.StatusBadRequest)
		}

	})

	http.HandleFunc("/header",func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("nombre","valor header")
		fmt.Fprintf(w, "response")

		// redireccionar
		// error 301- indica que el recurso se movio
		// estatus en pagina golang
		http.Redirect(w ,r , "https://www.google.com",http.StatusMovedPermanently)
	})

	http.HandleFunc("/dos",func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintf(w, "Hola mundo, dos")
		// pagina no encontrada
		http.NotFound(w ,r)
	})

	http.HandleFunc("/tres",func(w http.ResponseWriter, r *http.Request){
		// envio de errores
		http.Error(w, "este es un error", http.StatusConflict)
	})

	// log para el control de errores
	log.Fatal(http.ListenAndServe(":9000", nil))
}
