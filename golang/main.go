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
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// fmt.Println(getCPU())
	// fmt.Println(getMemory())
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
	router.HandleFunc("/ram", socketMemory)
	router.HandleFunc("/cpu", socketCpu)
	fmt.Println("server up in " + port + " port")
	http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router))
}

func welcome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Go api"))
}

func readerCPU(connection *websocket.Conn) {
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

func readerRam(connection *websocket.Conn) {
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
	for {
		data := getMemory()
		if err := connection.WriteJSON(data); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func writerCpu(connection *websocket.Conn) {
	for {
		data := getCPU()
		if err := connection.WriteJSON(data); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func socketMemory(response http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected to RAM")
	writerRam(ws)
	log.Println("Client disconected to RAM")
}

func socketCpu(response http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client conected to CPU")
	writerCpu(ws)
	log.Println("Client disconected to CPU")
}

func getCpuUsage() float64 {
	cmd := exec.Command("sh", "-c", `grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage ""}'`)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("error al correr comando", err)
	}
	salida := strings.Trim(strings.Trim(string(stdout), " "), "\n")
	valor, _ := strconv.ParseFloat(salida, 64)
	return (valor)
}

func getCache() float64 {
	cmd := exec.Command("sh", "-c", `free -m | head -n2 | tail -1 | awk '{print $6}'`)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("error al correr comando", err)
	}
	salida := strings.Trim(strings.Trim(string(stdout), " "), "\n")
	valor, _ := strconv.ParseFloat(salida, 64)
	return valor
}

func getMemory() structs.Memoria {
	ram, _ := ioutil.ReadFile("/proc/memo_201900289")
	var memoria structs.Memoria
	json.Unmarshal(ram, &memoria)
	memoria.Cache_memory = getCache()
	memoria.Used_memory = (memoria.Total_memory - memoria.Free_memory - int(getCache())) * 100 / memoria.Total_memory
	memoria.Available_memory = memoria.Free_memory + int(getCache())
	return memoria
}

func getCPU() structs.Cpu {
	processes, _ := ioutil.ReadFile("/proc/cpu_201900289")
	var cpu structs.Cpu
	json.Unmarshal(processes, &cpu)
	cpu.Usage = getCpuUsage()
	return cpu
}
