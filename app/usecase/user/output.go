package user

import "time"

type FindByIDOutput struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
}

type LoginOutput struct {
	ID          string
	AccessToken string
}
