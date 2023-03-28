package user

type FindByIDOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LoginOutput struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}
