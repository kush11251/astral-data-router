package models

import (
	"encoding/json"
)

type Data struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}
