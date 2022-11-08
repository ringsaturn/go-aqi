package goaqi

import (
	"fmt"
	"image/color"
)

type Var struct {
	P     Pollutant
	Value float32 // only co need float, other need int only
}

func (v *Var) MiuGPerM3ToPPM() *Var {
	v.Value = MiuGPerM3ToMgPerM3(v.Value)
	v.Value = MgPerM3ToPPM(v.P, v.Value)
	return v
}

func (v *Var) MiuGPerM3ToPPB() *Var {
	v.Value = MiuGPerM3ToMgPerM3(v.Value)
	v.Value = MgPerM3ToPPM(v.P, v.Value)
	return v
}

func (v *Var) MgPerM3ToPPM() *Var {
	v.Value = MgPerM3ToPPM(v.P, v.Value)
	return v
}

func (v *Var) PPMToMgPerM3() *Var {
	v.Value = PPMToMgPerM3(v.P, v.Value)
	return v
}

func (v *Var) MgPerM3ToPPB() *Var {
	v.MgPerM3ToPPM()
	v.Value = PPMToPPB(v.Value)
	return v
}

func (v *Var) PPMToPPB() *Var {
	v.Value = PPMToPPB(v.Value)
	return v
}

func (v *Var) PPBToPPM() *Var {
	v.Value = PPBToPPM(v.Value)
	return v
}

func (v *Var) PPBToMgPerM3() *Var {
	v = v.PPBToPPM()
	v = v.PPMToMgPerM3()
	return v
}

type Standard interface {
	Name() string
	// Calc 返回多个污染物表示并列首要污染物
	Calc(pollutantVars ...*Var) (int, []Pollutant, error)
}

type StandardWithColor interface {
	Standard
	AQIToColor(aqi int) (*color.RGBA, error)
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

// https://teesing.com/en/library/tools/ppm-mg3-converter
var MolecularWeight = map[Pollutant]float32{
	Pollutant_CO_1H:   28.01,
	Pollutant_CO_8H:   28.01,
	Pollutant_CO_24H:  28.01,
	Pollutant_NO2_1H:  46.01,
	Pollutant_NO2_24H: 46.01,
	Pollutant_O3_1H:   48,
	Pollutant_O3_8H:   48,
	Pollutant_SO2_1H:  64.06,
	Pollutant_SO2_24H: 64.06,
}

func PPMToPPB(value float32) float32 {
	return 1000 * value
}

func PPBToPPM(value float32) float32 {
	return value / 1000
}

func PPMToMgPerM3(p Pollutant, value float32) float32 {
	return 0.0409 * value * MolecularWeight[p]
}

func MgPerM3ToPPM(p Pollutant, value float32) float32 {
	return 24.45 * value / MolecularWeight[p]
}

func MiuGPerM3ToMgPerM3(v float32) float32 {
	return v / 1000
}

func MgGPerM3ToMiuGPerM3(v float32) float32 {
	return v * 1000
}
