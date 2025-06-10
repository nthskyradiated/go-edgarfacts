package data

import (
	"time"
)
type FactsLoaderHandlerInput struct {
	Cik   string `json:"cik"`
	Org   string `json:"org"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DataPointFacts struct {
	Start        string  `json:"start"`
	End          string  `json:"end"`
	Value        float64 `json:"val"`
	Account      string  `json:"account"`
	FiscalYear   int     `json:"fy"`
	FiscalPeriod string  `json:"fp"`
	Form         string  `json:"form"`
	Filed        string  `json:"filed"`
}

type MetricFacts struct {
	Label       string                      `json:"label"`
	Description string                      `json:"description"`
	Unit        map[string][]DataPointFacts `json:"units"`
}

type DataFacts struct {
	Cik    int                               `json:"cik"`
	Entity string                            `json:"entityName"`
	Facts  map[string]map[string]MetricFacts `json:"facts"`
}

func (f *DataFacts) Flatten() []OutFact {
	outfactsSlice := make([]OutFact, 0, 512)

	for framework, facts := range f.Facts {
		for labelMetric, metric := range facts {
			for labelUnit, unit := range metric.Unit {
				for _, dataPoint := range unit {
					formatTime := "2006-01-02"
					start, err := time.Parse(formatTime, dataPoint.Start)
					if err != nil {
						start = time.Time{}
					}
					end, err := time.Parse(formatTime, dataPoint.End)
					if err != nil {
						end = time.Time{}
					}
					filed, err := time.Parse(formatTime, dataPoint.Filed)
					if err != nil {
						filed = time.Time{}
					}
					factSingle := OutFact {
						Cik: f.Cik,
						Label: labelMetric,
						Unit: labelUnit,
						Start: start,
						End: end,
						Filed: filed,
						FiscalYear: dataPoint.FiscalYear,
						FiscalPeriod: dataPoint.FiscalPeriod,
						Acccount: dataPoint.Acccount,
						Form: dataPoint.Form,
						Framework: framework,
						Value: datapoint.Value
					}

					outfactsSlice = append(outfactsSlice, factSingle)
				}
			}
		}
	}

	return outfactsSlice
}

type ParseFactsHandlerInput struct {
	Cik string `json:"cik"`
}

type OutFact struct {
	Cik int `json:"cik"`
	Label string `json:"label"`
	Unit string `json:"unit"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Filed time.Time `json:"filed"`
	FiscalYear int `json:"fiscal_year"`
	FiscalPeriod string `json:"fiscal_period"`
	Account string `json:"account"`
	Form string `json:"form"`
	Framework string `json:"framework"`
	Value float64 `json:"value"`

}