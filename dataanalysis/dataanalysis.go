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
	HourMatrix []string `json:"HourMatrix"`
	FlowMatrix []int    `json:"FlowMatrix"`
}

type SensorFlowPerDay struct {
	DayTimeData              string `json:"DayTimeData"`
	QuantityOfSensorPackages int    `json:"QuantityOfSensorPackages"`
}

type SensorFlowPerDayPackageJson struct {
	DayMatrix  []string `json:"DayMatrix"`
	FlowMatrix []int    `json:"FlowMatrix"`
}

type TemperatureFlowPerHour struct {
	HourTimeData      string  `json:"HourTimeData"`
	TemperatureSum    float64 `json:"TemperatureSum"`
	QuantitiyOfEntrys float64 `json:"QuantitiyOfEntrys"`
}

type TemperatureFlowPerHourPackageJson struct {
	HourMatrix        []string  `json:"HourMatrix"`
	TemperatureMatrix []float64 `json:"TemperatureMatrix"`
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
	// log.Printf("SingleSensorData.Sensorvalues: %v", singleSensorData.SensorValues)
	// log.Printf("SingleSensorData.Sensorvalues[ReedSensor]: %v", singleSensorData.SensorValues["ReedSensor"])
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
	// currentWindowStatus = WindowContactsStatus{BakerStrFensterLi: RandomBool(),
	// 	BakerStrFensterRe: RandomBool(),
	// 	KuecheFensterLi:   RandomBool(),
	// 	KuecheFensterRe:   RandomBool()}
	// log.Printf("WindowStatus: %v", currentWindowStatus)
	// log.Printf("CurrentWindowStatus: %v", currentWindowStatus)
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

func TemperatureFlowHourArrayUpdate(singleSensorData sensors.SensorData, temperaturePackageFlowData []TemperatureFlowPerHour) []TemperatureFlowPerHour {
	timeOfSensor, _ := time.Parse(time.RFC3339, singleSensorData.Time)
	tempOfSensor, _ := singleSensorData.SensorValues["Temp"].(float64)
	// log.Printf("in TempFlowHourUpdate Temp of current Sensor: %v of Type %T", tempOfSensor, tempOfSensor)
	if len(temperaturePackageFlowData) != 0 {
		// log.Printf("from SensorFlowUpdate: SensorFlowHour is : %v", sensorPackageFlowData[0].HourTimeData)
		lastTempHour := len(temperaturePackageFlowData) - 1
		currentTempHour, _ := time.Parse(time.RFC3339, temperaturePackageFlowData[lastTempHour].HourTimeData)
		if timeOfSensor.Before(currentTempHour.Add(time.Hour * 1)) {
			// log.Printf("from SensorFlowUpdate-sensorBefore-true")
			temperaturePackageFlowData[lastTempHour].TemperatureSum += tempOfSensor
			temperaturePackageFlowData[lastTempHour].QuantitiyOfEntrys++
		} else {
			// log.Printf("from SensorFlowUpdate-sensorBefore-false - next hour")
			temperaturePackageFlowData = append(temperaturePackageFlowData, TemperatureFlowPerHour{HourTimeData: singleSensorData.Time, TemperatureSum: singleSensorData.SensorValues["Temp"].(float64), QuantitiyOfEntrys: 1.0})
			log.Printf("Next SensorFlowHour at: %v", temperaturePackageFlowData[1].HourTimeData)
		}
	}
	if len(temperaturePackageFlowData) == 0 {
		temperaturePackageFlowData = append(temperaturePackageFlowData, TemperatureFlowPerHour{HourTimeData: singleSensorData.Time, TemperatureSum: singleSensorData.SensorValues["Temp"].(float64), QuantitiyOfEntrys: 1.0})
		log.Printf("First TempFlowHour added!")
	}
	return temperaturePackageFlowData
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

func ParseSensorFlowPerHourJson(sensorPackageFlowData []SensorFlowPerHour) SensorFlowPerHourPackageJson {
	if len(sensorPackageFlowData) <= 24 {
		input := sensorPackageFlowData
		result := SensorFlowPerHourPackageJson{HourMatrix: make([]string, len(input)), FlowMatrix: make([]int, len(input))}
		for i, entry := range input {
			result.HourMatrix[i] = entry.HourTimeData[11:13] + ":00" // 2019-08-06T14:13:12.746280Z
			result.FlowMatrix[i] = entry.QuantityOfSensorPackages
		}
		return result
	} else {
		input := sensorPackageFlowData[(len(sensorPackageFlowData) - 24):len(sensorPackageFlowData)]
		result := SensorFlowPerHourPackageJson{HourMatrix: make([]string, len(input)), FlowMatrix: make([]int, len(input))}
		for i, entry := range input {
			result.HourMatrix[i] = entry.HourTimeData[11:13] + ":00" // 2019-08-06T14:13:12.746280Z
			result.FlowMatrix[i] = entry.QuantityOfSensorPackages
		}
		return result
	}
}

func ParseTemperatureFlowPerHourJson(temperaturePackageFlowData []TemperatureFlowPerHour) TemperatureFlowPerHourPackageJson {
	numOfDays := 3
	if len(temperaturePackageFlowData) <= (numOfDays * 24) {
		input := temperaturePackageFlowData
		result := TemperatureFlowPerHourPackageJson{HourMatrix: make([]string, len(input)), TemperatureMatrix: make([]float64, len(input))}
		for i, entry := range input {
			result.HourMatrix[i] = entry.HourTimeData[11:13] + ":00" // 2019-08-06T14:13:12.746280Z
			result.TemperatureMatrix[i] = (entry.TemperatureSum / entry.QuantitiyOfEntrys)
		}
		// log.Printf("Current TempFlowJSON: %v", result)
		return result
	} else {
		input := temperaturePackageFlowData[(len(temperaturePackageFlowData) - (numOfDays * 24)):len(temperaturePackageFlowData)]
		result := TemperatureFlowPerHourPackageJson{HourMatrix: make([]string, len(input)), TemperatureMatrix: make([]float64, len(input))}
		for i, entry := range input {
			result.HourMatrix[i] = entry.HourTimeData[11:13] + ":00" // 2019-08-06T14:13:12.746280Z
			result.TemperatureMatrix[i] = (entry.TemperatureSum / entry.QuantitiyOfEntrys)
		}
		return result
	}
}

func ParseSensorFlowPerDayJson(sensorPackageFlowData []SensorFlowPerDay) SensorFlowPerDayPackageJson {
	if len(sensorPackageFlowData) <= 14 {
		input := sensorPackageFlowData
		result := SensorFlowPerDayPackageJson{DayMatrix: make([]string, len(input)), FlowMatrix: make([]int, len(input))}
		for i, entry := range input {
			result.DayMatrix[i] = entry.DayTimeData[8:10] + "." + entry.DayTimeData[5:7] // 2019-08-06T14:13:12.746280Z
			result.FlowMatrix[i] = entry.QuantityOfSensorPackages
		}
		return result
	} else {
		input := sensorPackageFlowData[(len(sensorPackageFlowData) - 14):len(sensorPackageFlowData)]
		result := SensorFlowPerDayPackageJson{DayMatrix: make([]string, len(input)), FlowMatrix: make([]int, len(input))}
		for i, entry := range input {
			result.DayMatrix[i] = entry.DayTimeData[8:10] + "." + entry.DayTimeData[5:7] // 2019-08-06T14:13:12.746280Z
			result.FlowMatrix[i] = entry.QuantityOfSensorPackages
		}
		return result
	}
}

func UpdateAnalysisData(singleSensorData sensors.SensorData, currentWindowStatus WindowContactsStatus, temperatureFlowHourData []TemperatureFlowPerHour, sensorPackageHourFlowData []SensorFlowPerHour, sensorPackageDayFlowData []SensorFlowPerDay, currentSensorQuantities PackagesPerSensorCount) (WindowContactsStatus, []TemperatureFlowPerHour, []SensorFlowPerHour, []SensorFlowPerDay, PackagesPerSensorCount) {
	// log.Printf("From UpdateFunc: now SensorFlowHourArray length is : %v", len(sensorPackageFlowData))
	if singleSensorData.DeviceType == "FensterKontakt" {
		currentWindowStatus = WindowContactSensorsUpdate(singleSensorData, currentWindowStatus)
	} else if singleSensorData.DeviceType == "KuecheKombi" {
		temperatureFlowHourData = TemperatureFlowHourArrayUpdate(singleSensorData, temperatureFlowHourData)
		// log.Printf("Current TempFlowData[]: %v", temperatureFlowHourData)
	}
	sensorPackageHourFlowData = SensorFlowHourArrayUpdate(singleSensorData, sensorPackageHourFlowData)
	sensorPackageDayFlowData = SensorFlowDayArrayUpdate(singleSensorData, sensorPackageDayFlowData)
	currentSensorQuantities = QuantifyPerSensorPackages(singleSensorData, currentSensorQuantities)
	return currentWindowStatus, temperatureFlowHourData, sensorPackageHourFlowData, sensorPackageDayFlowData, currentSensorQuantities
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
