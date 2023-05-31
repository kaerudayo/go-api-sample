package user

import "github.com/api-sample/app/usecase/user"

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r loginReq) toInput() user.LoginInput {
	return user.LoginInput{
		Email:    r.Email,
		Password: r.Password,
	}
}

type signUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r signUpReq) toInput() user.SignUpInput {
	return user.SignUpInput{
		Email:    r.Email,
		Password: r.Password,
	}
}

type findByIDReq struct {
	ID string
}

func (r findByIDReq) toInput() user.FindByIDInput {
	return user.FindByIDInput{
		UserID: r.ID,
	}
}
