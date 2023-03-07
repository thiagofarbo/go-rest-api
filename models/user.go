package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	// Email       string
	// phoneNumber string
}

var Users []User
