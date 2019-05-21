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
	"strconv"
	"time"

	"db-training.de/campus-sensors/dataanalysis"
	"db-training.de/campus-sensors/sensors"
	"github.com/gorilla/mux"
)

var data []sensors.SensorData
var cacheData []sensors.SensorData
var dataInit = false
var dateStart time.Time
var dateEnd time.Time
var currentWindowStatus dataanalysis.WindowContactsStatus
var sensorPackageFlowData []dataanalysis.SensorFlowPerHour
var sentPackagesPerSensor dataanalysis.PackagesPerSensorCount

var authKey = os.Getenv("FIREFLY_APIKEY")

func Store(w http.ResponseWriter, r *http.Request) {
	if dataInit {
		requestDump, err := httputil.DumpRequest(r, false)
		log.Println("Got firefly push", string(requestDump))
		if err != nil {
			fmt.Println("Got error")
		}

		result, _ := ioutil.ReadAll(r.Body)
		newEntry := sensors.ConvertSingle(string(result))
		currentWindowStatus, sensorPackageFlowData, sentPackagesPerSensor = dataanalysis.UpdateAnalysisData(newEntry[0], currentWindowStatus, sensorPackageFlowData, sentPackagesPerSensor)
		data = append(data, newEntry...)
	}
}

func Infos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get data")
	dataLen := len(data)
	intervalShown := 10
	fmt.Fprintf(w, "Sensordaten in der Pseudo-Datenbank: %v, %q", dataLen, html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, "First %v Entries:", intervalShown)
	fmt.Fprintf(w, "\n")
	beginning := data[:intervalShown]
	json.NewEncoder(w).Encode(beginning)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Last %v Entries:", intervalShown)
	fmt.Fprintf(w, "\n")
	dataend := data[(dataLen - intervalShown):]
	json.NewEncoder(w).Encode(dataend)
	fmt.Fprintf(w, "\n \n")
	if dataInit {
		fmt.Fprintf(w, "Initialisation complete!")
	} else {
		fmt.Fprintf(w, "Initialisation running!")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Sensors(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sensordaten und Graphische Darstellungen, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, "Current Window Status:")
	fmt.Fprintf(w, "\n \n")
	if currentWindowStatus.BakerStrFensterLi == false {
		fmt.Fprintf(w, "BStr-F-L: geschlossen")
	} else {
		fmt.Fprintf(w, "BStr-F-L: offen")
	}
	fmt.Fprintf(w, "\n")
	if currentWindowStatus.BakerStrFensterRe == false {
		fmt.Fprintf(w, "BStr-F-R: geschlossen")
	} else {
		fmt.Fprintf(w, "BStr-F-R: offen")
	}
	fmt.Fprintf(w, "\n")
	if currentWindowStatus.KuecheFensterLi == false {
		fmt.Fprintf(w, "Kue-F-L: geschlossen")
	} else {
		fmt.Fprintf(w, "Kue-F-L: offen")
	}
	fmt.Fprintf(w, "\n")
	if currentWindowStatus.KuecheFensterRe == false {
		fmt.Fprintf(w, "Kue-F-R: geschlossen")
	} else {
		fmt.Fprintf(w, "Kue-F-R: offen")
	}
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, "Number of Packages per Sensor:")
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, "KueTHL:   %v", sentPackagesPerSensor.KuecheTempHumidLicht)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "BStr-F-L: %v", sentPackagesPerSensor.BakerStrFensterLi)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "BStr-F-R: %v", sentPackagesPerSensor.BakerStrFensterRe)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Kue-F-L:  %v", sentPackagesPerSensor.KuecheFensterLi)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Kue-F-R:  %v", sentPackagesPerSensor.KuecheFensterRe)
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, "Sensor Package Flow:")
	fmt.Fprintf(w, "\n \n")
	for _, entry := range sensorPackageFlowData[(len(sensorPackageFlowData) - 24):] {
		fmt.Fprintf(w, "time: %v  - ", entry.HourTimeData)
		for s := 0; s <= entry.QuantityOfSensorPackages; s += 5 {
			fmt.Fprintf(w, "#")
		}
		fmt.Fprintf(w, " - %v", entry.QuantityOfSensorPackages)
		fmt.Fprintf(w, "\n")
	}
	//[(len(sensorPackageFlowData) - 20):]
	fmt.Fprintf(w, "\n \n")

}

