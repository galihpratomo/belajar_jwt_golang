package controllers

import (
	"net/http"

	"github.com/galihpratomo/tes_jwt/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Succes API")

}