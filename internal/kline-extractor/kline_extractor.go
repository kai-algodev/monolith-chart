package kline_extractor

import (
	"bybit-kline-extractor/config"
	"bybit-kline-extractor/internal/kline-extractor/domain/service"
	"bybit-kline-extractor/internal/kline-extractor/infrastructure/repository"
	web_api "bybit-kline-extractor/internal/kline-extractor/infrastructure/web-api"
	"bybit-kline-extractor/pkg/bybit"
	"bybit-kline-extractor/pkg/postgres"
	"context"
	"fmt"
)

type KlineExtractor struct {
	chartService service.ChartService
}

func New(config *config.Config) (KlineExtractor, error) {
	bybit := bybit.New(bybit.BaseUrl(bybit.MAIN_NET))
	chartWebApi := web_api.NewBybitWebApi(bybit)
	fmt.Println("mod webapi", chartWebApi)
	postgresPkg, err := postgres.New(config.PG_URL)
	if err != nil {
		return KlineExtractor{}, fmt.Errorf("kline_extractor - New - postgres.New:%w", err)
	}

	postgresRepo := repository.NewPostgres(postgresPkg, context.Background())
	chartService := service.NewChart(postgresRepo, chartWebApi)

	return KlineExtractor{
		chartService: chartService,
	}, nil
}

func (k KlineExtractor) GetChartService() service.ChartService {
	return k.chartService
}
