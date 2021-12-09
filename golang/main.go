package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/structs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	ram, _ := ioutil.ReadFile("/proc/memo_201900289")
	processes, _ := ioutil.ReadFile("/proc/cpu_201900289")
	var memoria structs.Memoria
	var cpu structs.Cpu
	fmt.Println(string(processes))
	json.Unmarshal(ram, &memoria)
	json.Unmarshal(processes, &cpu)
	fmt.Println(memoria)
	fmt.Println(cpu)
	// makeServer()
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

func reader(connection *websocket.Conn) {
	for {
		messageType, p, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := connection.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func writerRam(connection *websocket.Conn) {
	data := structs.Memoria{Total_memory: 123, Free_memory: 321, Used_memory: 231}
	for {
		if err := connection.WriteJSON(data); err != nil {
			log.Println(err)
			return
		}
	}
}

func writerCpu(connection *websocket.Conn) {
	process := structs.Process{Pid: 1, Name: "buenas", User: 1, State: 2, Child: []int{1, 2, 3}}
	data := structs.Cpu{Processes: []structs.Process{process}, Running: 1, Sleeping: 2, Zombie: 3, Stopped: 4, Total: 5}
	for {
		if err := connection.WriteJSON(data); err != nil {
			log.Println(err)
			return
		}
	}
}

func socketMemory(response http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected to RAM")
	go writerRam(ws)
}

func socketCpu(response http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client conected to CPU")
	go writerCpu(ws)
}
