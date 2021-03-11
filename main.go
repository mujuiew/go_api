package main

import (
	"net/http"

	"github.com/go_api/internal/core/login"
	"github.com/go_api/internal/core/rest"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	rest.InitDB()
	r.HandleFunc("/login", login.Login).Methods("POST")
	r.HandleFunc("/getPromotion", rest.HomePage).Methods("POST")

	http.ListenAndServe(":8010", r)
}
