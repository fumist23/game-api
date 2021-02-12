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

// tokenが存在するかチェックする
func VerifyToken(ctx context.Context, token string) bool {
	row := DB.QueryRowContext(ctx, "SELECT * FROM users WHERE token=?", token)
	if err := row.Err(); err != nil {
		log.Printf("this token is invalid")
		return false
	}

	return true
}

// tokenを受け取って該当するuserのnameを取り出す

func GetUser(ctx context.Context, token string) (string, error) {
	row := DB.QueryRowContext(ctx, "SELECT name FROM users WHERE token=?", token)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Printf("failed to get user name from database")
		return name, err
	}

	return name, nil

}

// tokenとnameを受け取ってtokenに該当するuserのnameを更新する
func UpdateUser(ctx context.Context, token string, name string) error {
	_, err := DB.QueryContext(ctx, "UPDATE users SET name = ? WHERE token = ? ", name, token)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return err
	}
	return nil
}
