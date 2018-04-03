package main

import (
	"fmt"
	"os"

	"github.com/pleum/pubgo"
)

func main() {
	apiKey := os.Getenv("PUBGAPIKEY")
	pubgAPI, err := pubgo.NewClient(apiKey)
	if err != nil {
		panic(err.Error())
	}

	_, err = pubgAPI.GetStatus()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("[PUBG PUBLIC API EXAMPLE]")

	players, err := pubgAPI.GetPlayers("Player Name", "pc-as", pubgo.FilterPlayerName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Hello,", players.Data[0].Attributes.Name)
	fmt.Println("Getting last match information.")
	match := players.Data[0].Relationships.Matches.Data[0]
	lastMatch, err := pubgAPI.GetMatch(match.ID, "pc-as")
	if err != nil {
		panic(err.Error())
	}

	var suicidePlayers []string

	winPlace := 0
	for _, p := range lastMatch.Included {
		if p.Type == "participant" {
			if p.Attributes.Stats.DeathType == "suicide" {
				suicidePlayers = append(suicidePlayers, p.Attributes.Stats.Name)
			}
			if p.Attributes.Stats.Name == players.Data[0].Attributes.Name {
				winPlace = p.Attributes.Stats.WinPlace
			}
		}
	}

	fmt.Printf("You got %v place.\n", winPlace)
	fmt.Printf("Had %v players suicide in last match: %v\n", len(suicidePlayers), suicidePlayers)
}
