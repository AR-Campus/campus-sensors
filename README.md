# campus-sensors
Server and hosting for the LORA-Sensors in the CAMPUS

I. General Info

- Sensors send Packages to FireFlyIOT Server
- This Server gets new Packages via Push from FireFlyIOT Server
- On startup this Server downloads a specified amount of "last Sensor Packages"
- These Sensor Packages are stored in Memory in a DataBase liked by the Sensors UID
- This Server provides a WebPage with graphical Representations of those SensorDatas
