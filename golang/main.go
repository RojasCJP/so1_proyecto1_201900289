package main

import (
	"fmt"
	"main/structs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	prueba := structs.Prueba{Nombre: "juan", Edad: 21}
	fmt.Println(prueba)
	makeServer()
}

func makeServer() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}
	router.HandleFunc("/", welcome).Methods("GET")
	fmt.Println("server up in " + port + " port")
	http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router))
}

func welcome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Go api"))
}
