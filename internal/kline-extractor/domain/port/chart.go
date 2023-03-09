package port

import "bybit-kline-extractor/internal/kline-extractor/domain/model"

type ChartRepo interface {
	InsertKlineToChart(chart model.Chart) error
	SelectChart(exchange, pair, timeFrame string) (model.Chart, error)
}

type ChartApi interface {
	GetKlines(pair string, timeFrame model.TimeFrame, from int64) ([]model.Kline, error)
}
