package internal_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vitor-schipani/tibia-path-etl/src/kills-scraper/internal"
)

type ParseWorldNamesSuite struct {
	suite.Suite
}

func TestParseWorldNamesSuite(t *testing.T) {
	suite.Run(t, new(ParseWorldNamesSuite))
}

func (s *ParseWorldNamesSuite) TestCorrectBehavior() {
	resp := []byte(`{
		"worlds": {
			"regular_worlds": [
				{"name": "Aethera"},
				{"name": "Antica"}
			]
		}
	}`)

	expected := []string{"Aethera", "Antica"}
	actual, err := internal.ParseWorldsResponse(resp)

	s.Require().NoError(err)
	s.Equal(expected, actual)
}
