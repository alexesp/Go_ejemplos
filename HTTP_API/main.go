package main

import (
	"fmt"
	"log"
	"net/http"
)

// w - responseWriter (Donde escribir respuesta)
// r - request (Dedonde obtener la solicitud)
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi! I'm new web-server!</h1>")
}

// Un amigo que selecciona el controlador deseado según la dirección a la que llegó la solicitud.
func ReuqestHandler() {
	http.HandleFunc("/", GetGreet)               // Si llega una solicitud a la dirección "/", llame GetGreet
	log.Fatal(http.ListenAndServe(":8080", nil)) // Iniciamos el servidor web en modo “escucha”
}

func main() {
	ReuqestHandler()
}
