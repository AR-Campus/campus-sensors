package dataanalysis

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"db-training.de/campus-sensors/sensors"
)

type WindowContactsStatus struct {
	BakerStrFensterLi bool `json:"BakerStrFensterLi"`
	BakerStrFensterRe bool `json:"BakerStrFensterRe"`
	KuecheFensterLi   bool `json:"KuecheFensterLi"`
	KuecheFensterRe   bool `json:"KuecheFensterRe"`
}

type SensorFlowPerHour struct {
	HourTimeData             string `json:"HourTimeData"`
	QuantityOfSensorPackages int    `json:"QuantityOfSensorPackages"`
}

type SensorFlowPerHourPackageJson struct {
	// HourMatrix []map[String] `json:"HourMatrix"`
	FlowMatrix []int `json:"FlowMatrix"`
}

type SensorFlowPerDay struct {
	DayTimeData              string `json:"DayTimeData"`
	QuantityOfSensorPackages int    `json:"QuantityOfSensorPackages"`
}

type SensorFlowPerDayPackageJson struct {
	// DayMatrix  []String `json:"DayMatrix"`
	FlowMatrix []int `json:"FlowMatrix"`
}

// type SensorFlowPerMonth struct {
// 	MonthTimeData            string `json:"Hour-Time-Data"`
// 	QuantityOfSensorPackages int    `json:"Quantity-Of-Sensor-Packages"`
// }

type PackagesPerSensorCount struct {
	KuecheTempHumidLicht int64 `json:"KuecheTempHumidLicht"`
	BakerStrFensterLi    int64 `json:"BakerStrFensterLi"`
	BakerStrFensterRe    int64 `json:"BakerStrFensterRe"`
	KuecheFensterLi      int64 `json:"KuecheFensterLi"`
	KuecheFensterRe      int64 `json:"KuecheFensterRe"`
}

func RandomBool() bool {
	return rand.Intn(2) == 0
}

func WindowContactSensorsUpdate(singleSensorData sensors.SensorData, currentWindowStatus WindowContactsStatus) WindowContactsStatus {
	switch singleSensorData.DeviceID {
	case "BakerStrFensterLi":
		currentWindowStatus.BakerStrFensterLi = singleSensorData.SensorValues["ReedSensor"].(bool)
	case "BakerStrFensterRe":
		currentWindowStatus.BakerStrFensterRe = singleSensorData.SensorValues["ReedSensor"].(bool)
	case "KuecheFensterLi":
		currentWindowStatus.KuecheFensterLi = singleSensorData.SensorValues["ReedSensor"].(bool)
	case "KuecheFensterRe":
		currentWindowStatus.KuecheFensterRe = singleSensorData.SensorValues["ReedSensor"].(bool)
	}
	currentWindowStatus = WindowContactsStatus{BakerStrFensterLi: RandomBool(),
		BakerStrFensterRe: RandomBool(),
		KuecheFensterLi:   RandomBool(),
		KuecheFensterRe:   RandomBool()}
	log.Printf("WindowStatus: %v", currentWindowStatus)
	return currentWindowStatus
}

