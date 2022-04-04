package mep

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ringsaturn/aqi"
)

var _ aqi.StandardWithColor = &Algo{}

func ExampleAlgo_Calc() {
	algo := &Algo{}
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
	// Output: aqi=69 with primary pollutant as [PM10_1H]
}

func ExampleAlgo_AQIToColor() {
	algo := &Algo{}
	rgba, err := algo.AQIToColor(33)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rgba)
	// Output: &{0 255 0 0}
}

func TestAlgo_Calc(t *testing.T) {
	type fields struct {
		FailedWhenNotSupported bool
	}
	type args struct {
		pollutantVars []aqi.Var
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   []aqi.Pollutant
		wantErr bool
	}{
		{
			name: "example",
			fields: fields{
				FailedWhenNotSupported: false,
			},
			args: args{
				pollutantVars: []aqi.Var{
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
				},
			},
			want:    69,
			want1:   []aqi.Pollutant{aqi.Pollutant_PM10_1H},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algo{
				FailedWhenNotSupported: tt.fields.FailedWhenNotSupported,
			}
			got, got1, err := a.Calc(tt.args.pollutantVars...)
			if (err != nil) != tt.wantErr {
				t.Errorf("aqi.Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("aqi.Calc() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("aqi.Calc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkAlgoCalc(b *testing.B) {
	algo := &Algo{}
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
	for i := 0; i < b.N; i++ {
		_, _, _ = algo.Calc(inputs...)
	}
}
