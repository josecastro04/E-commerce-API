package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
}
