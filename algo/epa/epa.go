// Package epa is impl for EPA 454/B-18-007
//
// Offical Doc
// https://nepis.epa.gov/Exe/ZyPURL.cgi?Dockey=P100W5UG.TXT
package epa

import (
	"fmt"

	"github.com/ringsaturn/aqi"
	"github.com/ringsaturn/aqi/algo"
)

var Tables = map[aqi.Pollutant][]float32{
	aqi.Pollutant_AQI:       {0, 50, 100, 150, 200, 300, 400, 500},
	aqi.Pollutant_CO_8H:     {0, 4.4, 9.4, 12.4, 15.4, 30.4, 40.4, 50.4},      // ppm
	aqi.Pollutant_SO2_1H:    {0, 35, 75, 185, 304, 604, 804, 1004},            // ppb
	aqi.Pollutant_NO2_1H:    {0, 53, 100, 360, 649, 1249, 1649, 2049},         // ppb
	aqi.Pollutant_O3_8H:     {0, 0.054, 0.070, 0.085, 0.105, 0.2},             // ppm
	aqi.Pollutant_O3_1H:     {0, 0, 0.125, 0.164, 0.204, 0.404, 0.504, 0.604}, // ppm
	aqi.Pollutant_PM2_5_1H:  {0, 12, 35.4, 55.4, 150.4, 250.4, 350.4, 500.4},  // μg/m3
	aqi.Pollutant_PM2_5_24H: {0, 12, 35.4, 55.4, 150.4, 250.4, 350.4, 500.4},  // μg/m3
	aqi.Pollutant_PM10_1H:   {0, 54, 154, 254, 354, 424, 504, 604},            // μg/m3
	aqi.Pollutant_PM10_24H:  {0, 54, 154, 254, 354, 424, 504, 604},            // μg/m3
}

type Algo struct {
	FailedWhenNotSupported bool
}

func (a *Algo) Name() string {
	return "epa"
}

// Calc is func for realtime AQI report computing.
func (a *Algo) Calc(pollutantVars ...aqi.Var) (int, []aqi.Pollutant, error) {
	var (
		results = make(map[aqi.Pollutant]int)
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
		if pollutantVar.P == aqi.Pollutant_O3_8H && pollutantVar.Value > 0.2 {
			continue
		}

		// 1-hour SO 2 values do not define higher AQI values (≥ 200).
		// AQI values of 200 or greater are calculated with 24-hour SO 2 concentrations.
		if pollutantVar.P == aqi.Pollutant_SO2_1H && pollutantVar.Value > 304 {
			continue
		}

		iaqiLo, iaqiHi, pLo, pHi, err := algo.GetRanges(pollutantVar.Value, pollutantIndexrange, Tables[aqi.Pollutant_AQI])
		if err != nil {
			return 0, nil, err
		}

		aqi, err := algo.CalcViaHiLo(pollutantVar.Value, iaqiLo, iaqiHi, pLo, pHi)
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
	primaryPollutants := make([]aqi.Pollutant, 0)
	for pollutant, value := range results {
		if value == maxAQI {
			primaryPollutants = append(primaryPollutants, pollutant)
		}
	}
	return maxAQI, primaryPollutants, nil
}
