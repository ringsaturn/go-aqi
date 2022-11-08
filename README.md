# Compute AQI in Go [![ci](https://github.com/ringsaturn/go-aqi/actions/workflows/ci.yml/badge.svg)](https://github.com/ringsaturn/go-aqi/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/ringsaturn/go-aqi.svg)](https://pkg.go.dev/github.com/ringsaturn/go-aqi)

```bash
go get github.com/ringsaturn/go-aqi
```

For usage see [examples](_example/).

NOTE: Currently the algo impl is based on the different standard files and
different AQI Standard use different units.
Please ensure the input value has been converted to the algo expect unit.

| Pollutant           | MEP(China)       | EPA(USA)         |
| ------------------- | ---------------- | ---------------- |
| CO                  | mg/m<sup>3</sup> | ppm              |
| PM 2.5              | μg/m<sup>3</sup> | μg/m<sup>3</sup> |
| PM 10               | μg/m<sup>3</sup> | μg/m<sup>3</sup> |
| SO<sub>2</sub>      | μg/m<sup>3</sup> | ppb              |
| NO<sub>2</sub>      | μg/m<sup>3</sup> | ppb              |
| Ozone/O<sub>3</sub> | μg/m<sup>3</sup> | ppm              |
