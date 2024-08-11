package goaqi

type Pollutant int32

const (
	UNKNOWN   Pollutant = 0  // Unknown
	AQI       Pollutant = 1  // Air Quality Index
	O3_1H     Pollutant = 10 // Ozone 1 hour
	O3_8H     Pollutant = 11 // Ozone 8 hour
	PM2_5_1H  Pollutant = 20 // PM2.5 1 hour
	PM2_5_24H Pollutant = 21 // PM2.5 24 hour
	PM10_1H   Pollutant = 30 // PM10 1 hour
	PM10_24H  Pollutant = 31 // PM10 24 hour
	SO2_1H    Pollutant = 40 // Sulfur Dioxide 1 hour
	SO2_24H   Pollutant = 41 // Sulfur Dioxide 24 hour
	NO2_1H    Pollutant = 50 // Nitrogen Dioxide 1 hour
	NO2_24H   Pollutant = 51 // Nitrogen Dioxide 24 hour
	CO_1H     Pollutant = 60 // Carbon Monoxide 1 hour
	CO_8H     Pollutant = 61 // Carbon Monoxide 8 hour
	CO_24H    Pollutant = 62 // Carbon Monoxide 24 hour
)

type AQIStandard int32

const (
	AQISTANDARD_UNSPECIFIED AQIStandard = 0 // Unspecified
	AQISTANDARD_US          AQIStandard = 1 // US AQI
	AQISTANDARD_CN          AQIStandard = 2 // China AQI
	AQISTANDARD_EU          AQIStandard = 3 // EU AQI
)
