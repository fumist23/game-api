package database

import (
	"context"
	"log"
)

// nameをtokenを受け取って保存する
func CreateUser(ctx context.Context, name string, token string) error {
	_, err := DB.QueryContext(ctx, "INSERT INTO users(name, token) VALUES(?, ?)", name, token)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return err
	}

	return nil
}

// tokenを受け取って該当するuserのnameを取り出す

// tokenとnameを受け取ってtokenに該当するuserのnameを更新する
