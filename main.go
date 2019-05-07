package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"db-training.de/campus-sensors/sensors"
	"github.com/gorilla/mux"
)

var data []sensors.SensorInfo

var authKey = os.Getenv("FIREFLY_APIKEY")

func Store(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
	if err != nil {
		fmt.Println("Got error")
	}

	result, _ := ioutil.ReadAll(r.Body)
	data = append(data, sensors.ConvertInfos(string(result))...)
}

func Infos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get data")
	json.NewEncoder(w).Encode(data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Sensors(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sensordaten und Graphische Darstellungen, %q", html.EscapeString(r.URL.Path))
}

func initData(lastN int64) {
	log.Printf("Load last %v packets from Firefly", lastN)
	FireFlyURL := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&limit_to_last=%v", authKey, lastN)
	response, err := http.Get(FireFlyURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	data = sensors.ConvertInfos(string(responseData))
}

func main() {

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/store", Store)
	router.HandleFunc("/infos", Infos)
	router.HandleFunc("/sensors", Sensors)

	herokuPort := os.Getenv("PORT")
	if herokuPort == "" {
		log.Fatal("$PORT must be set")
	}
	var port string
	if herokuPort == "" {
		port = "localhost:5555"
	} else {
		port = ":" + herokuPort
	}
	// var lastN int64
	// lastN, err := strconv.ParseInt(os.Getenv("NUMBER_OF_FIREFLY_ROWS"), 10, 64)
	// if err != nil {
	// 	lastN = 10
	// }

	//go initData(lastN)
	log.Print("Starting server at: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
