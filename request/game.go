package request

import (
	"errors"
	"strings"
	"time"

	"github.com/feribeirods/api-go-gin/model"
)

type GameRequest struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	PublishDate time.Time  `json:"publish_date"`
	Beated      *bool      `json:"beated"`
	BeatedAt    *time.Time `json:"beated_at"`
	Rating      int64      `json:"rating"`
}

// Validate the values recieved
func (r *GameRequest) Validate() []error {
	var errs []error

	if strings.TrimSpace(r.Name) == "" {
		errs = append(errs, errors.New("you must send us a name"))
	}
	if r.Rating < 0 || r.Rating > 10 {
		errs = append(errs, errors.New("rating must be between 0 and 10"))
	}
	if r.Beated != nil && *r.Beated && r.BeatedAt == nil {
		errs = append(errs, errors.New("beated_at is required if beated is true"))
	}
	if r.PublishDate.After(time.Now()) {
		errs = append(errs, errors.New("publish_date cannot be in the future"))
	}
	return errs
}

// Translate to the model
func (r *GameRequest) ToModel() *model.Game {
	return &model.Game{
		ID:          r.ID,
		Name:        r.Name,
		PublishDate: r.PublishDate,
		Beated:      r.Beated,
		BeatedAt:    r.BeatedAt,
		Rating:      r.Rating,
	}
}
