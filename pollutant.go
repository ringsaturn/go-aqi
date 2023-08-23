package goaqi

type Pollutant int32

const (
	UNKNOWN   Pollutant = 0
	AQI       Pollutant = 1
	O3_1H     Pollutant = 10
	O3_8H     Pollutant = 11
	PM2_5_1H  Pollutant = 20
	PM2_5_24H Pollutant = 21
	PM10_1H   Pollutant = 30
	PM10_24H  Pollutant = 31
	SO2_1H    Pollutant = 40
	SO2_24H   Pollutant = 41
	NO2_1H    Pollutant = 50
	NO2_24H   Pollutant = 51
	CO_1H     Pollutant = 60
	CO_8H     Pollutant = 61
	CO_24H    Pollutant = 62
)
