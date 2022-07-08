package entities

import "encoding/json"

type CachedPlayer struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func NewCachedPlayerFromJSON(jsonString string) (*CachedPlayer, error) {
	var cachedPlayer CachedPlayer
	err := json.Unmarshal([]byte(jsonString), &cachedPlayer)
	if err != nil {
		return nil, err
	}
	return &cachedPlayer, nil
}
