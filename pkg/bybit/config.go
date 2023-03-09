package bybit

import "time"

type BybitOption func(*Bybit)

func TimeOut(timeOut time.Duration) BybitOption {
	return (BybitOption)(func(bybit *Bybit) {
		bybit.client.Timeout = timeOut
	})
}

func BaseUrl(url string) BybitOption {
	return (BybitOption)(func(bybit *Bybit) {
		bybit.baseURL = url
	})
}
