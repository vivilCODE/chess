package models

import "time"


type User struct {
	Id       string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"` 
	SignedUp time.Time `json:"signedUp"`
}