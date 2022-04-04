package aqi

import (
	"fmt"
	"image/color"
)

type Var struct {
	P     Pollutant
	Value float32 // only co need float, other need int only
}

type Standard interface {
	Name() string
	// Calc 返回多个污染物表示并列首要污染物
	Calc(...Var) ([]*Var, error)
}

type StandardWithColor interface {
	Standard
	AQIToColor(aqi int) (color.Alpha, error)
}

func GetRanges(value float32, pIndexRange []float32, aqiIndexRange []float32) (iaqiLo, iaqiHi, pLo, pHi float32, err error) {
	for i, v := range pIndexRange {
		if i == len(pIndexRange)-1 {
			return aqiIndexRange[i-1], aqiIndexRange[i], pIndexRange[i-1], v, nil
		}
		if pIndexRange[i] < value && value <= pIndexRange[i+1] {
			return aqiIndexRange[i], aqiIndexRange[i+1], v, pIndexRange[i+1], nil
		}
	}
	return 0, 0, 0, 0, fmt.Errorf("bad range value=%+v for pIndexRange=%+v", value, pIndexRange)
}

func CalcViaHiLo(value, iaqiLo, iaqiHi, pLo, pHi float32) (int, error) {
	return int((iaqiHi-iaqiLo)/(pHi-pLo)*(value-pLo) + iaqiLo), nil
}
