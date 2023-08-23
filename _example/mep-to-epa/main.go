package main

import (
	"fmt"

	"github.com/ringsaturn/aqi/epa"
	goaqi "github.com/ringsaturn/go-aqi"
)

func main() {
	algo := epa.Algo{}

	var inputs = make([]*goaqi.Var, 0)
	pm25_1h := &goaqi.Var{P: goaqi.PM2_5_1H, Value: 16}
	pm10_1h := &goaqi.Var{P: goaqi.PM10_1H, Value: 88}
	co_1h := &goaqi.Var{P: goaqi.CO_1H, Value: 0.2}
	so2_1h := &goaqi.Var{P: goaqi.SO2_1H, Value: 3}
	no2_1h := &goaqi.Var{P: goaqi.NO2_1H, Value: 3}
	o3_1h := &goaqi.Var{P: goaqi.O3_1H, Value: 3}

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
