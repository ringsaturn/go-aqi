# Compute AQI in Go [![Go Reference](https://pkg.go.dev/badge/github.com/ringsaturn/go-aqi.svg)](https://pkg.go.dev/github.com/ringsaturn/go-aqi)

```bash
go get github.com/ringsaturn/go-aqi
```

For usage see [examples](_example/).

NOTE: Currently the algo impl is based on the different standard files and
different AQI Standard use different units. Please ensure the input value has
been converted to the algo expect unit.

|                | CO               | PM 2.5           | PM 10            | SO<sub>2</sub>   | NO<sub>2</sub>   | Ozone/O<sub>3</sub> |
| -------------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- | ------------------- |
| MEP(China)[^1] | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup> | μg/m<sup>3</sup>    |
| EPA(USA)[^2]   | ppm              | μg/m<sup>3</sup> | μg/m<sup>3</sup> | ppb              | ppb              | ppm                 |

[^1]: [环境空气质量指数（AQI）技术规定](https://www.mee.gov.cn/ywgz/fgbz/bz/bzwb/jcffbz/201203/W020120410332725219541.pdf)

[^2]: [Guideline for Reporting of Daily Air Quality: Air Quality Index](https://www.airnow.gov/sites/default/files/2020-05/aqi-technical-assistance-document-sept2018.pdf)
