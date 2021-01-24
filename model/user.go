package model

//type User struct {
//	Id int `json:"id"`
//	Name string `json:"name"`
//	Token string `json:"token"`
//}

type UserCreateRequest struct {
	name string `json:"name"`
}

type UserCreateResponse struct {
	token string
}

type UserGetResponse struct {
	name string
}

type UserUpdateRequest struct {
	name string
}
