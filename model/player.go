package model

type PlayerResp struct {
	Data  []Player       `json:"data"`
	Links Links          `json:"links"`
	Meta  map[string]any `json:"meta,omitempty"`
}

type Player struct {
	ID            string           `json:"id"`
	Type          DataType         `json:"type"`
	Attributes    PlayerAttributes `json:"attributes"`
	Relationships struct {
		Assets  DataWrapper `json:"assets"`
		Matches DataWrapper `json:"matches"`
	} `json:"relationships"`
	Links Links `json:"links"`
}

type PlayerAttributes struct {
	Stats        any    `json:"stats,omitempty"`
	TitleID      string `json:"titleId"`
	ShardID      string `json:"shardId"`
	PatchVersion string `json:"patchVersion"`
	Name         string `json:"name"`
}
