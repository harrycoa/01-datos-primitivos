package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.URL.Path[1:])
		http.ServeFile(w,r,r.URL.Path[1:])
	})
	// Levantamos el servidor en el puerto 9000
	http.ListenAndServe(":9000",nil)

}

