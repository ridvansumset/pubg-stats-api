package main

import (
	"example/pubg-stats-api/client"
	"example/pubg-stats-api/model"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID       string         `json:"id"`
	Type     model.DataType `json:"type"`
	Name     string         `json:"name"`
	MatchIDs []string       `json:"matchIds"`
}

func (s *HTTPServer) getPubgPlayerByID(c *gin.Context) {
	qs := url.Values{}
	qs.Add("filter[playerIds]", c.Param("id"))

	res, h := s.getPubgPlayer(qs)
	if res != nil {
		c.IndentedJSON(http.StatusOK, res)
	} else {
		c.IndentedJSON(http.StatusNotFound, h)
	}
}

func (s *HTTPServer) getPubgPlayerByName(c *gin.Context) {
	qs := url.Values{}
	qs.Add("filter[playerNames]", c.Param("name"))

	res, h := s.getPubgPlayer(qs)
	if res != nil {
		c.IndentedJSON(http.StatusOK, res)
	} else {
		c.IndentedJSON(http.StatusNotFound, h)
	}
}

func (s *HTTPServer) getPubgPlayer(queryString url.Values) (*resWrapper, *gin.H) {
	qd := time.Now().Unix()
	data, err := s.http.Request(client.ShardSteam, client.EndpointPlayers, &queryString)
	if err != nil {
		return nil, &gin.H{"message": err.Error()}
	}
	if data == nil {
		return nil, &gin.H{"message": "nil response"}
	}
	defer data.Close()

	r := model.PlayerResp{}
	if err := fastjson.NewDecoder(data).Decode(&r); err != nil {
		return nil, &gin.H{"message": "could not decode response"}
	}

	u := r.Data[0]
	var res = resWrapper{
		QueryTimestamps: []int64{qd},
		Player: &player{
			ID:   u.ID,
			Type: u.Type,
			Name: u.Attributes.Name,
			MatchIDs: func() []string {
				ms := u.Relationships.Matches.Data
				var ids []string
				for i := 0; i < len(ms); i++ {
					ids = append(ids, ms[i].ID)
				}
				return ids
			}(),
		},
	}

	return &res, nil
}
