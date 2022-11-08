package mep_test

import (
	"fmt"
	"reflect"
	"testing"

	goaqi "github.com/ringsaturn/go-aqi"
	"github.com/ringsaturn/go-aqi/mep"
)

var _ goaqi.StandardWithColor = &mep.Algo{}

func ExampleAlgo_Calc() {
	algo := &mep.Algo{}
	inputs := []*goaqi.Var{
		{
			P:     goaqi.Pollutant_PM2_5_1H,
			Value: 16,
		},
		{
			P:     goaqi.Pollutant_PM10_1H,
			Value: 88,
		},
		{
			P:     goaqi.Pollutant_CO_1H,
			Value: 0.2,
		},
		{
			P:     goaqi.Pollutant_SO2_1H,
			Value: 3,
		},
		{
			P:     goaqi.Pollutant_NO2_1H,
			Value: 11,
		},
		{
			P:     goaqi.Pollutant_O3_1H,
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
	algo := &mep.Algo{}
	rgba, err := algo.AQIToColor(33)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rgba)
	// Output: &{0 255 0 0}
}

func ExampleAlgo_AQIToDesc() {
	algo := &mep.Algo{}
	desc, err := algo.AQIToDesc(33)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", desc)
	// Output: 优
}

func TestAlgo_Calc(t *testing.T) {
	type fields struct {
		FailedWhenNotSupported bool
	}
	type args struct {
		pollutantVars []*goaqi.Var
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   []goaqi.Pollutant
		wantErr bool
	}{
		{
			name: "example",
			fields: fields{
				FailedWhenNotSupported: false,
			},
			args: args{
				pollutantVars: []*goaqi.Var{
					{
						P:     goaqi.Pollutant_PM2_5_1H,
						Value: 16,
					},
					{
						P:     goaqi.Pollutant_PM10_1H,
						Value: 88,
					},
					{
						P:     goaqi.Pollutant_CO_1H,
						Value: 0.2,
					},
					{
						P:     goaqi.Pollutant_SO2_1H,
						Value: 3,
					},
					{
						P:     goaqi.Pollutant_NO2_1H,
						Value: 11,
					},
					{
						P:     goaqi.Pollutant_O3_1H,
						Value: 75,
					},
				},
			},
			want:    69,
			want1:   []goaqi.Pollutant{goaqi.Pollutant_PM10_1H},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &mep.Algo{
				FailedWhenNotSupported: tt.fields.FailedWhenNotSupported,
			}
			got, got1, err := a.Calc(tt.args.pollutantVars...)
			if (err != nil) != tt.wantErr {
				t.Errorf("goaqi.Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("goaqi.Calc() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("goaqi.Calc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkAlgoCalc(b *testing.B) {
	algo := &mep.Algo{}
	inputs := []*goaqi.Var{
		{
			P:     goaqi.Pollutant_PM2_5_1H,
			Value: 16,
		},
		{
			P:     goaqi.Pollutant_PM10_1H,
			Value: 88,
		},
		{
			P:     goaqi.Pollutant_CO_1H,
			Value: 0.2,
		},
		{
			P:     goaqi.Pollutant_SO2_1H,
			Value: 3,
		},
		{
			P:     goaqi.Pollutant_NO2_1H,
			Value: 11,
		},
		{
			P:     goaqi.Pollutant_O3_1H,
			Value: 75,
		},
	}
	for i := 0; i < b.N; i++ {
		_, _, _ = algo.Calc(inputs...)
	}
}

var largeData [][]*goaqi.Var

func init() {
	for i := 0; i < 384; i++ {
		largeData = append(largeData, []*goaqi.Var{
			{
				P:     goaqi.Pollutant_PM2_5_1H,
				Value: 16,
			},
			{
				P:     goaqi.Pollutant_PM10_1H,
				Value: 88,
			},
			{
				P:     goaqi.Pollutant_CO_1H,
				Value: 0.2,
			},
			{
				P:     goaqi.Pollutant_SO2_1H,
				Value: 3,
			},
			{
				P:     goaqi.Pollutant_NO2_1H,
				Value: 11,
			},
			{
				P:     goaqi.Pollutant_O3_1H,
				Value: 75,
			},
		},
		)
	}
}

func BenchmarkAlgoCalcOnLargeData(b *testing.B) {
	algo := &mep.Algo{}
	for i := 0; i < b.N; i++ {
		func() {
			for _, input := range largeData {
				_, _, _ = algo.Calc(input...)
			}
		}()
	}
}
