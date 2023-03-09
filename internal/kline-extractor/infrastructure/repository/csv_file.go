package repository

import (
	"bybit-kline-extractor/internal/kline-extractor/domain/model"
	"bybit-kline-extractor/internal/kline-extractor/infrastructure/repository/dto"
	"bybit-kline-extractor/pkg/csv_file"
)

type CsvFile struct {
	csvPkg *csv_file.CsvFile
}

func NewCsvFile(csvPkg *csv_file.CsvFile) *CsvFile {
	return &CsvFile{csvPkg: csvPkg}
}

func (c *CsvFile) AddChartsModelToCsvs(charts ...model.Chart) error {
	chartSeries := make([][][]string, len(charts))
	for i, chart := range charts {
		chartDto := dto.ChartModelToDto(chart)
		chartString := dto.ChartSliceToStringSlices(chartDto)
		chartSeries[i] = chartString
	}

	return c.csvPkg.WriteMultipleSeriesInMultipleCsvs(chartSeries...)
}
