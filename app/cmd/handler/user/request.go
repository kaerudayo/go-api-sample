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

type findByIdReq struct {
	ID string
}

func (r findByIdReq) toInput() user.FindByIdInput {
	return user.FindByIdInput{
		ID: r.ID,
	}
}
