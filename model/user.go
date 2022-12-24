package model

type PubgType string

const (
	PubgTypePlayer = PubgType("player")
	PubgTypeMatch  = PubgType("match")
)

type PubgResp struct {
	Data  []PubgUser     `json:"data"`
	Links Links          `json:"links"`
	Meta  map[string]any `json:"meta"`
}

type PubgUser struct {
	ID            string     `json:"id"`
	Type          PubgType   `json:"type"`
	Attributes    Attributes `json:"attributes"`
	Relationships struct {
		Assets  Assets  `json:"assets"`
		Matches Matches `json:"matches"`
	} `json:"relationships"`
	Links Links `json:"links"`
}

type Attributes struct {
	Stats        any    `json:"stats"`
	TitleId      string `json:"titleId"`
	ShardId      string `json:"shardId"`
	PatchVersion string `json:"patchVersion"`
	Name         string `json:"name"`
}

type Assets struct {
	Data []Data `json:"data"`
}

type Matches struct {
	Data []Data `json:"data"`
}

type Data struct {
	ID   string   `json:"id"`
	Type PubgType `json:"type"`
}

type Links struct {
	Self   string  `json:"self"`
	Schema *string `json:"schema,omitempty"`
}
