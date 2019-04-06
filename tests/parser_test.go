package tests

import (
	"github.com/smartystreets/assertions"
	"miniature-parakeet/parse"
	"testing"
)

func TestGetAddress(t *testing.T) {
	parser := parse.Parser{}
	result := parser.GetAddressOfAction("Frontier")
	assertions.ShouldEqual(result, "https://quote.rbc.ru/ticker/177752")
}
