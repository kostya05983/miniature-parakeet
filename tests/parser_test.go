package tests

import (
	"github.com/smartystreets/assertions"
	"miniature-parakeet/parse"
	"testing"
)

func TestGetAddress(t *testing.T) {
	parser := parse.ParserLink{}
	result := parser.GetAddressOfAction("Frontier")
	assertions.ShouldEqual(result, "")
}
