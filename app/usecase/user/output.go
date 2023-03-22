package user

import "time"

type FindByIdOutput struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birth_day"`
}
