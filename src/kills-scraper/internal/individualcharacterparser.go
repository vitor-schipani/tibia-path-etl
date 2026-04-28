package internal

import (
	"encoding/json"
)

type RawCharacterResponse struct {
	CharacterHeader CharacterWrapper `json:"character"`
}

type CharacterWrapper struct {
	Deaths    []Death   `json:"deaths"`
	Character Character `json:"character"`
}

type Character struct {
	Name              string `json:"name"`
	Sex               string `json:"sex"`
	Title             string `json:"title"`
	Vocation          string `json:"vocation"`
	World             string `json:"world"`
	Residence         string `json:"residence"`
	AccountStatus     string `json:"account_status"`
	UnlockedTitles    int    `json:"unlocked_titles"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievement_points"`
}

type Death struct {
	Killers []Killer `json:"killers"`
	Level   int      `json:"level"`
}

type Killer struct {
	Name     string `json:"name"`
	IsPlayer bool   `json:"player"`
}

func IndividualCharacterParser(data []byte) (CharacterWrapper, error) {
	var r RawCharacterResponse
	err := json.Unmarshal(data, &r)
	if err != nil {
		return CharacterWrapper{}, err
	}
	return r.CharacterHeader, nil
}
