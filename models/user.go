package models

type User struct {
	CreatedTime int64  `json:"created_time"`
	Role        string `json:"role"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required,max=20,min=6"`
}
