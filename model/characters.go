package model

type Character struct {
	Name    string `json:"name"`
	Reality int    `json:"reality"`
}

type UserCharacter struct {
	UserID  int `json:"userId"`
	Reality int `json:"reality"`
}
