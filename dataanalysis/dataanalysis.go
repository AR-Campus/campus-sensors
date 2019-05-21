package dataanalysis

type windowContactsStatus struct {
	BakerStrFensterLi bool `json:"BakerStr-Fenster-Li"`
	BakerStrFensterRe bool `json:"BakerStr-Fenster-Re"`
	KuecheFensterLi   bool `json:"Kueche-Fenster-Li"`
	KuecheFensterRe   bool `json:"Kueche-Fenster-Re"`
}

// func corelateDevIDtoName(deviceID string) string {
// 	switch deviceID {
// 	case "A81758FFFE031A09":
// 		return "Kueche-Temp-Humid-Licht"
// 	case "0E7E34643331041C":
// 		return "Kueche-Fenster-Li"
// 	case "0E7E34643331041D":
// 		return "Kueche-Fenster-Re"
// 	case "0E7E346433310415":
// 		return "BakerStr-Fenster-Li"
// 	case "0E7E346433310418":
// 		return "BakerStr-Fenster-Re"
// 		// case "":
// 		// 	return ""
// 		// case "":
// 		// 	return ""
// 		// case "":
// 		// 	return ""
// 	}
// 	return "no-specified-Case-Matched"
// }

func windowContactSensorsUpdate(singleSensorData SensorData) {

}
