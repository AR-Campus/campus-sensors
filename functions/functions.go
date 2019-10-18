package functions

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"db-training.de/campus-sensors/dataanalysis"
	"db-training.de/campus-sensors/sensors"
)

func DisplayInfos(data []sensors.SensorData, dataInit bool, w http.ResponseWriter, r *http.Request) {
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
	return
}

func UpdateWindowsFrontEndData(currentWindowStatus dataanalysis.WindowContactsStatus, w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(currentWindowStatus)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	return
}

func UpdateTopChartFrontEndData(sensorPackageHourFlowData []dataanalysis.SensorFlowPerHour, w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(dataanalysis.ParseSensorFlowPerHourJson(sensorPackageHourFlowData))
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	return
}

// UPDATE FLOW PER DAY
// func UpdateBottomChartFrontEndData(sensorPackageDayFlowData []dataanalysis.SensorFlowPerDay, w http.ResponseWriter, r *http.Request) {
// 	payload, err := json.Marshal(dataanalysis.ParseSensorFlowPerDayJson(sensorPackageDayFlowData))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(payload)
// 	return
// }

// UPDATE TEMP OVER TIME
func UpdateBottomChartFrontEndData(temperaturePackageFlowData []dataanalysis.TemperatureFlowPerHour, w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(dataanalysis.ParseTemperatureFlowPerHourJson(temperaturePackageFlowData))
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	return
}

//
// fmt.Fprintf(w, "Sensordaten und Graphische Darstellungen, %q", html.EscapeString(r.URL.Path))
// fmt.Fprintf(w, "\n \n")
// fmt.Fprintf(w, "Current Window Status:")
// fmt.Fprintf(w, "\n \n")
// dataanalysis.DrawWindowStatus(w, currentWindowStatus)
// fmt.Fprintf(w, "\n \n")
// fmt.Fprintf(w, "Number of Packages per Sensor:")
// fmt.Fprintf(w, "\n \n")
// fmt.Fprintf(w, "KueTHL:   %v", sentPackagesPerSensor.KuecheTempHumidLicht)
// fmt.Fprintf(w, "\n")
// fmt.Fprintf(w, "BStr-F-L: %v", sentPackagesPerSensor.BakerStrFensterLi)
// fmt.Fprintf(w, "\n")
// fmt.Fprintf(w, "BStr-F-R: %v", sentPackagesPerSensor.BakerStrFensterRe)
// fmt.Fprintf(w, "\n")
// fmt.Fprintf(w, "Kue-F-L:  %v", sentPackagesPerSensor.KuecheFensterLi)
// fmt.Fprintf(w, "\n")
// fmt.Fprintf(w, "Kue-F-R:  %v", sentPackagesPerSensor.KuecheFensterRe)
// fmt.Fprintf(w, "\n \n")
// fmt.Fprintf(w, "Sensor Package Flow:")
// fmt.Fprintf(w, "\n \n")
// for _, entry := range sensorPackageHourFlowData[(len(sensorPackageHourFlowData) - 10):] {
//   fmt.Fprintf(w, "time: %v  - ", entry.HourTimeData)
//   for s := 0; s <= entry.QuantityOfSensorPackages; s += 5 {
//     fmt.Fprintf(w, "#")
//   }
//   fmt.Fprintf(w, " - %v", entry.QuantityOfSensorPackages)
//   fmt.Fprintf(w, "\n")
// }
// fmt.Fprintf(w, "\n \n")
// for _, entry := range sensorPackageDayFlowData { //[(len(sensorPackageDayFlowData) - 1):]
//   fmt.Fprintf(w, "time: %v  - ", entry.DayTimeData)
//   for s := 0; s <= entry.QuantityOfSensorPackages; s += 100 {
//     fmt.Fprintf(w, "#")
//   }
//   fmt.Fprintf(w, " - %v", entry.QuantityOfSensorPackages)
//   fmt.Fprintf(w, "\n")
// }
// //[(len(sensorPackageHourFlowData) - 20):]
// fmt.Fprintf(w, "\n \n")
