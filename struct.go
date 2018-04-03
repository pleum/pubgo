package pubgo

import "time"

type Status struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			ReleasedAt time.Time `json:"releasedAt"`
			Version    string    `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
}

type playerData struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt    time.Time   `json:"createdAt"`
		Name         string      `json:"name"`
		PatchVersion string      `json:"patchVersion"`
		ShardID      string      `json:"shardId"`
		Stats        interface{} `json:"stats"`
		TitleID      string      `json:"titleId"`
		UpdatedAt    time.Time   `json:"updatedAt"`
	} `json:"attributes"`
	Relationships struct {
		Assets struct {
			Data []interface{} `json:"data"`
		} `json:"assets"`
		Matches struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"matches"`
	} `json:"relationships"`
	Links struct {
		Schema string `json:"schema"`
		Self   string `json:"self"`
	} `json:"links"`
}

type Player struct {
	Data  playerData `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}

type Players struct {
	Data  []playerData `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}

type matchData struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt    time.Time   `json:"createdAt"`
		Duration     int         `json:"duration"`
		GameMode     string      `json:"gameMode"`
		PatchVersion string      `json:"patchVersion"`
		ShardID      string      `json:"shardId"`
		Stats        interface{} `json:"stats"`
		Tags         interface{} `json:"tags"`
		TitleID      string      `json:"titleId"`
	} `json:"attributes"`
	Relationships struct {
		Assets struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"assets"`
		Rosters struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"rosters"`
		Rounds struct {
			Data []interface{} `json:"data"`
		} `json:"rounds"`
		Spectators struct {
			Data []interface{} `json:"data"`
		} `json:"spectators"`
	} `json:"relationships"`
	Links struct {
		Schema string `json:"schema"`
		Self   string `json:"self"`
	} `json:"links"`
}

type Match struct {
	Data     matchData `json:"data"`
	Included []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Actor   string `json:"actor"`
			ShardID string `json:"shardId"`
			Stats   struct {
				DBNOs           int     `json:"DBNOs"`
				Assists         int     `json:"assists"`
				Boosts          int     `json:"boosts"`
				DamageDealt     float32 `json:"damageDealt"`
				DeathType       string  `json:"deathType"`
				HeadshotKills   int     `json:"headshotKills"`
				Heals           int     `json:"heals"`
				KillPlace       int     `json:"killPlace"`
				KillPoints      int     `json:"killPoints"`
				KillPointsDelta float32 `json:"killPointsDelta"`
				KillStreaks     int     `json:"killStreaks"`
				Kills           int     `json:"kills"`
				LastKillPoints  int     `json:"lastKillPoints"`
				LastWinPoints   int     `json:"lastWinPoints"`
				LongestKill     int     `json:"longestKill"`
				MostDamage      int     `json:"mostDamage"`
				Name            string  `json:"name"`
				PlayerID        string  `json:"playerId"`
				Revives         int     `json:"revives"`
				RideDistance    float32 `json:"rideDistance"`
				RoadKills       int     `json:"roadKills"`
				TeamKills       int     `json:"teamKills"`
				TimeSurvived    int     `json:"timeSurvived"`
				VehicleDestroys int     `json:"vehicleDestroys"`
				WalkDistance    float64 `json:"walkDistance"`
				WeaponsAcquired int     `json:"weaponsAcquired"`
				WinPlace        int     `json:"winPlace"`
				WinPoints       int     `json:"winPoints"`
				WinPointsDelta  float32 `json:"winPointsDelta"`
			} `json:"stats"`
		} `json:"attributes"`
		Relationships struct {
			Participants struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"participants"`
			Team struct {
				Data interface{} `json:"data"`
			} `json:"team"`
		} `json:"relationships,omitempty"`
	} `json:"included"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}
