package model

type SignInRequest struct {
	Code string `json:"code"`
}

type SignInResponse struct {
	// This should also contain a token that keeps the user logged in for a while
	User User
}
