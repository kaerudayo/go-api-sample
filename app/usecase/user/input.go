package user

type LoginInput struct {
	Email    string
	Password string
}

type SignUpInput struct {
	Email    string
	Password string
}

type FindByIDInput struct {
	ID string
}
