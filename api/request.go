package api

import (
	"encoding/json"
	"strings"
)

type Request struct {
}

// Types of items.
type Type int

const (
	Unknown Type = iota
	Psalm
	Acclamation
	Song
)

// Custom parsing for item type JSON.
func (_type *Type) UnmarshalJSON(source []byte) error {
	var str string
	if err := json.Unmarshal(source, &str); err != nil {
		return err
	}
	switch strings.ToLower(str) {
	case "psalm":
		*_type = Psalm
	case "acclamation":
		*_type = Acclamation
	case "song":
		*_type = Song
	default:
		*_type = Unknown
	}

	return nil
}

// Request item data.
type Item struct {
	Id      string   `json:"id"`
	Type    Type     `json:"type"`
	Content []string `json:"content"`
	Parts   []int    `json:"parts"`
}
