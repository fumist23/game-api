package model

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Reality int    `json:"reality"`
}

type UserCharacter struct {
	UserID        int    `json:"userId"`
	CharacterID   int    `json:"characterId"`
	CharacterName string `json:"characterName"`
}
