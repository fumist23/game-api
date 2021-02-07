package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fumist23/game-api/database"
	"github.com/fumist23/game-api/model"
)

// /user/createに対するハンドラ
// requestからnameを取り出してtoken生成してDBに保存して返す
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	defer body.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var userCreateRequest model.UserCreateRequest
	json.Unmarshal(buf.Bytes(), &userCreateRequest)

	name := userCreateRequest.name

	tokenStr, err := GenerateTokenWithName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Pringf("error occurred in GenerateTokenWithName %v", err)
	}

	err = database.CreateUser(ctx, name, tokenStr)
	if err != nil {
		log.Printf("error occured in Cdatabase.reateUser: %v", err)
	}

	userCreateResponse := model.UserCreateResponse{
		token: token,
	}

	w.Write([]byte(userCreateResponse))
	w.WriteHeader(http.StatusCreated)
}

// /user/getに対するハンドラ
// headerのx-tokenからtokenを取り出してDBからfetchして該当するuserのnameを取得して返す
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("token must be included in header")
		w.WriteHeader(http.StatusBadRequest)
	}

	name, err := database.GetUser(ctx, token)

	if err != nil {
		log.Printf("error occurred in database.GetUser: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(name); err != nil {
		log.Printf("failed to encode name to json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// /user/updateに対するハンドラ
// x-tokenからtokenを取り出して該当するuserを検証し、受け取ったnameを更新してDB更新して返す
