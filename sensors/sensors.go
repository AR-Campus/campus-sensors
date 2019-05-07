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

type FireFlyPackage struct {
	Packets []SensorInfo `json:"packets"`
}

type FireFlySingle struct {
	Packet SensorInfo `json:"up_packet"`
}

func ConvertInfos(rawFireFlyData string) []SensorInfo {
	log.Print("Raw data", rawFireFlyData)
	var sensorJson FireFlyPackage
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorInfo, 0)
	}
	var result []SensorInfo
	for _, entry := range sensorJson.Packets {
		result = append(result, entry)
	}
	log.Print("Result", result)
	return result
}

func ConvertSingle(rawFireFlyData string) []SensorInfo {
	log.Print("Raw data", rawFireFlyData)
	var sensorJson FireFlySingle
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorInfo, 0)
	}
	var result []SensorInfo
	result = append(result, sensorJson.Packet)
	log.Print("Result", result)
	return result
}
