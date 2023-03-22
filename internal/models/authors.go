package models

type AuthorNotifications struct {
	Email bool `json:"email"`
	Sms   bool `json:"sms"`
}
