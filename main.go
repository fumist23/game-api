package main

import (
	"net/http"
	"practices/github.com/labstack/gommon/log"

)

func main() {

	r := router.Router()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}