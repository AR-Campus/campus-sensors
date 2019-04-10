package sensors

import (
	"encoding/json"
	"fmt"
)

// type SensorInfo struct {
// 	device_eui string
// 	payload    string
// 	parsed     []struct {
// 		data map[string]string
// 	}
// 	// Hier Struktur von FireFly json
// }

type SensorInfo struct {
	Device_EUI string                 `json:"device_eui"`
	RawPayload string                 `json:"raw_payload"`
	Parsed     map[string]interface{} `json:"parsed"`
}

type FireFlyPackage struct {
	Packets []SensorInfo `json:"packets"`
}

// func ConvertInfos(rawFireFlyData string) []SensorInfo {
// 	var sensorJson []SensorInfo
// 	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
// 	if err != nil {
// 		fmt.Println("Json Unmarshal failed:", err)
// 		return make([]SensorInfo, 0)
// 	}
// 	// fmt.Println("Len of Packets", len(sensorJson))
// 	var result []SensorInfo
// 	for _, entry := range sensorJson {
// 		result = append(result, entry)
// 		fmt.Println("entry", result)
// 	}
// 	return result
// }

func ConvertInfos(rawFireFlyData string) []SensorInfo {
	var sensorJson FireFlyPackage
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		fmt.Println("Json Unmarshal failed:", err)
		return make([]SensorInfo, 0)
	}
	// fmt.Println("Len of Packets", len(sensorJson))
	var result []SensorInfo
	for _, entry := range sensorJson.Packets {
		result = append(result, entry)
		// fmt.Println("entry", result)
	}
	fmt.Println("Test entrys of parsed_Map:", result[6].Parsed)
	return result
}

// // `&myStoredVariable` is the address of the variable we want to store our
// // parsed data in
// var loraData []LoraData
//
// json.Unmarshall([]byte(loraDatenRAW), &loraData)
//
// fmt.Printf(loraData)
//...
