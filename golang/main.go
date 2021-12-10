package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/structs"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
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

func getCpuUsage() float64 {
	cmd := exec.Command("grep", "cpu ", `/proc/stat`)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("error al correr comando", err)
	}
	totals := strings.Split(string(stdout), " ")
	_2, _ := strconv.Atoi(totals[2])
	_4, _ := strconv.Atoi(totals[4])
	_5, _ := strconv.Atoi(totals[5])
	return (float64((_2+_4)*100) / float64(_2+_4+_5))
}

func getMemory() structs.Memoria {
	ram, _ := ioutil.ReadFile("/proc/memo_201900289")
	var memoria structs.Memoria
	json.Unmarshal(ram, &memoria)
	return memoria
}

func getCPU() structs.Cpu {
	processes, _ := ioutil.ReadFile("/proc/cpu_201900289")
	var cpu structs.Cpu
	json.Unmarshal(processes, &cpu)
	return cpu
}
