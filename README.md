# Compute AQI in Go [![ci](https://github.com/ringsaturn/aqi/actions/workflows/ci.yml/badge.svg)](https://github.com/ringsaturn/aqi/actions/workflows/ci.yml)

```bash
go install github.com/ringsaturn/aqi
```

Usage:

```go
package main

import (
	"fmt"

	"github.com/ringsaturn/aqi"
	"github.com/ringsaturn/aqi/algo/mep"
)

func main() {
	algo := &mep.Algo{}
	inputs := []aqi.Var{
		{
			P:     aqi.Pollutant_PM2_5_1H,
			Value: 16,
		},
		{
			P:     aqi.Pollutant_PM10_1H,
			Value: 88,
		},
		{
			P:     aqi.Pollutant_CO_1H,
			Value: 0.2,
		},
		{
			P:     aqi.Pollutant_SO2_1H,
			Value: 3,
		},
		{
			P:     aqi.Pollutant_NO2_1H,
			Value: 11,
		},
		{
			P:     aqi.Pollutant_O3_1H,
			Value: 75,
		},
	}
	aqi, primaryPollutant, err := algo.Calc(inputs...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("aqi=%v with primary pollutant as %v\n", aqi, primaryPollutant)
}
```

NOTE: Currently the algo impl is based on the different standard files and
different AQI Standard use different unit system.
Please ensure the input value has been converted to the algo expect unit system.
