package main

import (
	"api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	var pathApi = "/api/"
	mux := mux.NewRouter()
	mux.HandleFunc(pathApi+"asistencia/{codigo_asistencia}", handlers.ObtnerAsistencia).Methods(http.MethodGet)
	return mux
}
