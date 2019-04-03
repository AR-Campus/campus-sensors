package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
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
	// Wie HTTP - request Body augeben, lesbar
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Store data", bodyBuffer)
	//	var sensorInfo SensorInfo
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

	herokuPort := os.Getenv("PORT")
	var port string
	if herokuPort == "" {
		port = "locahost:5555"
	} else {
		port = ":" + herokuPort
	}
	log.Fatal(http.ListenAndServe(port, router))
}
