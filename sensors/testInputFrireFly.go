package sensors

func singelFireFlyData() string {
	dataString := `[{
		"uid":"Your-Personal-Device-UID",
		"spreading_factor":7,
		"size":26,
		"received_at":"some-Date-Time-Format",
		"port":5,
		"payload_encrypted":false,
		"payload":"Some-HexaDeciaml-Value",
		"parsed":{
			"VDDID":001,
			"VDD":3.1415,
			"TempID":1,
			"Temp":180,
			"MotionID":2,
			"Motion":42,
			"LightID":21,
			"Light":42,
			"HumidityID":21,
			"Humidity":42},
		"mtype":"soem-String",
		"modu":"LORA",
		"mic_pass":true,
		"gwrx":[{
			"tmst":31415926,
			"time":"another-Date-Time-Format",
			"srv_rcv_time":1234567891,
			"rssi":-42,
			"lsnr":10.0,
			"gweui":"Another-HexaDecimal-Value"}],
		"freq":868.1,
		"fopts":"",
		"fcnt":12345,
		"device_eui":"Your-Device-EUI",
		"datr":"HexaDecimalValue",
		"codr":"1/3",
		"bandwidth":123,
		"ack":false}]`
	return dataString
}

func singelFireFlyPackage() string {
	dataString := `{
	"packets":[{
    "uid":"Your-Personal-Device-UID",
		"spreading_factor":7,
		"size":26,
		"received_at":"some-Date-Time-Format",
		"port":5,
		"payload_encrypted":false,
		"payload":"Some-HexaDeciaml-Value",
		"parsed":{
			"VDDID":001,
			"VDD":3.1415,
			"TempID":1,
			"Temp":180,
			"MotionID":2,
			"Motion":42,
			"LightID":21,
			"Light":42,
			"HumidityID":21,
			"Humidity":42},
		"mtype":"soem-String",
		"modu":"LORA",
		"mic_pass":true,
		"gwrx":[{
			"tmst":31415926,
			"time":"another-Date-Time-Format",
			"srv_rcv_time":1234567891,
			"rssi":-42,
			"lsnr":10.0,
			"gweui":"Another-HexaDecimal-Value"}],
		"freq":868.1,
		"fopts":"",
		"fcnt":12345,
		"device_eui":"Your-Device-EUI",
		"datr":"HexaDecimalValue",
		"codr":"1/3",
		"bandwidth":123,
		"ack":false}],
	"more":true}`
	return dataString
}

func rawFireFlyPackage() string {
	return `{"packets":[
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:31:29.816557","port":5,"payload_encrypted":false,"payload":"0100D6022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.4,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1665368427,"time":"2019-04-05T14:31:28.993660Z","srv_rcv_time":1554474689055466,"rssi":-73,"lsnr":7.5,"gweui":"00000008004A0826"}],"freq":868.5,"fopts":"","fcnt":112594,"device_eui":"Your-Device-EUI","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:30:29.662156","port":5,"payload_encrypted":false,"payload":"0100D7022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.5,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1605371547,"time":"2019-04-05T14:30:28.972066Z","srv_rcv_time":1554474629007541,"rssi":-72,"lsnr":9.8,"gweui":"00000008004A0826"}],"freq":868.1,"fopts":"","fcnt":112593,"device_eui":"Your-Device-EUI","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:29:30.255333","port":5,"payload_encrypted":false,"payload":"0100D7022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.5,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1545374667,"time":"2019-04-05T14:29:28.969711Z","srv_rcv_time":1554474569147433,"rssi":-77,"lsnr":7.5,"gweui":"00000008004A0826"}],"freq":868.5,"fopts":"","fcnt":112592,"device_eui":"Your-Device-EUI","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:28:30.547328","port":5,"payload_encrypted":false,"payload":"0100D7022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.5,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1485377787,"time":"2019-04-05T14:28:28.965336Z","srv_rcv_time":1554474509187409,"rssi":-73,"lsnr":11.2,"gweui":"00000008004A0826"}],"freq":868.1,"fopts":"","fcnt":112591,"device_eui":"Your-Device-EUI","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:27:30.433267","port":5,"payload_encrypted":false,"payload":"0100D7022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.5,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1425380899,"time":"2019-04-05T14:27:28.975782Z","srv_rcv_time":1554474449067503,"rssi":-78,"lsnr":7.8,"gweui":"00000008004A0826"}],"freq":868.5,"fopts":"","fcnt":112590,"device_eui":"Your-Device-EUI","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:26:29.790622","port":5,"payload_encrypted":false,"payload":"0100D7022A0400120500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.5,"MotionID":5,"Motion":0,"LightID":64,"Light":18,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1365384019,"time":"2019-04-05T14:26:28.969453Z","srv_rcv_time":1554474389308827,"rssi":-75,"lsnr":7.0,"gweui":"00000008004A0826"}],"freq":868.5,"fopts":"","fcnt":112589,"device_eui":"A81758FFFE031A09","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":17,"received_at":"2019-04-05T14:25:38.706860","port":9,"payload_encrypted":false,"payload":"DE02000F","parsed":{"reserved":true,"Tamper":true,"ReedSensor":false,"OpeningCounter":512,"BatteryStatus":true},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1313672179,"time":"2019-04-05T14:25:37.260902Z","srv_rcv_time":1554474337347287,"rssi":-89,"lsnr":9.2,"gweui":"00000008004A0826"}],"freq":867.1,"fopts":"","fcnt":2086,"device_eui":"0E7E34643331041C","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":26,"received_at":"2019-04-05T14:25:29.455256","port":5,"payload_encrypted":false,"payload":"0100D6022A0400130500070DF1","parsed":{"VDDID":112,"VDD":3.569,"TempID":16,"Temp":21.4,"MotionID":5,"Motion":0,"LightID":64,"Light":19,"HumidityID":2,"Humidity":42},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1305387147,"time":"2019-04-05T14:25:28.975071Z","srv_rcv_time":1554474328996584,"rssi":-75,"lsnr":10.2,"gweui":"00000008004A0826"}],"freq":868.1,"fopts":"","fcnt":112588,"device_eui":"A81758FFFE031A09","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":17,"received_at":"2019-04-05T14:24:41.191445","port":9,"payload_encrypted":false,"payload":"E0020023","parsed":{"reserved":true,"Tamper":false,"ReedSensor":false,"OpeningCounter":512,"BatteryStatus":false},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1256127596,"time":"2019-04-05T14:24:39.725150Z","srv_rcv_time":1554474279827179,"rssi":-101,"lsnr":5.8,"gweui":"00000008004A0826"}],"freq":867.7,"fopts":"","fcnt":2118,"device_eui":"0E7E346433310415","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":false},
  {"uid":"Your-Personal-Device-UID","spreading_factor":7,"size":17,"received_at":"2019-04-05T14:24:36.068549","port":9,"payload_encrypted":false,"payload":"DE020050","parsed":{"reserved":true,"Tamper":true,"ReedSensor":false,"OpeningCounter":512,"BatteryStatus":true},"mtype":"unconfirmed_data_up","modu":"LORA","mic_pass":true,"gwrx":[{"tmst":1251016740,"time":"2019-04-05T14:24:34.608326Z","srv_rcv_time":1554474274707256,"rssi":-87,"lsnr":6.8,"gweui":"00000008004A0826"}],"freq":867.5,"fopts":"","fcnt":2201,"device_eui":"0E7E346433310418","datr":"SF7BW125","codr":"4/5","bandwidth":125,"ack":true}],
  "more":true}`
}
