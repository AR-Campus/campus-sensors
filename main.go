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

	var sensorInfo sensors.SensorInfo
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

	// Fetch last Sensor-Package
	FireFlyCall := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&limit_to_last=1", authKey)
	response, err := http.Get(FireFlyCall)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		responseData, _ := ioutil.ReadAll(response.Body)
		data = append(data, sensors.ConvertInfos(string(responseData))...)
	}
	// } else {
	// 	responseData, _ := ioutil.ReadAll(response.Body)
	// 	data = responseData
	// 	fmt.Println("Called SensorPackages from FireFly")
	// }
	// post, err := http.Post(port+"/store", marshal(data), r.io.Reader)
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println("Stored Packages")
	// }

}
