package models

type JwtUserAuth struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Exp        string `json:"exp"`
	Role       string `json:"role"`
}
