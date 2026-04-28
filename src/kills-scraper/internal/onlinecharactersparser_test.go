package internal_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vitor-schipani/tibia-path-etl/src/kills-scraper/internal"
)

type ParseOnlineCharsSuite struct {
	suite.Suite
}

func TestParseOnlineCharsSuite(t *testing.T) {
	suite.Run(t, new(ParseOnlineCharsSuite))
}

func (s *ParseOnlineCharsSuite) TestCorrectBehavior() {
	resp := []byte(`{
		"world": {
			"online_players": [
				{"name":"PlayerOne"},
				{"name":"PlayerTwo"}
			]
		}
	}`)

	expected := []string{"PlayerOne", "PlayerTwo"}
	actual, err := internal.ParseOnlineCharacters(resp)

	s.Require().NoError(err)
	s.Equal(expected, actual)
}
