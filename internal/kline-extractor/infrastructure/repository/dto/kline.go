package dto

type Kline struct {
	OpenTime int64   `db:"open_time"`
	Open     float64 `db:"open"`
	High     float64 `db:"high"`
	Low      float64 `db:"low"`
	Close    float64 `db:"close"`
	Volume   float64 `db:"volume"`
}
