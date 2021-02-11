package model

type Character struct {
	name    string `json:"name"`
	reality int    `json:"reality"`
}

type UserCharacter struct {
	userID  int `json:"userId"`
	reality int `json:"reality"`
}
