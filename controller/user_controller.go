package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
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

// /user/updateに対するハンドラ
// x-tokenからtokenを取り出して該当するuserを検証し、受け取ったnameを更新してDB更新して返す
