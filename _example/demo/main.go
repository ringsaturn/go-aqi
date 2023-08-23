package main

import (
	"fmt"

	goaqi "github.com/ringsaturn/go-aqi"
	"github.com/ringsaturn/go-aqi/mep"
)

func main() {
	algo := &mep.Algo{}
	inputs := []*goaqi.Var{
		{
			P:     goaqi.PM2_5_1H,
			Value: 16,
		},
		{
			P:     goaqi.PM10_1H,
			Value: 88,
		},
		{
			P:     goaqi.CO_1H,
			Value: 0.2,
		},
		{
			P:     goaqi.SO2_1H,
			Value: 3,
		},
		{
			P:     goaqi.NO2_1H,
			Value: 11,
		},
		{
			P:     goaqi.O3_1H,
			Value: 75,
		},
	}
	aqi, primaryPollutant, err := algo.Calc(inputs...)
	if err != nil {
		panic(err)
	}
	levelDesc, err := algo.AQIToDesc(aqi)
	if err != nil {
		panic(err)
	}
	fmt.Printf("aqi=%v as level=%v with primary pollutant as %v\n", aqi, levelDesc, primaryPollutant)
}
