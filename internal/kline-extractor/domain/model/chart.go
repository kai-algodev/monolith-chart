package model

type Chart struct {
	Klines []Kline
	ChartDescription
}

func NewChart(kline []Kline, exchage, timeFrame, cryptoCurrencyPair string) Chart {
	return Chart{
		Klines: kline,
		ChartDescription: ChartDescription{
			Exchange:           exchage,
			TimeFrame:          timeFrame,
			CryptoCurrencyPair: cryptoCurrencyPair,
		},
	}
}

type ChartDescription struct {
	Exchange           string
	TimeFrame          string
	CryptoCurrencyPair string
}
