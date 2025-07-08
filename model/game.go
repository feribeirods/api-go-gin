package model

import "time"

type Game struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	PublishDate time.Time  `json:"publish_date"`
	Beated      *bool      `json:"beated"`
	BeatedAt    *time.Time `json:"beated_at"`
	Rating      int64      `json:"rating"`
}
