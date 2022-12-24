package main

import (
	"net/http"
	"net/url"

	"example/pubg-stats-api/model"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func (s *HTTPServer) getPubgUserByID(c *gin.Context) {
	qs := url.Values{}
	qs.Add("filter[playerIds]", c.Param("id"))

	data, err := s.http.Request("/steam/players", qs)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if data == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "nil response"})
		return
	}
	defer data.Close()

	r := model.PubgResp{}
	if err := fastjson.NewDecoder(data).Decode(&r); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not decode response"})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func (s *HTTPServer) getPubgUserByName(c *gin.Context) {
	qs := url.Values{}
	qs.Add("filter[playerNames]", c.Param("name"))

	data, err := s.http.Request("/steam/players", qs)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if data == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "nil response"})
		return
	}
	defer data.Close()

	r := model.PubgResp{}
	if err := fastjson.NewDecoder(data).Decode(&r); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not decode response"})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

var fastjson = jsoniter.ConfigCompatibleWithStandardLibrary
