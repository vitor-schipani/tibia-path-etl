package internal_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vitor-schipani/tibia-path-etl/src/kills-scraper/internal"
)

type ParseCharacterSuite struct {
	suite.Suite
}

func TestParseCharacterSuite(t *testing.T) {
	suite.Run(t, new(ParseCharacterSuite))
}

func (s *ParseCharacterSuite) TestCorrectBehavior() {
	resp := []byte(`{
		"character":{
			"character":{
				"name":"CharOne",
				"sex":"Male",
				"title":"None",
				"unlocked_titles":19,
				"vocation":"Royal Paladin",
				"level":1111,
				"achievement_points":69,
				"world":"Example",
				"residence":"Example",
				"account_status":"Example"
			},
			"deaths":[{
				"level":600,
				"killers":[{
					"name":"lala",
					"player":false
				}]
			}]
		}
	}`)

	parsed, err := internal.IndividualCharacterParser(resp)

	s.Require().NoError(err)
	s.Equal("CharOne", parsed.Character.Name)
	s.Equal("Royal Paladin", parsed.Character.Vocation)
}
