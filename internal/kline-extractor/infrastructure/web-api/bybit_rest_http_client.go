package web_api

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"bybit-kline-extractor/pkg/bybit"
)

type BybitRestHttpClient struct {
	bybit *bybit.Bybit
}

func NewBybitWebApi(bybit *bybit.Bybit) BybitRestHttpClient {
	return BybitRestHttpClient{
		bybit: bybit,
	}
}

func (s BybitRestHttpClient) GetKlines(pair string, timeFrame model.TimeFrame, from int64) ([]model.Kline, error) {
	_, _, ohlcs, err := s.bybit.LinearGetKLine(pair, string(timeFrame), from, 200)
	if err != nil {
		return nil, err
	}
	klines := (bybit.OHLCLinearSlice)(ohlcs).ToKline()
	return klines, nil
}
