package model

type Kline struct {
	OpenTime int64
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
}

func NewKline(openTime int64, open, high, low, close, volume float64) Kline {
	return Kline{
		OpenTime: openTime,
		Open:     open,
		High:     high,
		Low:      low,
		Close:    close,
		Volume:   volume,
	}
}
