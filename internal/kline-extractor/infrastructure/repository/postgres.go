package repository

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"bybit-kline-extractor/internal/kline-extractor/infrastructure/repository/dto"
	"bybit-kline-extractor/pkg/postgres"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type ChartDB struct {
	pg      *postgres.Postgres
	context context.Context
}

func NewPostgres(pg *postgres.Postgres, context context.Context) ChartDB {
	return ChartDB{
		pg:      pg,
		context: context,
	}
}

func (c ChartDB) InsertKlineToChart(chart model.Chart) error {
	sql, args, err := c.pg.SqlBuilder.Insert("charts").Rows(dto.ChartModelToDto(chart)).ToSQL()

	if err != nil {
		return fmt.Errorf("postgres - AddChart - tosql: %w", err)
	}

	_, err = c.pg.PgPool.Exec(c.context, sql, args...)
	if err != nil {
		return fmt.Errorf("postgres - AddChart - Exec: %w", err)
	}

	return nil
}

func (c ChartDB) SelectChart(exchange, pair, timeFrame string) (model.Chart, error) {
	sql, args, err := c.pg.SqlBuilder.Select("*").
		From("charts").
		Where(
			goqu.Ex{
				"exchange":   exchange,
				"pair":       pair,
				"time_frame": timeFrame,
			}).
		Order(goqu.I("time_frame").Asc()).
		ToSQL()

	if err != nil {
		return model.Chart{
				Klines: []model.Kline{},
				ChartDescription: model.ChartDescription{
					Exchange:           exchange,
					TimeFrame:          timeFrame,
					CryptoCurrencyPair: pair,
				},
			},
			fmt.Errorf("postgres - GetKlines - ToSql: %w", err)
	}

	rows, err := c.pg.PgPool.Query(c.context, sql, args...)
	if err != nil {
		return model.Chart{}, fmt.Errorf("postgres - GetChart - Query:%w", err)
	} else if rows.Err() != nil {
		return model.Chart{}, fmt.Errorf("postgres - GetChart - Query:%w", rows.Err())
	}

	charts := make([]dto.Chart, 0)

	for rows.Next() {
		chart := dto.Chart{
			Exchange:  exchange,
			Pair:      pair,
			TimeFrame: timeFrame,
			Kline:     dto.Kline{},
		}
		rows.Scan(&chart.Exchange,
			&chart.Pair,
			&chart.TimeFrame,
			&chart.OpenTime,
			&chart.Open,
			&chart.High,
			&chart.Low,
			&chart.Close,
			&chart.Volume,
		)

		charts = append(charts, chart)
	}

	chartModel := dto.ChartsDtoToChartModel(charts)

	return chartModel, nil
}
