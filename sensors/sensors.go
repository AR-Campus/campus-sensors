package sensors

import (
	"encoding/json"
	"log"
)

type SensorInfo struct {
	Device_EUI string                 `json:"device_eui"`
	RawPayload string                 `json:"raw_payload"`
	Parsed     map[string]interface{} `json:"parsed"`
}

type FireFlyPackage struct {
	Packets []SensorInfo `json:"packets"`
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