func SensorFlowHourArrayUpdate(singleSensorData sensors.SensorData, sensorPackageFlowData []SensorFlowPerHour) []SensorFlowPerHour {
	timeOfSensor, _ := time.Parse(time.RFC3339, singleSensorData.Time)
	// log.Printf("Time of sensor: %v", timeOfSensor)
	if len(sensorPackageFlowData) != 0 {
		// log.Printf("from SensorFlowUpdate: SensorFlowHour is : %v", sensorPackageFlowData[0].HourTimeData)
		lastFlowHour := len(sensorPackageFlowData) - 1
		currentSensorFlowHour, _ := time.Parse(time.RFC3339, sensorPackageFlowData[lastFlowHour].HourTimeData)
		if timeOfSensor.Before(currentSensorFlowHour.Add(time.Hour * 1)) {
			// log.Printf("from SensorFlowUpdate-sensorBefore-true")
			sensorPackageFlowData[lastFlowHour].QuantityOfSensorPackages++
		} else {
			// log.Printf("from SensorFlowUpdate-sensorBefore-false - next hour")
			sensorPackageFlowData = append(sensorPackageFlowData, SensorFlowPerHour{HourTimeData: singleSensorData.Time, QuantityOfSensorPackages: 1})
			// log.Printf("Next SensorFlowHour at: %v", sensorPackageFlowData[1].HourTimeData)
		}
	}
	if len(sensorPackageFlowData) == 0 {
		sensorPackageFlowData = append(sensorPackageFlowData, SensorFlowPerHour{HourTimeData: singleSensorData.Time, QuantityOfSensorPackages: 1})
		log.Printf("First SensorFlowHour added!")
	}
	return sensorPackageFlowData
}

func SensorFlowDayArrayUpdate(singleSensorData sensors.SensorData, sensorPackageFlowData []SensorFlowPerDay) []SensorFlowPerDay {
	timeOfSensor, _ := time.Parse(time.RFC3339, singleSensorData.Time)
	// log.Printf("Time of sensor: %v", timeOfSensor)
	if len(sensorPackageFlowData) != 0 {
		// log.Printf("from SensorFlowUpdate: SensorFlowHour is : %v", sensorPackageFlowData[0].HourTimeData)
		lastFlowHour := len(sensorPackageFlowData) - 1
		currentSensorFlowHour, _ := time.Parse(time.RFC3339, sensorPackageFlowData[lastFlowHour].DayTimeData)
		if timeOfSensor.Before(currentSensorFlowHour.Add(time.Hour * 24)) {
			// log.Printf("from SensorFlowUpdate-sensorBefore-true")
			sensorPackageFlowData[lastFlowHour].QuantityOfSensorPackages++
		} else {
			// log.Printf("from SensorFlowUpdate-sensorBefore-false - next hour")
			sensorPackageFlowData = append(sensorPackageFlowData, SensorFlowPerDay{DayTimeData: singleSensorData.Time, QuantityOfSensorPackages: 1})
			// log.Printf("Next SensorFlowHour at: %v", sensorPackageFlowData[1].HourTimeData)
		}
	}
	if len(sensorPackageFlowData) == 0 {
		sensorPackageFlowData = append(sensorPackageFlowData, SensorFlowPerDay{DayTimeData: singleSensorData.Time, QuantityOfSensorPackages: 1})
		log.Printf("First SensorFlowHour added!")
	}
	return sensorPackageFlowData
}

func QuantifyPerSensorPackages(singleSensorData sensors.SensorData, currentSensorQuantities PackagesPerSensorCount) PackagesPerSensorCount {
	switch singleSensorData.DeviceID {
	case "KuecheTempHumidLicht":
		currentSensorQuantities.KuecheTempHumidLicht++
	case "BakerStrFensterLi":
		currentSensorQuantities.BakerStrFensterLi++
	case "BakerStrFensterRe":
		currentSensorQuantities.BakerStrFensterRe++
	case "KuecheFensterLi":
		currentSensorQuantities.KuecheFensterLi++
	case "KuecheFensterRe":
		currentSensorQuantities.KuecheFensterRe++
	}
	return currentSensorQuantities
}

func UpdateAnalysisData(singleSensorData sensors.SensorData, currentWindowStatus WindowContactsStatus, sensorPackageHourFlowData []SensorFlowPerHour, sensorPackageDayFlowData []SensorFlowPerDay, currentSensorQuantities PackagesPerSensorCount) (WindowContactsStatus, []SensorFlowPerHour, []SensorFlowPerDay, PackagesPerSensorCount) {
	// log.Printf("From UpdateFunc: now SensorFlowHourArray length is : %v", len(sensorPackageFlowData))
	currentWindowStatus = WindowContactSensorsUpdate(singleSensorData, currentWindowStatus)
	sensorPackageHourFlowData = SensorFlowHourArrayUpdate(singleSensorData, sensorPackageHourFlowData)
	sensorPackageDayFlowData = SensorFlowDayArrayUpdate(singleSensorData, sensorPackageDayFlowData)
	currentSensorQuantities = QuantifyPerSensorPackages(singleSensorData, currentSensorQuantities)
	return currentWindowStatus, sensorPackageHourFlowData, sensorPackageDayFlowData, currentSensorQuantities
}

