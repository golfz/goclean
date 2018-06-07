package router

import (
	"github.com/gorilla/mux"
	"goclean/handler"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/example/{key}", handler.InitiateExampleHandler().GetExampleData).Methods("GET")

	return r
}
