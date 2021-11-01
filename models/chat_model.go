package models

type ChatMeta struct {
	ID       string `bson:"_id"`
	Messages []Chat
}

type Attachment struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type Chat struct {
	Type        string     `json:"type"`
	Text        string     `json:"text"`
	From        string     `json:"from"`
	SenderName  string     `json:"sender_name"`
	To          string     `json:"to"`
	CreatedTime string     `json:"created_time"`
	Attachment  Attachment `json:"attachment"`
}
