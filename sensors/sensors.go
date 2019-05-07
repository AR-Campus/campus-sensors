package sensors

import (
	"encoding/json"
	"log"
	"time"
)

type SensorInfo struct {
	Device_EUI string                   `json:"device_eui"`
	GWRX       []map[string]interface{} `json:"gwrx"`
	Parsed     map[string]interface{}   `json:"parsed"`
}

type SensorData struct {
	DeviceID     string                 `json:"device_id"`
	Time         time.Time              `json:"time"`
	SensorValues map[string]interface{} `json:"sensor_values"`
}

type FireFlyPackage struct {
	Packets []SensorInfo `json:"packets"`
}

type FireFlySingle struct {
	Packet SensorInfo `json:"up_packet"`
}

func ConvertSensorType(sensorInfoTypeArray []SensorInfo) []SensorData {
	result := make([]SensorData, len(sensorInfoTypeArray))
	for i, entry := range sensorInfoTypeArray {
		time, _ := time.Parse(time.RFC3339, entry.GWRX[0]["time"].(string))
		result[i] = SensorData{DeviceID: entry.Device_EUI, Time: time, SensorValues: entry.Parsed}
	}
	return result
}

func ConvertInfos(rawFireFlyData string) []SensorData {
	log.Print("Raw data", rawFireFlyData)
	var sensorJson FireFlyPackage
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorData, 0)
	}
	return ConvertSensorType(sensorJson.Packets)
}

func ConvertSingle(rawFireFlyData string) []SensorData {
	log.Print("Raw data", rawFireFlyData)
	var sensorJson FireFlySingle
	err := json.Unmarshal([]byte(rawFireFlyData), &sensorJson)
	if err != nil {
		log.Println("Json Unmarshal failed:", err)
		return make([]SensorData, 0)
	}
	return ConvertSensorType([]SensorInfo{sensorJson.Packet})
}
