package main

import (
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"sync"
	"time"

	"db-training.de/campus-sensors/dataanalysis"
	"db-training.de/campus-sensors/functions"
	"db-training.de/campus-sensors/sensors"
	"github.com/gorilla/mux"
)

var data []sensors.SensorData
var cacheData []sensors.SensorData
var dataInit = false
var dateStart time.Time
var dateEnd time.Time
var currentWindowStatus dataanalysis.WindowContactsStatus
var sensorPackageHourFlowData []dataanalysis.SensorFlowPerHour
var sensorPackageHourFlowJson dataanalysis.SensorFlowPerHourPackageJson
var sensorPackageDayFlowData []dataanalysis.SensorFlowPerDay
var sensorPackageDayFlowJson dataanalysis.SensorFlowPerHourPackageJson
var sentPackagesPerSensor dataanalysis.PackagesPerSensorCount
var homepageTpl *template.Template

// var navigationBarHTML string

var authKey = os.Getenv("FIREFLY_APIKEY")

// Config provides basic configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

func Store(w http.ResponseWriter, r *http.Request) {
	if dataInit {
		requestDump, err := httputil.DumpRequest(r, false)
		log.Println("Got firefly push", string(requestDump))
		if err != nil {
			fmt.Println("Got error")
		}

		result, _ := ioutil.ReadAll(r.Body)
		newEntry := sensors.ConvertSingle(string(result))
		currentWindowStatus, sensorPackageHourFlowData, sensorPackageDayFlowData, sentPackagesPerSensor = dataanalysis.UpdateAnalysisData(newEntry[0], currentWindowStatus, sensorPackageHourFlowData, sensorPackageDayFlowData, sentPackagesPerSensor)
		log.Printf("/Store newWindowStatus?: %v", currentWindowStatus)
		data = append(data, newEntry...)
	}
}

func Infos(w http.ResponseWriter, r *http.Request) {
	functions.DisplayInfos(data, dataInit, w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Sensors(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/campus_sensors.html")
}

func UpdateWindowsFrontEnd(w http.ResponseWriter, r *http.Request) {
	functions.UpdateWindowsFrontEndData(currentWindowStatus, w, r)
}

func UpdateTopChartFrontEnd(w http.ResponseWriter, r *http.Request) {
	functions.UpdateTopChartFrontEndData(sensorPackageHourFlowData, w, r)
}

func UpdateBottomChartFrontEnd(w http.ResponseWriter, r *http.Request) {
	functions.UpdateBottomChartFrontEndData(sensorPackageDayFlowData, w, r)
}

func getLastSensorPackageDateTime() string {
	FireFlyURL := fmt.Sprintf("https://api.fireflyiot.com/api/v1/packets?auth=%v&limit_to_last=1", authKey)
	response, err := http.Get(FireFlyURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	cacheData = sensors.ConvertInfos(string(responseData))
	return cacheData[0].Time
}

func initData() { // For Loop untli All Packets from starting Date on are received
	currentWindowStatus = dataanalysis.WindowContactsStatus{BakerStrFensterLi: false, BakerStrFensterRe: false, KuecheFensterLi: false, KuecheFensterRe: false}
	sensorPackageHourFlowData = make([]dataanalysis.SensorFlowPerHour, 0)
	sensorPackageDayFlowData = make([]dataanalysis.SensorFlowPerDay, 0)
	startDate := os.Getenv("START_DATE")
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
				currentWindowStatus, sensorPackageHourFlowData, sensorPackageDayFlowData, sentPackagesPerSensor = dataanalysis.UpdateAnalysisData(entry, currentWindowStatus, sensorPackageHourFlowData, sensorPackageDayFlowData, sentPackagesPerSensor)
				log.Printf("initData update windowStatus?: %v from %v", currentWindowStatus, entry)
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
		log.Printf("Length of SensorFlowHourArray: %v", len(sensorPackageHourFlowData))
		log.Printf("Length of SensorFlowDayArray: %v", len(sensorPackageDayFlowData))
		time.Sleep(1 * time.Second)
	}
	dataInit = true
	log.Printf("Initialisation complete!")
}

func main() {

	router := mux.NewRouter().StrictSlash(false)
	// router.HandleFunc("/", Index)
	router.HandleFunc("/store", Store)
	router.HandleFunc("/infos", Infos)
	router.HandleFunc("/sensors", Sensors)
	router.HandleFunc("/updatewindows", UpdateWindowsFrontEnd)
	router.HandleFunc("/updatetopchart", UpdateTopChartFrontEnd)
	router.HandleFunc("/updatebottomchart", UpdateBottomChartFrontEnd)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("html/"))))

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
	go initData()

	log.Print("Starting server at: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
