package dto

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"fmt"
)

type Chart struct {
	Exchange  string `db:"exchange"`
	Pair      string `db:"pair"`
	TimeFrame string `db:"time_frame"`
	Kline
}

func ChartModelToDto(chart model.Chart) []Chart {
	chartDtoList := make([]Chart, 0)
	for _, c := range chart.Klines {
		var chartDto Chart
		chartDto.Exchange = chart.Exchange
		chartDto.Pair = chart.CryptoCurrencyPair
		chartDto.TimeFrame = chart.TimeFrame
		chartDto.OpenTime = c.OpenTime
		chartDto.Open = c.Open
		chartDto.High = c.High
		chartDto.Low = c.Low
		chartDto.Close = c.Close
		chartDto.Volume = c.Volume
		chartDtoList = append(chartDtoList, chartDto)
	}
	return chartDtoList
}

func ChartsDtoToChartModel(charts []Chart) model.Chart {
	chart := model.Chart{
		Klines: []model.Kline{},
		ChartDescription: model.ChartDescription{
			Exchange:           charts[0].Exchange,
			TimeFrame:          charts[0].TimeFrame,
			CryptoCurrencyPair: charts[0].Pair,
		},
	}

	for _, c := range charts {
		kline := model.NewKline(c.OpenTime, c.Open, c.High, c.Low, c.Close, c.Volume)
		chart.Klines = append(chart.Klines, kline)
	}

	return chart
}

func ChartSliceToStringSlices(charts []Chart) [][]string {
	stringSlices := make([][]string, 0)
	for index := range stringSlices {
		stringSlices[index] = make([]string, 0)
	}

	for _, chart := range charts {
		stringSlice := make([]string, 0)

		stringSlice = append(stringSlice,
			chart.Exchange,
			chart.Pair,
			chart.TimeFrame,
			fmt.Sprintf("%v", chart.OpenTime),
			fmt.Sprintf("%v", chart.Open),
			fmt.Sprintf("%v", chart.High),
			fmt.Sprintf("%v", chart.Low),
			fmt.Sprintf("%v", chart.Close),
			fmt.Sprintf("%v", chart.Volume),
		)

		stringSlices = append(stringSlices, stringSlice)
	}
	return stringSlices
}
