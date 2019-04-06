package commands

import (
	"miniature-parakeet/models"
	"miniature-parakeet/parse"
)

type Average struct {
}

func (average Average) Calc(model models.AverageModel) float64 {
	parserLink := parse.ParserLink{}
	link := parserLink.GetAddressOfAction(model.CompanyName)
	parserAverage := parse.AverageParser{}
	return parserAverage.GetAverage(link, model.Month, model.Day)
}
