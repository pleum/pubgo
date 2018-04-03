package pubgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type Client struct {
	apiKey string
	gzip   bool
}

func (c Client) GetStatus() (*Status, error) {
	urlAPI := pubgAPIBase + statusEndpoint

	body, err := createRequest(c, urlAPI, nil, false)
	if err != nil {
		return nil, err
	}

	status := &Status{}
	err = json.NewDecoder(body).Decode(&status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c Client) GetPlayer(playerID string, shard shard) (*Player, error) {
	urlAPI := fmt.Sprintf(pubgAPIShard+"/%s", string(shard), playersEndpoint, playerID)
	body, err := createRequest(c, urlAPI, nil, true)
	if err != nil {
		return nil, err
	}

	player := &Player{}
	err = json.NewDecoder(body).Decode(&player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (c Client) GetPlayers(player string, shard shard, filter filterPlayer) (*Players, error) {
	urlAPI := fmt.Sprintf(pubgAPIShard, string(shard), playersEndpoint)
	query := url.Values{}

	if filter == FilterPlayerID {
		query.Set("filter[playerIds]", player)
	} else if filter == FilterPlayerName {
		query.Set("filter[playerNames]", player)
	} else {
		return nil, errors.New("Invalid filter type")
	}

	body, err := createRequest(c, urlAPI, query, true)
	if err != nil {
		return nil, err
	}

	players := &Players{}
	err = json.NewDecoder(body).Decode(&players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (c Client) GetMatch(matchID string, shard shard) (*Match, error) {
	urlAPI := fmt.Sprintf(pubgAPIShard+"/%s", string(shard), matchesEndpoint, matchID)
	body, err := createRequest(c, urlAPI, nil, true)
	if err != nil {
		return nil, err
	}

	match := &Match{}
	err = json.NewDecoder(body).Decode(&match)
	if err != nil {
		return nil, err
	}

	return match, nil
}
