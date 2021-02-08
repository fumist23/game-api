package controller

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fumist23/game-api/database"
	"github.com/fumist23/game-api/model"
)

// /user/createに対するハンドラ
// requestからnameを取り出してtoken生成してDBに保存して返す
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body := r.Body
	defer body.Close()

	var userCreateRequest model.UserCreateRequest
	json.NewDecoder(body).Decode(&userCreateRequest)

	name := userCreateRequest.Name

	tokenStr, err := GenerateTokenWithName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred in GenerateTokenWithName %v", err)
	}

	err = database.CreateUser(ctx, name, tokenStr)
	if err != nil {
		log.Printf("error occurred in database.requestUser: %v", err)
	}

	userCreateResponse := model.UserCreateResponse{
		Token: tokenStr,
	}

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(userCreateResponse)
	if err != nil {
		log.Printf("failed to encode userCreateResponse: %v", err)
	}

	w.Write(buf.Bytes())
	w.WriteHeader(http.StatusCreated)
}

// /user/getに対するハンドラ
// headerのx-tokenからtokenを取り出してDBからfetchして該当するuserのnameを取得して返す
//func GetUser(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	token := r.Header.Get("x-token")
//	if token == "" {
//		log.Println("token must be included in header")
//		w.WriteHeader(http.StatusBadRequest)
//	}
//
//	name, err := database.GetUser(ctx, token)
//
//	w
//
//}

// /user/updateに対するハンドラ
// x-tokenからtokenを取り出して該当するuserを検証し、受け取ったnameを更新してDB更新して返す