// func ReInitialize(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Sensordaten mit aktueller Anzahl an Sensordaten aus GetEnv in x100, %q", html.EscapeString(r.URL.Path))
// 	var lastN int64
// 	lastN, err := strconv.ParseInt(os.Getenv("NUMBER_OF_FIREFLY_ROWS"), 10, 64)
// 	if err != nil {
// 		lastN = 10
// 	}
// 	go initData(lastN)
// }

func getLastSensorPackageDateTime() string {
	FireFlyURL := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&limit_to_last=1", authKey)
	response, err := http.Get(FireFlyURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	cacheData = sensors.ConvertInfos(string(responseData))
	// fmt.Println(cacheData)
	return cacheData[0].Time
}

func initData(lastN int64) { // For Loop untli All Packets from starting Date on are received
	currentWindowStatus = dataanalysis.WindowContactsStatus{BakerStrFensterLi: false, BakerStrFensterRe: false, KuecheFensterLi: false, KuecheFensterRe: false}
	sensorPackageFlowData = make([]dataanalysis.SensorFlowPerHour, 0)
	// sentPackagesPerSensor =
	startDate := os.Getenv("START_DATE")
	// sensorPackageFlowData[0] = dataanalysis.SensorFlowPerHour{HourTimeData: startDate, QuantityOfSensorPackages: 1}
	log.Printf("Load last packets from Firefly starting currently at %v", startDate)
	endDate := getLastSensorPackageDateTime()
	log.Printf("Current EndDate: %v", endDate)
	dateStart, _ = time.Parse(time.RFC3339, startDate)
	dateEnd, _ = time.Parse(time.RFC3339, endDate)
	for dateStart.Before(dateEnd) {
		FireFlyURL := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&direction=asc&received_after=%v&limit_to_last=100", authKey, startDate)
		response, err := http.Get(FireFlyURL)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return
		}
		responseData, _ := ioutil.ReadAll(response.Body)
		cacheData = sensors.ConvertInfos(string(responseData))
		if len(cacheData) == 0 {
			log.Printf("Current cacheData is empty!")
			data = append(make([]sensors.SensorData, 0))
		} else {
			for _, entry := range cacheData {
				// log.Printf(" from UpdateLoop: now SensorFlowHourArray length is : %v", len(sensorPackageFlowData))
				// currentWindowStatus = dataanalysis.WindowContactSensorsUpdate(entry, currentWindowStatus)
				// sensorPackageFlowData = dataanalysis.SensorFlowArrayUpdate(entry, sensorPackageFlowData)
				// sentPackagesPerSensor = dataanalysis.QuantifyPerSensorPackages(entry, sentPackagesPerSensor)
				currentWindowStatus, sensorPackageFlowData, sentPackagesPerSensor = dataanalysis.UpdateAnalysisData(entry, currentWindowStatus, sensorPackageFlowData, sentPackagesPerSensor)
			}
			startDate = cacheData[(len(cacheData) - 1)].Time
			dateStart, _ = time.Parse(time.RFC3339, startDate)
			log.Printf("Load last packets from Firefly starting currently at %v", startDate)
			data = append(data, cacheData[:(len(cacheData)-1)]...)
			log.Printf("DataBase current size: %v", len(data))
		}
		endDate := getLastSensorPackageDateTime()
		time.Sleep(2 * time.Second)
		log.Printf("Check for new EndDate, now new at %v", endDate)
		log.Printf("Length of SensorFlowArray: %v", len(sensorPackageFlowData))
		time.Sleep(1 * time.Second)
	}
	dataInit = true
	log.Printf("Initialisation complete!")
}

func main() {

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", Index)
	router.HandleFunc("/store", Store)
	router.HandleFunc("/infos", Infos)
	router.HandleFunc("/sensors", Sensors)
	// router.HandleFunc("/reinit", ReInitialize).Methods("POST")

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
	var lastN int64
	lastN, err := strconv.ParseInt(os.Getenv("NUMBER_OF_FIREFLY_ROWS"), 10, 64)
	if err != nil {
		lastN = 10
	}
	go initData(lastN)

	log.Print("Starting server at: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
