package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var data []string

type SensorInfo struct {
	Message string // Hier Struktur von FireFly json
}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Store data", r.Body)
	var sensorInfo SensorInfo

	fmt.Println(sensorInfo)

	// Wie HTTP - request Body augeben, lesbar

	//_ = json.NewDecoder(r.Body).Decode(&sensorInfo)
	//data = append(data, sensorInfo.Message)
}

func Infos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get data")
	json.NewEncoder(w).Encode(data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/store", Store)
	router.HandleFunc("/infos", Infos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "locahost:5555"
	}
	log.Fatal(http.ListenAndServe(port, router))

}
