package main

import (
	"bybit-kline-extractor/app"
	"bybit-kline-extractor/config"
)

func main() {
	cfg := config.New()
	app.Run(cfg)
}
