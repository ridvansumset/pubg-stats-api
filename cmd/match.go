package main

import (
	"fmt"
	"net/http"
	"time"

	"example/pubg-stats-api/client"
	"example/pubg-stats-api/model"

	"github.com/gin-gonic/gin"
)

type match struct {
	ID          string         `json:"id"`
	Type        model.DataType `json:"type"`
	MapName     string         `json:"mapName"`
	CreatedAt   time.Time      `json:"createdAt"`
	GameMode    string         `json:"gameMode"`
	TitleID     string         `json:"titleId"`
	Duration    int            `json:"duration"`
	Participant *participant   `json:"participant,omitempty"`
}

type participant struct {
	ID            string         `json:"id"`
	Type          model.DataType `json:"type"`
	Assists       int            `json:"assists"`
	DamageDealt   float64        `json:"damageDealt"`
	DeathType     string         `json:"deathType"`
	HeadshotKills int            `json:"headshotKills"`
	KillPlace     int            `json:"killPlace"`
	KillStreaks   int            `json:"killStreaks"`
	Kills         int            `json:"kills"`
	LongestKill   float64        `json:"longestKill"`
	Name          string         `json:"name"`
	PlayerId      string         `json:"playerId"`
	RideDistance  float64        `json:"rideDistance"`
	TimeSurvived  int            `json:"timeSurvived"`
	WalkDistance  float64        `json:"walkDistance"`
	WinPlace      int            `json:"winPlace"`
}

func (s *HTTPServer) getPubgMatchByID(c *gin.Context) {
	qd := time.Now().Unix()
	data, err := s.http.Request(client.ShardSteam, fmt.Sprintf(client.EndpointMatches, c.Param("id")), nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if data == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "nil response"})
		return
	}
	defer data.Close()

	r := model.MatchResp{}
	if err := fastjson.NewDecoder(data).Decode(&r); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not decode response"})
		return
	}

	m := r.Data
	var resp = resWrapper{
		QueryTimestamps: []int64{qd},
		Match: &match{
			ID:        m.ID,
			Type:      m.Type,
			MapName:   m.Attributes.MapName,
			CreatedAt: m.Attributes.CreatedAt,
			GameMode:  m.Attributes.GameMode,
			TitleID:   m.Attributes.TitleID,
			Duration:  m.Attributes.Duration,
		},
	}
	if c.Query("participantId") != "" || c.Query("participantName") != "" {
		resp.Match.Participant = func() *participant {
			var ps []model.MatchParticipant
			for i := 0; i < len(r.Included); i++ {
				if r.Included[i].Type == model.DataTypeParticipant {
					ps = append(ps, r.Included[i])
				}
			}

			var mp *model.MatchParticipant
			for i := range ps {
				if ps[i].Attributes.Stats.PlayerId == c.Query("participantId") ||
					ps[i].Attributes.Stats.Name == c.Query("participantName") {
					mp = &ps[i]
					break
				}
			}

			if mp != nil {
				stats := mp.Attributes.Stats
				return &participant{
					ID:            mp.ID,
					Type:          mp.Type,
					Assists:       stats.Assists,
					DamageDealt:   stats.DamageDealt,
					DeathType:     stats.DeathType,
					HeadshotKills: stats.HeadshotKills,
					KillPlace:     stats.KillPlace,
					KillStreaks:   stats.KillStreaks,
					Kills:         stats.Kills,
					LongestKill:   stats.LongestKill,
					Name:          stats.Name,
					PlayerId:      stats.PlayerId,
					RideDistance:  stats.RideDistance,
					TimeSurvived:  stats.TimeSurvived,
					WalkDistance:  stats.WalkDistance,
					WinPlace:      stats.WinPlace,
				}
			}

			return nil
		}()
	}

	c.IndentedJSON(http.StatusOK, resp)
}
