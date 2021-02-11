package controller

import (
	"log"
	"net/http"

	"github.com/fumist23/game-api/database"
)

func DrawGacha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	characters, err := database.GetCharacters(ctx)
	if err != nil {
		log.Printf("failed to get charatcers: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	gachaConfigs, err := database.GetGachaConfigs(ctx)
	if err != nil {
		log.Printf("failed to get gachaConfigs: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("characters: %v, gachaConfigs: %v", characters, gachaConfigs)

	w.WriteHeader(http.StatusOK)
}
