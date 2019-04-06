package tests

import (
	"miniature-parakeet/parse"
	"testing"
)

func TestGetAverage(t *testing.T) {
	parser := parse.AverageParser{}
	parser.GetAverage("", "апр", 1)
}
