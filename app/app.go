package app

import (
	"bybit-kline-extractor/config"
	kline_extractor "bybit-kline-extractor/internal/kline-extractor"
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"bybit-kline-extractor/internal/kline-extractor/infrastructure/repository"
	"bybit-kline-extractor/pkg/csv_file"
	"fmt"
	"log"
	"time"
)

// TODO: change from and to
func Run(cfg *config.Config) {
	startTime := time.Now().UTC().Unix()
	klineExtractor, err := kline_extractor.New(cfg)
	if err != nil {
		log.Fatalln(fmt.Errorf("app - New - kline_extractor.New: %w", err))
	}
	timeFrames := make([]model.TimeFrame, 0)
	timeFrames = append(timeFrames,
		model.ONE_MIN,
		model.ONE_WEEK,
		model.ONE_MONTH,
	)
	exchange := "bybit"
	pair := "ETHUSDT"
	//	from := int64(1604192461)
	//	to := int64(1675248799)

	chartDescriptions := make([]model.ChartDescription, 0)

	for _, timeFrame := range timeFrames {
		chartDescription := model.ChartDescription{
			Exchange:           exchange,
			TimeFrame:          string(timeFrame),
			CryptoCurrencyPair: pair,
		}

		chartDescriptions = append(chartDescriptions, chartDescription)
	}

	charts := klineExtractor.GetChartService().GetMultipleChartsFromRepository(chartDescriptions...)
	csvPkg := csv_file.New()
	csvRepo := repository.NewCsvFile(csvPkg)
	err = csvRepo.AddChartsModelToCsvs(charts...)
	if err != nil {
		log.Fatalln(err)
	}

	finalTime := time.Now().UTC().Unix()
	duration := finalTime - startTime
	fmt.Println("duration in minutes", duration/60)
}
