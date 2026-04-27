package internal

import "encoding/json"

type worldsResponse struct {
	Worlds struct {
		RegularWorlds []struct {
			Name string `json:"name"`
		} `json:"regular_worlds"`
	} `json:"worlds"`
}

func ParseWorldsResponse(data []byte) ([]string, error) {
	var r worldsResponse
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(r.Worlds.RegularWorlds))
	for _, w := range r.Worlds.RegularWorlds {
		names = append(names, w.Name)
	}

	return names, nil
}
