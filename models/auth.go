package models

type Auth struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLogin struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	ProfileManagerId int    `json:"manager_id"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Action   string `json:"action" validate:"required,oneof=create login"`
}

type AuthLoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
