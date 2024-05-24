package main

import (
	"net/http"

	"github.com/driv/golang-apiserver/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) root(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": "Welcome to the API!",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
