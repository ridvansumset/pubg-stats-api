package main

import (
	jsoniter "github.com/json-iterator/go"
)

type resWrapper struct {
	QueryTimestamps []int64 `json:"queryTimestamps"`
	Player          *player `json:"player,omitempty"`
	Match           *match  `json:"match,omitempty"`
}

var fastjson = jsoniter.ConfigCompatibleWithStandardLibrary
