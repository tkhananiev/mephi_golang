package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"rest_module/rest"
)

func Router(api *rest.API) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is alive"))
	}).Methods("GET")

	// Подключаем все реальные маршруты
	r.PathPrefix("/").Handler(api.Router())

	return r
}
