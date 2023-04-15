# Compute AQI in Go [![ci](https://github.com/ringsaturn/go-aqi/actions/workflows/ci.yml/badge.svg)](https://github.com/ringsaturn/go-aqi/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/ringsaturn/go-aqi.svg)](https://pkg.go.dev/github.com/ringsaturn/go-aqi)

```bash
go get github.com/ringsaturn/go-aqi
```

For usage see [examples](_example/).

NOTE: Currently the algo impl is based on the different standard files and
different AQI Standard use different units. Please ensure the input value has
been converted to the algo expect unit.

|            | CO               | PM 2.5           | PM 10            | SO<sub>2</sub>   | NO<sub>2</sub>   | Ozone/O<sub>3</sub> |
| ---------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- | ------------------- |
| MEP(China) | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup>    |
| EPA(USA)   | ppm              | μg/m<sup>3</sup> | μg/m<sup>3</sup> | ppb              | ppb              | ppm                 |
