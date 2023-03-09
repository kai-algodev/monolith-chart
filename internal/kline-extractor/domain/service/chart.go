package service

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"bybit-kline-extractor/internal/kline-extractor/domain/port"
	"fmt"
	"log"
	"sync"
)

type ChartService interface {
	AddKlineToChart(chart model.Chart) error
	StoreCharts(timeFrames []model.TimeFrame, exchange, pair string, from, to int64)
	GetMultipleChartsFromRepository(chartDescriptions ...model.ChartDescription) []model.Chart
}

type Chart struct {
	repo   port.ChartRepo
	webApi port.ChartApi
}

func NewChart(repo port.ChartRepo, webApi port.ChartApi) Chart {
	return Chart{
		repo:   repo,
		webApi: webApi,
	}
}

func (c Chart) AddKlineToChart(chart model.Chart) error {
	return c.repo.InsertKlineToChart(chart)
}

func (c Chart) StoreCharts(timeFrames []model.TimeFrame, exchange, pair string, from, to int64) {
	var wg = new(sync.WaitGroup)
	wg.Add(len(timeFrames))

	for _, t := range timeFrames {
		go c.addChartToRepository(t, exchange, pair, wg, from, to)
	}
	wg.Wait()

}

func (c Chart) addChartToRepository(timeFrame model.TimeFrame, exchange, pair string, wg *sync.WaitGroup, from, to int64) {
	timeFrameToSecond, err := timeFrame.ToSecond()
	if err != nil {
		log.Fatalln(err)
	}
	lthr := to - (200 * timeFrameToSecond)
	remainKlines := (to - from) / timeFrameToSecond
	fmt.Println("remaining klines in timeframe:", timeFrame, "->", remainKlines, "klines")
	klines, err := c.webApi.GetKlines(pair, timeFrame, from)
	if err != nil {
		log.Fatalln(err)
	}

	var startTime int64 = 0

	if len(klines) != 0 {
		startTime = klines[0].OpenTime
		chart := model.NewChart(klines, exchange, string(timeFrame), pair)
		err = c.repo.InsertKlineToChart(chart)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if from >= lthr {
		wg.Done()
		return
	}
	newFrom := startTime + (200 * timeFrameToSecond)
	c.addChartToRepository(timeFrame, exchange, pair, wg, newFrom, to)
}

func (c Chart) GetMultipleChartsFromRepository(chartDescriptions ...model.ChartDescription) []model.Chart {

	chartsCount := len(chartDescriptions)
	charts := make([]model.Chart, 0)
	var wg = new(sync.WaitGroup)

	wg.Add(chartsCount)
	for _, chartDescription := range chartDescriptions {
		cdd := chartDescription
		go func(cd model.ChartDescription) {
			chart, err := c.repo.SelectChart(cd.Exchange, cd.CryptoCurrencyPair, cd.TimeFrame)
			charts = append(charts, chart)
			if err != nil {
				log.Println("could not get chart from repo, chart:", cd, "Error:", err)
				return
			}
			wg.Done()
		}(cdd)
	}
	wg.Wait()

	return charts
}
