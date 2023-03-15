package handlers

import (
	"api/db/repository"
	model "api/models"
	"api/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ObtnerAsistencia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)
	codigoAsistencia := urlParams["codigo_asistencia"]
	asistencia, err := repository.ObtenerDatosAsistencia(codigoAsistencia)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorRes, _ := json.Marshal(model.ErrorRes{Mensaje: util.FormatearError(err.Error())})
		BadRequest(w, string(errorRes))
		return
	}

	resultado, err := json.Marshal(asistencia)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorRes, _ := json.Marshal(model.ErrorRes{Mensaje: util.FormatearError(err.Error())})
		BadRequest(w, string(errorRes))
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, string(resultado))
}

func BadRequest(w http.ResponseWriter, error string) {
	fmt.Fprintln(w, error)
}
