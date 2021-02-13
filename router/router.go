package router

import (
	"github.com/fumist23/game-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/user/create", controller.CreateUser)
	r.HandleFunc("/user/get", controller.GetUser)
	r.HandleFunc("/user/update", controller.UpdateUser)
	r.HandleFunc("/gacha/draw", controller.DrawGacha)

	return r
}
