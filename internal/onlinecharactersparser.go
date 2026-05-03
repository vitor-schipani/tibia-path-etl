package internal

import "encoding/json"

type OnlineCharacters struct {
	World struct {
		OnlinePlayers []struct {
			Name string `json:"name"`
		} `json:"online_players"`
	} `json:"world"`
}

func ParseOnlineCharacters(data []byte) ([]string, error) {
	var r OnlineCharacters
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	chars := make([]string, 0, len(r.World.OnlinePlayers))
	for _, w := range r.World.OnlinePlayers {
		chars = append(chars, w.Name)
	}

	return chars, nil
}
