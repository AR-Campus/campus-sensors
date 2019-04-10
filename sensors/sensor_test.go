package sensors

import (
	"testing"

	"github.com/laliluna/expectations"
)

func TestTest(t *testing.T) {

}

func TestInputTestValues(t *testing.T) {
	et := expectations.NewT(t)
	et.ExpectThat(singelFireFlyData()).String()
	et.ExpectThat(singelFireFlyPackage()).String()
	et.ExpectThat(rawFireFlyPackage()).String()
}

func TestPasedSensorData(t *testing.T) {
	et := expectations.NewT(t)
	var result []SensorInfo = ConvertInfos(rawFireFlyPackage())
	// fmt.Println("Input String length", len(rawSensorDataPackets))
	et.ExpectThat(len(result)).DoesNotEqual(0)

	result = ConvertInfos(`[{"SomeJibberish"}]`)

	et.ExpectThat(len(result)).Equals(0)

}
