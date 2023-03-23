package user

import "github.com/api-sample/app/usecase/user"

type signUpReq struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (r signUpReq) toInput() user.SignUpInput {
	return user.SignUpInput{
		ID:       r.ID,
		Password: r.Password,
	}
}

type findByIDReq struct {
	ID string
}

func (r findByIDReq) toInput() user.FindByIDInput {
	return user.FindByIDInput{
		ID: r.ID,
	}
}
