package bybit

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"time"
)

type BaseResult struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	ExtInfo          string      `json:"ext_info"`
	Result           interface{} `json:"result"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}

type ResultStringArrayResponse struct {
	BaseResult
	Result []string `json:"result"`
}

type Item struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size"`
}

type OrderBook struct {
	Asks []Item    `json:"asks"`
	Bids []Item    `json:"bids"`
	Time time.Time `json:"time"`
}

type RawItem struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Size   float64 `json:"size"`
	Side   string  `json:"side"` // Buy/Sell
}

type GetOrderBookResult struct {
	BaseResult
	Result []RawItem `json:"result"`
}

type OHLC struct {
	Symbol   string  `json:"symbol"`
	Interval string  `json:"interval"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open,string"`
	High     float64 `json:"high,string"`
	Low      float64 `json:"low,string"`
	Close    float64 `json:"close,string"`
	Volume   float64 `json:"volume,string"`
	Turnover float64 `json:"turnover,string"`
}

type GetKlineResult struct {
	BaseResult
	Result []OHLC `json:"result"`
}

type OHLCLinear struct {
	Symbol   string  `json:"symbol"`
	Period   string  `json:"period"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"`
}

func (o OHLCLinear) ToKline() model.Kline {
	kline := model.Kline{
		OpenTime: 0,
		Open:     0,
		High:     0,
		Low:      0,
		Close:    0,
		Volume:   0,
	}
	kline.Open = o.Open
	kline.High = o.High
	kline.Low = o.Low
	kline.Close = o.Close
	kline.Volume = o.Volume
	kline.OpenTime = o.OpenTime
	return kline
}

type OHLCLinearSlice []OHLCLinear

func (o OHLCLinearSlice) ToKline() []model.Kline {
	klines := make([]model.Kline, 0)
	for _, ohlc := range o {
		klines = append(klines, ohlc.ToKline())
	}
	return klines
}

type GetLinearKlineResult struct {
	BaseResult
	Result []OHLCLinear `json:"result"`
}
