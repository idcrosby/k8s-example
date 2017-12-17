package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Count int = 0
var logger *log.Logger
var f *os.File
var countF *os.File
var message string = "Hello From Kubernetes." 

func main() {

	f, _ = os.OpenFile("/var/log/app.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	countF, _ = os.OpenFile("/var/log/counter", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)

	defer f.Close()
	fmt.Println("init web server...")

	http.HandleFunc("/", hello)
	http.HandleFunc("/readiness", readiness)
	http.HandleFunc("/liveness", liveness)
	http.HandleFunc("/counter", counter)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request recived at %s", time.Now().String())
	f.WriteString("Request recived at " + time.Now().String() + "\n")
	
	host, _ := os.Hostname()
	fmt.Fprint(w, message + "\nI'm running on ", host, "\n")
}

func readiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Readiness Probe called at %s", time.Now().String())
	f.WriteString("Readiness Probe called at " + time.Now().String() + "\n")
	time.Sleep(500 * time.Millisecond)
	w.Write([]byte("I'm ready."))
}

func liveness(w http.ResponseWriter, r *http.Request) {
	f.WriteString("Liveness probe called at %s" + time.Now().String() + "\n")
	fmt.Println("Liveness probe called at %s", time.Now().String())
	time.Sleep(500 * time.Millisecond)
	w.Write([]byte("I am Alive!"))
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Counter called at %s", time.Now().String())
	f.WriteString("Counter called at %s" + time.Now().String() + "\n")
	Count = Count +1
	countF.WriteString("Count: " + strconv.Itoa(Count) + "\n")
	fmt.Fprintf(w,"Count: %d", Count)
}
