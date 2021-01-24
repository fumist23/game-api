package router

import (
	"github.com/fumist23/game-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/user/create", controller.CreateUser)

	return r
}
