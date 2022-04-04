// Package mep is for HJ633-2012
//
// Offical Doc
// https://www.mee.gov.cn/ywgz/fgbz/bz/bzwb/jcffbz/201203/t20120302_224166.shtml
package mep

import (
	"fmt"

	"github.com/ringsaturn/aqi"
)

var Tables = map[aqi.Pollutant][]float32{
	aqi.Pollutant_AQI:       {0, 50, 100, 150, 200, 300, 400, 500},
	aqi.Pollutant_CO_1H:     {0, 5, 10, 35, 60, 90, 120, 150},           // mg/m3
	aqi.Pollutant_CO_24H:    {0, 2, 4, 14, 24, 36, 48, 60},              // mg/m3
	aqi.Pollutant_SO2_24H:   {0, 50, 150, 475, 800, 1600, 2100, 2620},   // μg/m3
	aqi.Pollutant_SO2_1H:    {0, 150, 500, 650, 800},                    // μg/m3
	aqi.Pollutant_NO2_24H:   {0, 40, 80, 180, 280, 565, 750, 940},       // μg/m3
	aqi.Pollutant_NO2_1H:    {0, 100, 200, 700, 1200, 2340, 3090, 3840}, // μg/m3
	aqi.Pollutant_O3_1H:     {0, 160, 200, 300, 400, 800, 1000, 1200},   // μg/m3
	aqi.Pollutant_O3_8H:     {0, 100, 160, 215, 265, 800},               // μg/m3
	aqi.Pollutant_PM10_1H:   {0, 50, 150, 250, 350, 420, 500, 600},      // μg/m3
	aqi.Pollutant_PM10_24H:  {0, 50, 150, 250, 350, 420, 500, 600},      // μg/m3
	aqi.Pollutant_PM2_5_1H:  {0, 35, 75, 115, 150, 250, 350, 500},       // μg/m3
	aqi.Pollutant_PM2_5_24H: {0, 35, 75, 115, 150, 250, 350, 500},       // μg/m3
}

type Algo struct {
	FailedWhenNotSupported bool
}

func (a *Algo) Name() string {
	return "mep"
}

// Calc 计算策略是 HJ633-2012 中的实时报，其中采用 PM2.5 1H, PM10 1H 变量计算
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

		// 二氧化硫（SO2）1 小时平均浓度值高于 800 μg/m 3的，不再进行其空气质量分指数计算；
		// 二氧化硫（SO2） 空气质量分指数按 24 小时平均浓度计算的分指数报告。
		if pollutantVar.P == aqi.Pollutant_SO2_1H && pollutantVar.Value > 800 {
			continue
		}
		// 臭氧（O3）8 小时平均浓度值高于 800 μg/m 3的，不再进行其空气质量分指数计算；
		// 臭氧（O3）空气质量分指数按 1 小时平均浓度计算的分指数报告。
		if pollutantVar.P == aqi.Pollutant_O3_8H && pollutantVar.Value > 800 {
			continue
		}

		iaqiLo, iaqiHi, pLo, pHi, err := aqi.GetRanges(pollutantVar.Value, pollutantIndexrange, Tables[aqi.Pollutant_AQI])
		if err != nil {
			return 0, nil, err
		}

		aqi, err := aqi.CalcViaHiLo(pollutantVar.Value, iaqiLo, iaqiHi, pLo, pHi)
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