func DrawWindowStatus(w http.ResponseWriter, currentWindowStatus WindowContactsStatus) {
	fmt.Fprintf(w, "#=====================#\n")
	fmt.Fprintf(w, "|_______|             |\n")
	fmt.Fprintf(w, "|  |                  |\n")
	fmt.Fprintf(w, "|  |      <b>KÜCHE       |\n")
	fmt.Fprintf(w, "|__|                  |\n")
	fmt.Fprintf(w, "|_                    |\n")
	fmt.Fprintf(w, " /                    |\n")
	fmt.Fprintf(w, "/                     |\n")
	if currentWindowStatus.KuecheFensterLi == true {
		fmt.Fprintf(w, "|                  |--|\n")
		fmt.Fprintf(w, "|           offen  |  |\n")
		fmt.Fprintf(w, "|                  |--|\n")
		fmt.Fprintf(w, "|                     |\n")
		fmt.Fprintf(w, "|                     |\n")
	} else {
		fmt.Fprintf(w, "|                    ||\n")
		fmt.Fprintf(w, "|              zu    ||\n")
		fmt.Fprintf(w, "|                    ||\n")
		fmt.Fprintf(w, "|                     |\n")
		fmt.Fprintf(w, "|                     |\n")
	}
	if currentWindowStatus.KuecheFensterRe == true {
		fmt.Fprintf(w, "|                  |--|\n")
		fmt.Fprintf(w, "|           offen  |  |\n")
		fmt.Fprintf(w, "|                  |--|\n")
		fmt.Fprintf(w, "|                     |\n")
		fmt.Fprintf(w, "#===#=================#\n")
		fmt.Fprintf(w, "    |                 |\n")
	} else {
		fmt.Fprintf(w, "|                    ||\n")
		fmt.Fprintf(w, "|              zu    ||\n")
		fmt.Fprintf(w, "|                    ||\n")
		fmt.Fprintf(w, "|                     |\n")
		fmt.Fprintf(w, "#===#=================#\n")
		fmt.Fprintf(w, "    |    BAKER-STR    |\n")
	}
	if currentWindowStatus.BakerStrFensterLi == true {
		fmt.Fprintf(w, "    |              |--|\n")
		fmt.Fprintf(w, "    |       offen  |  |\n")
		fmt.Fprintf(w, "    |              |--|\n")
		fmt.Fprintf(w, "    |                 |\n")
		fmt.Fprintf(w, "    |                 |\n")
	} else {
		fmt.Fprintf(w, "    |                ||\n")
		fmt.Fprintf(w, "    |          zu    ||\n")
		fmt.Fprintf(w, "    |                ||\n")
		fmt.Fprintf(w, "    |                 |\n")
		fmt.Fprintf(w, "    |                 |\n")
	}
	if currentWindowStatus.BakerStrFensterRe == true {
		fmt.Fprintf(w, "    |              |--|\n")
		fmt.Fprintf(w, "     /      offen  |  |\n")
		fmt.Fprintf(w, "    /              |--|\n")
		fmt.Fprintf(w, "    |                 |\n")
		fmt.Fprintf(w, "#===#=================#\n")
	} else {
		fmt.Fprintf(w, "    |                ||\n")
		fmt.Fprintf(w, "     /         zu    ||\n")
		fmt.Fprintf(w, "    /                ||\n")
		fmt.Fprintf(w, "    |                 |\n")
		fmt.Fprintf(w, "#===#=================#\n")
	}
}
