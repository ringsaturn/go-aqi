// Package epa is impl for EPA 454/B-18-007
//
// Offical Doc
// https://nepis.epa.gov/Exe/ZyPURL.cgi?Dockey=P100W5UG.TXT
package epa

import (
	"fmt"
	"image/color"

	goaqi "github.com/ringsaturn/go-aqi"
)

var Tables = map[goaqi.Pollutant][]float64{
	goaqi.AQI:       {0, 50, 100, 150, 200, 300, 400, 500},
	goaqi.CO_8H:     {0, 4.4, 9.4, 12.4, 15.4, 30.4, 40.4, 50.4},      // ppm
	goaqi.SO2_1H:    {0, 35, 75, 185, 304, 604, 804, 1004},            // ppb
	goaqi.NO2_1H:    {0, 53, 100, 360, 649, 1249, 1649, 2049},         // ppb
	goaqi.O3_8H:     {0, 0.054, 0.070, 0.085, 0.105, 0.2},             // ppm
	goaqi.O3_1H:     {0, 0, 0.125, 0.164, 0.204, 0.404, 0.504, 0.604}, // ppm
	goaqi.PM2_5_1H:  {0, 12, 35.4, 55.4, 150.4, 250.4, 350.4, 500.4},  // μg/m3
	goaqi.PM2_5_24H: {0, 12, 35.4, 55.4, 150.4, 250.4, 350.4, 500.4},  // μg/m3
	goaqi.PM10_1H:   {0, 54, 154, 254, 354, 424, 504, 604},            // μg/m3
	goaqi.PM10_24H:  {0, 54, 154, 254, 354, 424, 504, 604},            // μg/m3
}

type AQILevel int

const (
	LEVEL_UNDEFINE AQILevel = iota
	LEVEL1
	LEVEL2
	LEVEL3
	LEVEL4
	LEVEL5
	LEVEL6
)

var LevelToColor = map[AQILevel]*color.RGBA{
	LEVEL1: {R: 0, G: 228, B: 0},
	LEVEL2: {R: 255, G: 255, B: 0},
	LEVEL3: {R: 255, G: 126, B: 0},
	LEVEL4: {R: 255, G: 0, B: 0},
	LEVEL5: {R: 143, G: 63, B: 151},
	LEVEL6: {R: 126, G: 0, B: 35},
}

var LevelToDesc = map[AQILevel]string{
	LEVEL1: "Good",
	LEVEL2: "Moderate",
	LEVEL3: "Unhealthy for Sensitive Groups",
	LEVEL4: "Unhealthy",
	LEVEL5: "Very Unhealthy",
	LEVEL6: "Hazardous",
}

type Algo struct {
	FailedWhenNotSupported bool
}

func (a *Algo) Name() string {
	return "epa"
}

// Calc is func for realtime AQI report computing.
func (a *Algo) Calc(pollutantVars ...*goaqi.Var) (int, []goaqi.Pollutant, error) {
	var (
		results = make(map[goaqi.Pollutant]int)
		maxAQI  int
	)

	for _, pollutantVar := range pollutantVars {
		pollutantIndexrange, ok := Tables[pollutantVar.P]
		if !ok {
			if a.FailedWhenNotSupported {
				return 0, nil, fmt.Errorf("pollutant %v not supported yet", pollutantVar.P.String())
			}
			// allow input not supported pollutant, just continue
			continue
		}

		// 8-hour O 3 values do not define higher AQI values (≥ 301).
		// AQI values of 301 or higher are calculated with 1-hour O 3 concentrations.
		if pollutantVar.P == goaqi.O3_8H && pollutantVar.Value > 0.2 {
			continue
		}

		// 1-hour SO 2 values do not define higher AQI values (≥ 200).
		// AQI values of 200 or greater are calculated with 24-hour SO 2 concentrations.
		if pollutantVar.P == goaqi.SO2_1H && pollutantVar.Value > 304 {
			continue
		}

		aqi, err := func() (int, error) {
			if pollutantVar.Value > pollutantIndexrange[len(pollutantIndexrange)-1] {
				return 500, nil
			}
			iaqiLo, iaqiHi, pLo, pHi, err := goaqi.GetRanges(pollutantVar.Value, pollutantIndexrange, Tables[goaqi.AQI])
			if err != nil {
				return 0, err
			}
			return goaqi.CalcViaHiLo(pollutantVar.Value, iaqiLo, iaqiHi, pLo, pHi)
		}()
		if err != nil {
			return 0, nil, err
		}

		if aqi > maxAQI {
			maxAQI = aqi
		}
		results[pollutantVar.P] = aqi
	}
	if maxAQI <= 50 {
		return maxAQI, nil, nil
	}
	primaryPollutants := make([]goaqi.Pollutant, 0)
	for pollutant, value := range results {
		if value == maxAQI {
			primaryPollutants = append(primaryPollutants, pollutant)
		}
	}
	return maxAQI, primaryPollutants, nil
}

func (a *Algo) AQIToLevel(aqi int) AQILevel {
	if aqi <= 50 {
		return LEVEL1
	}
	if aqi <= 100 {
		return LEVEL2
	}
	if aqi <= 150 {
		return LEVEL3
	}
	if aqi <= 200 {
		return LEVEL4
	}
	if aqi <= 300 {
		return LEVEL5
	}
	return LEVEL6
}

func (a *Algo) AQIToColor(aqi int) (*color.RGBA, error) {
	rgba, ok := LevelToColor[a.AQIToLevel(aqi)]
	if !ok {
		return nil, fmt.Errorf("unknown aqi level for color")
	}
	return rgba, nil
}

func (a *Algo) AQIToDesc(aqi int) (string, error) {
	desc, ok := LevelToDesc[a.AQIToLevel(aqi)]
	if !ok {
		return "", fmt.Errorf("unknown aqi level for desc")
	}
	return desc, nil
}
