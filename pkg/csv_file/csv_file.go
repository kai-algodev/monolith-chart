package csv_file

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

type CsvFile struct {
}

func New() *CsvFile {
	return &CsvFile{}
}

func (c *CsvFile) WriteMultipleSeriesInMultipleCsvs(series ...[][]string) error {
	var wg = new(sync.WaitGroup)
	var internalError error

	wg.Add(len(series))
	for _, serie := range series {
		go func(serie [][]string) {
			exchange := serie[0][0]
			pair := serie[0][1]
			timeFrame := serie[0][2]
			firstOpenTime := serie[0][3]
			lastOpentime := serie[len(serie)-1][4]
			path := fmt.Sprintf("data/%v-%v-%v-%v-%v.csv", exchange, pair, timeFrame, firstOpenTime, lastOpentime)
			file, err := os.Create(path)
			if err != nil {
				log.Println(fmt.Errorf("csv_file - WriteMultipleSeriesInMultipleCsvs - Create: %w", err))
				internalError = err
			}

			writer := csv.NewWriter(file)
			writer.WriteAll(serie)
			if err = writer.Error(); err != nil {
				log.Println(fmt.Errorf("csv_file - WriteMultipleSeriesInMultipleCsvs - WriteAll: %w", err))
				internalError = err
			}

			wg.Done()
		}(serie)
	}

	wg.Wait()

	return internalError
}
