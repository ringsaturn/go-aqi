// Package eea is for AQI calculation based on European Environment Agency (EEA) standard.
//
// Rules:
//
//	+------------------------------------+-------------------------------------------------------------------+
//	| Pollutant                          | Index level                                                       |
//	|                                    | (based on pollutant concentrations in µg/m3)                      |
//	|                                    +-------+---------+----------+---------+-----------+----------------+
//	|                                    | Good  | Fair    | Moderate | Poor    | Very poor | Extremely poor |
//	+====================================+=======+=========+==========+=========+===========+================+
//	| Particles less than 2.5 µm (PM2.5) | 0-10  | 10-20   | 20-25    | 25-50   | 50-75     | 75-800         |
//	+------------------------------------+-------+---------+----------+---------+-----------+----------------+
//	| Particles less than 10 µm (PM10)   | 0-20  | 20-40   | 40-50    | 50-100  | 100-150   | 150-1200       |
//	+------------------------------------+-------+---------+----------+---------+-----------+----------------+
//	| Nitrogen dioxide (NO2)             | 0-40  | 40-90   | 90-120   | 120-230 | 230-340   | 340-1000       |
//	+------------------------------------+-------+---------+----------+---------+-----------+----------------+
//	| Ozone (O3)                         | 0-50  | 50-100  | 100-130  | 130-240 | 240-380   | 380-800        |
//	+------------------------------------+-------+---------+----------+---------+-----------+----------------+
//	| Sulphur dioxide (SO2)              | 0-100 | 100-200 | 200-350  | 350-500 | 500-750   | 750-1250       |
//	+------------------------------------+-------+---------+----------+---------+-----------+----------------+
//
// Please Notes that:
//
//	> For NO2, O3 and SO2, hourly concentrations are fed into the calculation of
//	> the index.
//	>
//	> For PM10 and PM2.5, the 24-hour running means for the past 24 hours are
//	> fed into the calculation of the index. A 24-hour running mean will be
//	> calculated if there are values for at least 18 out of the 24 hours.
//
// References:
//
//  1. https://environment.ec.europa.eu/topics/air_en
//  2. https://www.eea.europa.eu/highlights/european-air-quality-index-current
//  3. https://www.eea.europa.eu/themes/air/air-quality-index
//  4. https://www.eea.europa.eu/highlights/european-air-quality-index-current
//  5. https://airindex.eea.europa.eu/Map/AQI/Viewer/
package eea

import (
	"fmt"
	"image/color"

	goaqi "github.com/ringsaturn/go-aqi"
)

const Standard = goaqi.AQISTANDARD_EU

var tables = map[goaqi.Pollutant][]float64{
	goaqi.AQI:       {0, 50, 100, 150, 200, 300, 400, 500},
	goaqi.SO2_1H:    {0, 150, 500, 650, 800},                    // μg/m3
	goaqi.NO2_1H:    {0, 100, 200, 700, 1200, 2340, 3090, 3840}, // μg/m3
	goaqi.O3_1H:     {0, 160, 200, 300, 400, 800, 1000, 1200},   // μg/m3
	goaqi.PM10_1H:   {0, 50, 150, 250, 350, 420, 500, 600},      // μg/m3
	goaqi.PM10_24H:  {0, 50, 150, 250, 350, 420, 500, 600},      // μg/m3
	goaqi.PM2_5_1H:  {0, 35, 75, 115, 150, 250, 350, 500},       // μg/m3
	goaqi.PM2_5_24H: {0, 35, 75, 115, 150, 250, 350, 500},       // μg/m3
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

var levelToColor = map[AQILevel]*color.RGBA{
	LEVEL1: {R: 104, G: 233, B: 226}, // Good
	LEVEL2: {R: 95, G: 193, B: 162},  // Fair
	LEVEL3: {R: 238, G: 224, B: 85},  // Moderate
	LEVEL4: {R: 241, G: 84, B: 76},   // Poor
	LEVEL5: {R: 131, G: 27, B: 45},   // Very poor
	LEVEL6: {R: 106, G: 40, B: 104},  // Extremely poor
}

var levelToDesc = map[AQILevel]string{
	LEVEL1: "Good",
	LEVEL2: "Fair",
	LEVEL3: "Moderate",
	LEVEL4: "Poor",
	LEVEL5: "Very poor",
	LEVEL6: "Extremely poor",
}

type Algo struct{}

func (a *Algo) Name() string {
	return "eea"
}

// Calc is func for realtime AQI report computing.
func (a *Algo) Calc(pollutantVars ...*goaqi.Var) (int, []goaqi.Pollutant, error) {
	var (
		results = make(map[goaqi.Pollutant]int)
		maxAQI  int
	)

	for _, pollutantVar := range pollutantVars {
		pollutantIndexRange, ok := tables[pollutantVar.P]
		if !ok {
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
			if pollutantVar.Value > pollutantIndexRange[len(pollutantIndexRange)-1] {
				return 500, nil
			}
			iaqiLo, iaqiHi, pLo, pHi, err := goaqi.GetRanges(pollutantVar.Value, pollutantIndexRange, tables[goaqi.AQI])
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
	rgba, ok := levelToColor[a.AQIToLevel(aqi)]
	if !ok {
		return nil, fmt.Errorf("unknown aqi level for color")
	}
	return rgba, nil
}

func (a *Algo) AQIToDesc(aqi int) (string, error) {
	desc, ok := levelToDesc[a.AQIToLevel(aqi)]
	if !ok {
		return "", fmt.Errorf("unknown aqi level for desc")
	}
	return desc, nil
}
