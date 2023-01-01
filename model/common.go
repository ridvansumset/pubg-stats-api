package model

type DataType string

const (
	DataTypePlayer      = DataType("player")
	DataTypeMatch       = DataType("match")
	DataTypeParticipant = DataType("participant")
	DataTypeRoster      = DataType("roster")
)

type DataWrapper struct {
	Data []Data `json:"data"`
}

type Data struct {
	ID   string   `json:"id"`
	Type DataType `json:"type"`
}

type Links struct {
	Self   string  `json:"self"`
	Schema *string `json:"schema,omitempty"`
}
