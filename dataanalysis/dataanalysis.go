package dataanalysis

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"db-training.de/campus-sensors/sensors"
)

type WindowContactsStatus struct {
	BakerStrFensterLi bool `json:"BakerStr-Fenster-Li"`
	BakerStrFensterRe bool `json:"BakerStr-Fenster-Re"`
	KuecheFensterLi   bool `json:"Kueche-Fenster-Li"`
	KuecheFensterRe   bool `json:"Kueche-Fenster-Re"`
}

type SensorFlowPerHour struct {
	HourTimeData             string `json:"Hour-Time-Data"`
	QuantityOfSensorPackages int    `json:"Quantity-Of-Sensor-Packages"`
}

type SensorFlowPerDay struct {
	DayTimeData              string `json:"Hour-Time-Data"`
	QuantityOfSensorPackages int    `json:"Quantity-Of-Sensor-Packages"`
}

// type SensorFlowPerMonth struct {
// 	MonthTimeData            string `json:"Hour-Time-Data"`
// 	QuantityOfSensorPackages int    `json:"Quantity-Of-Sensor-Packages"`
// }

type PackagesPerSensorCount struct {
	KuecheTempHumidLicht int64 `json:"Kueche-Temp-Humid-Licht"`
	BakerStrFensterLi    int64 `json:"BakerStr-Fenster-Li"`
	BakerStrFensterRe    int64 `json:"BakerStr-Fenster-Re"`
	KuecheFensterLi      int64 `json:"Kueche-Fenster-Li"`
	KuecheFensterRe      int64 `json:"Kueche-Fenster-Re"`
}

func WindowContactSensorsUpdate(singleSensorData sensors.SensorData, currentWindowStatus WindowContactsStatus) WindowContactsStatus {
	switch singleSensorData.DeviceID {
	case "BakerStr-Fenster-Li":
		currentWindowStatus.BakerStrFensterLi = singleSensorData.SensorValues["ReedSensor"].(bool)
	case "BakerStr-Fenster-Re":
		currentWindowStatus.BakerStrFensterRe = !singleSensorData.SensorValues["ReedSensor"].(bool)
	case "Kueche-Fenster-Li":
		currentWindowStatus.KuecheFensterLi = !singleSensorData.SensorValues["ReedSensor"].(bool)
	case "Kueche-Fenster-Re":
		currentWindowStatus.KuecheFensterRe = singleSensorData.SensorValues["ReedSensor"].(bool)
	}
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
	case "Kueche-Temp-Humid-Licht":
		currentSensorQuantities.KuecheTempHumidLicht++
	case "BakerStr-Fenster-Li":
		currentSensorQuantities.BakerStrFensterLi++
	case "BakerStr-Fenster-Re":
		currentSensorQuantities.BakerStrFensterRe++
	case "Kueche-Fenster-Li":
		currentSensorQuantities.KuecheFensterLi++
	case "Kueche-Fenster-Re":
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
	fmt.Fprintf(w, "|  |      KÜCHE       |\n")
	fmt.Fprintf(w, "|__|                  |\n")
	fmt.Fprintf(w, "|_                    |\n")
	fmt.Fprintf(w, " /                    |\n")
	fmt.Fprintf(w, "/                     |\n")
	if currentWindowStatus.KuecheFensterLi == false {
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
	if currentWindowStatus.KuecheFensterRe == false {
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