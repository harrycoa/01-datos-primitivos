package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("")

		// creamos la estructura curson donde pasaremos a json
		cursos := Cursos{
			Curso{"matematicas",4},
			Curso{"lenguaje",8},
			Curso{"node JS",9},
		}
		json.NewEncoder(w).Encode(cursos)
	})
	// Levantamos el servidor en el puerto 9000
	http.ListenAndServe(":9000",nil)

}

type Curso struct {
	Titulo string `json:"titulaso"`
	NumeroVideos int
}

// conjunto de estas estructuras
type Cursos []Curso


