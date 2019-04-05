package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gorilla/mux"
)

var data []SensorInfo

var authKey = os.Getenv("FIREFLY_APIKEY")

type SensorInfo struct {
	device_eui string
	payload    string
	parsed     []struct {
		data map[string]string
	}
	// Hier Struktur von FireFly json
}

func Store(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
	if err != nil {
		fmt.Println("Got error")
	}

	var sensorInfo SensorInfo
	_ = json.NewDecoder(r.Body).Decode(&sensorInfo)
	data = append(data, sensorInfo)
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

func main() {

	// Fetch last Sensor-Package
	FireFlyCall := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&limit_to_last=10", authKey)
	response, err := http.Get(FireFlyCall)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/store", Store)
	router.HandleFunc("/infos", Infos)
	router.HandleFunc("/sensors", Sensors)

	herokuPort := os.Getenv("PORT")
	var port string
	if herokuPort == "" {
		port = "localhost:5555"
	} else {
		port = ":" + herokuPort
	}
	log.Fatal(http.ListenAndServe(port, router))
}
