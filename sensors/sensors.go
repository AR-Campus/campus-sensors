package sensors

import (
	"encoding/json"
	"log"
)

type SensorInfo struct {
	Device_EUI string                   `json:"device_eui"`
	GWRX       []map[string]interface{} `json:"gwrx"`
	Parsed     map[string]interface{}   `json:"parsed"`
}

type SensorData struct {
	DeviceID     string                 `json:"device_id"`
	DeviceType   string                 `json:"device_type"`
	Time         string                 `json:"time"`
	SensorValues map[string]interface{} `json:"sensor_values"`
}

type FireFlyPackage struct {
	Packets []SensorInfo `json:"packets"`
}

type FireFlySingle struct {
	Packet SensorInfo `json:"up_packet"`
}

// type WindowContactsStatus struct {
// 	BakerStrFensterLi bool `json:"BakerStr-Fenster-Li"`
// 	BakerStrFensterRe bool `json:"BakerStr-Fenster-Re"`
// 	KuecheFensterLi   bool `json:"Kueche-Fenster-Li"`
// 	KuecheFensterRe   bool `json:"Kueche-Fenster-Re"`
// }

func corelateDevIDtoName(deviceID string) string {
	switch deviceID {
	case "A81758FFFE031A09":
		return "KuecheTempHumidLicht"
	case "0E7E34643331041C":
		return "KuecheFensterLi"
	case "0E7E34643331041D":
		return "KuecheFensterRe"
	case "0E7E346433310415":
		return "BakerStrFensterLi"
	case "0E7E346433310418":
		return "BakerStrFensterRe"
		// case "":
		// 	return ""
		// case "":
		// 	return ""
		// case "":
		// 	return ""
	}
	return "no-specified-Case-Matched"
}

func ConvertSensorType(sensorInfoTypeArray []SensorInfo) []SensorData {
	result := make([]SensorData, len(sensorInfoTypeArray))
	for i, entry := range sensorInfoTypeArray {
		// time, _ := time.Parse(time.RFC3339, entry.GWRX[0]["time"].(string))
		result[i] = SensorData{DeviceID: corelateDevIDtoName(entry.Device_EUI), DeviceType: "", Time: entry.GWRX[0]["time"].(string), SensorValues: entry.Parsed}
		switch result[i].DeviceID {
		case "KuecheTempHumidLicht":
			result[i].DeviceType = "KuecheKombi"
		case "KuecheFensterLi":
			result[i].DeviceType = "FensterKontakt"
		case "KuecheFensterRe":
			result[i].DeviceType = "FensterKontakt"
		case "BakerStrFensterLi":
			result[i].DeviceType = "FensterKontakt"
		case "BakerStrFensterRe":
			result[i].DeviceType = "FensterKontakt"
			// case "":
			// 	return ""
			// case "":
			// 	return ""
			// case "":
			// 	return ""
		default:
			result[i].DeviceType = "NotClassified"
		}
	}
	return result
}

func ConvertInfos(rawFireFlyData string) []SensorData {
	log.Print("Converting many sensor values.")
	var sensorJson FireFlyPackage
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorData, 0)
	}
	result := ConvertSensorType(sensorJson.Packets)
	log.Printf("Converted %v Sensorvalues.", len(result))
	return result
}

func ConvertSingle(rawFireFlyData string) []SensorData {
	log.Print("Converting a single sensor value")
	var sensorJson FireFlySingle
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorData, 0)
	}
	result := ConvertSensorType([]SensorInfo{sensorJson.Packet})
	log.Printf("Converted %v Sensorvalues.", len(result))
	return result
}

// func windowContactSensorsUpdate(singleSensorData SensorData, currentWindowStatus WindowContactsStatus) {
// 	switch singleSensorData.DeviceID {
// 	case "BakerStr-Fenster-Li":
// 		currentWindowStatus.BakerStrFensterLi = singleSensorData.SensorValues["ReedSensor"].(bool)
// 	case "BakerStr-Fenster-Re":
// 		currentWindowStatus.BakerStrFensterRe = singleSensorData.SensorValues["ReedSensor"].(bool)
// 	case "Kueche-Fenster-Li":
// 		currentWindowStatus.KuecheFensterLi = singleSensorData.SensorValues["ReedSensor"].(bool)
// 	case "Kueche-Fenster-Re":
// 		currentWindowStatus.KuecheFensterRe = singleSensorData.SensorValues["ReedSensor"].(bool)
// 	}
// }
