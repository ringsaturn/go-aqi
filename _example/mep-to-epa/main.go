package main

import (
	"fmt"

	"github.com/ringsaturn/aqi"
	"github.com/ringsaturn/aqi/epa"
)

func main() {
	algo := epa.Algo{}

	var inputs = make([]*aqi.Var, 0)
	pm25_1h := &aqi.Var{P: aqi.Pollutant_PM2_5_1H, Value: 16}
	pm10_1h := &aqi.Var{P: aqi.Pollutant_PM10_1H, Value: 88}
	co_1h := &aqi.Var{P: aqi.Pollutant_CO_1H, Value: 0.2}
	so2_1h := &aqi.Var{P: aqi.Pollutant_SO2_1H, Value: 3}
	no2_1h := &aqi.Var{P: aqi.Pollutant_NO2_1H, Value: 3}
	o3_1h := &aqi.Var{P: aqi.Pollutant_O3_1H, Value: 3}

	inputs = append(inputs, pm25_1h)
	inputs = append(inputs, pm10_1h)
	inputs = append(inputs, co_1h.MgPerM3ToPPM())
	inputs = append(inputs, so2_1h.MiuGPerM3ToPPB())
	inputs = append(inputs, no2_1h.MiuGPerM3ToPPB())
	inputs = append(inputs, o3_1h.MiuGPerM3ToPPM())

	aqi, primaryPollutant, err := algo.Calc(inputs...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v\n", aqi, primaryPollutant)
	// Output: 67 [PM10_1H]
}
