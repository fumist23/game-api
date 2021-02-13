package controller

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/fumist23/game-api/database"
	"github.com/fumist23/game-api/model"
)

// 指定された回数分、ランダムにキャラクターを返す
func getRandomCharacters(ctx context.Context, count int) ([]model.Character, error) {
	// 対象となるキャラクターの取得
	characters, err := database.GetCharacters(ctx)
	if err != nil {
		log.Printf("failed to get charatcers: %v", err)
		return nil, err
	}

	// ガチャの設定を取得
	gachaConfigs, err := database.GetGachaConfigs(ctx)
	if err != nil {
		log.Printf("failed to get gachaConfigs: %v", err)
		return nil, err
	}

	var groupedCharactersList []model.GroupedCharacters
	// realityごとにキャラクターをグループ化する
	for _, gachaConfig := range gachaConfigs {
		ids := make([]int, 0)
		for _, character := range characters {
			if character.Reality == gachaConfig.Reality {
				ids = append(ids, character.ID)
			}
		}
		groupedCharacters := model.GroupedCharacters{
			Reality:     gachaConfig.Reality,
			Probability: gachaConfig.Probability,
			IDs:         ids,
		}
		groupedCharactersList = append(groupedCharactersList, groupedCharacters)
	}

	rand.Seed(time.Now().UnixNano())

	selectedCharacterIDs := make([]int, 0)

	for i := 0; i < count; i++ {
		randomNum := rand.Float64()
		selectedCharacterID := gacha(groupedCharactersList, randomNum)
		selectedCharacterIDs = append(selectedCharacterIDs, selectedCharacterID)
	}

	selectedCharacters := make([]model.Character, 0, count)

	for _, selectedCharacterID := range selectedCharacterIDs {
		for _, character := range characters {
			if character.ID == selectedCharacterID {
				selectedCharacters = append(selectedCharacters, character)
				break
			}
		}
	}
	return selectedCharacters, nil
}

// グループ化されたキャラクターとランダムな数字からキャラクターを選び、IDだけ返す
func gacha(groupedCharactersList []model.GroupedCharacters, randomNum float64) int {
	var accum float64 = 0
	var selectedCharacterID int

	for _, groupedCharacters := range groupedCharactersList {
		accum += groupedCharacters.Probability
		if randomNum < accum {
			selectedCharacterID = groupedCharacters.IDs[int(math.Round(randomNum*float64(len(groupedCharacters.IDs)-1)))]
			break
		}
	}

	return selectedCharacterID
}

// DrawGacha ガチャを引く
func DrawGacha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("x-token")

	//tokenの検証
	isValidToken := database.VerifyToken(ctx, token)
	if !isValidToken {
		log.Printf("this token is invalid")
		w.WriteHeader(http.StatusBadRequest)
	}

	body := r.Body
	defer body.Close()

	// ガチャを引く回数
	var gachaDrawRequest model.GachaDrawRequest
	if err := json.NewDecoder(body).Decode(&gachaDrawRequest); err != nil {
		log.Printf("failed to decode json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	count := gachaDrawRequest.Count

	if count <= 0 {
		log.Printf("time must be more than 1")
		w.WriteHeader(http.StatusBadRequest)
	}

	//引かれたキャラクターを取得
	selectedCharacters, err := getRandomCharacters(ctx, count)
	if err != nil {
		log.Printf("failed to getRandomCharacters: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// ガチャを引くユーザーの情報を取得
	user, err := database.GetUser(ctx, token)
	if err != nil {
		log.Printf("failoed to get user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 取得したキャラクターをuserCharacterテーブルに入れる
	if err := database.PostUserCharacters(ctx, selectedCharacters, user.Id); err != nil {
		log.Printf("failed to PostUserCharacters: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	//引いたキャラクターを返す
	if err := json.NewEncoder(w).Encode(selectedCharacters); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
