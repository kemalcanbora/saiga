package models

type GetChatHistory struct {
	ID string `json:"id" validate:"required"`
}
