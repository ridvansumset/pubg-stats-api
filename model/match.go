package model

import "time"

type MatchResp struct {
	Data     Match              `json:"data"`
	Included []MatchParticipant `json:"included"`
	Links    Links              `json:"links"`
	Meta     map[string]any     `json:"meta,omitempty"`
}

type Match struct {
	ID            string          `json:"id"`
	Type          DataType        `json:"type"`
	Attributes    MatchAttributes `json:"attributes"`
	Relationships struct {
		Rosters DataWrapper `json:"rosters"`
		Assets  DataWrapper `json:"assets"`
	} `json:"relationships"`
	Links Links `json:"links"`
}

type MatchAttributes struct {
	MapName       string    `json:"mapName"`
	MatchType     string    `json:"matchType"`
	SeasonState   string    `json:"seasonState"`
	CreatedAt     time.Time `json:"createdAt"`
	Tags          any       `json:"tags"`
	GameMode      string    `json:"gameMode"`
	TitleID       string    `json:"titleId"`
	ShardID       string    `json:"shardId"`
	IsCustomMatch bool      `json:"isCustomMatch"`
	Duration      int       `json:"duration"`
	Stats         any       `json:"stats"`
}

type MatchParticipant struct {
	ID         string                     `json:"id"`
	Type       DataType                   `json:"type"`
	Attributes MatchParticipantAttributes `json:"attributes"`
}

type MatchParticipantAttributes struct {
	Stats struct {
		DBNOs           int     `json:"DBNOs"`
		Assists         int     `json:"assists"`
		Boosts          int     `json:"boosts"`
		DamageDealt     float64 `json:"damageDealt"`
		DeathType       string  `json:"deathType"`
		HeadshotKills   int     `json:"headshotKills"`
		Heals           int     `json:"heals"`
		KillPlace       int     `json:"killPlace"`
		KillStreaks     int     `json:"killStreaks"`
		Kills           int     `json:"kills"`
		LongestKill     float64 `json:"longestKill"`
		Name            string  `json:"name"`
		PlayerId        string  `json:"playerId"`
		Revives         int     `json:"revives"`
		RideDistance    float64 `json:"rideDistance"`
		RoadKills       int     `json:"roadKills"`
		SwimDistance    float64 `json:"swimDistance"`
		TeamKills       int     `json:"teamKills"`
		TimeSurvived    int     `json:"timeSurvived"`
		VehicleDestroys int     `json:"vehicleDestroys"`
		WalkDistance    float64 `json:"walkDistance"`
		WeaponsAcquired int     `json:"weaponsAcquired"`
		WinPlace        int     `json:"winPlace"`
	} `json:"stats"`
	Actor string `json:"actor"`
}
