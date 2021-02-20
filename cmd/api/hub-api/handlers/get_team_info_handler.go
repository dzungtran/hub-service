package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"hub-service/domain"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
)

type GetTeamInfoHandler struct {
	UseCase usecase.GetTeamInfoUseCase
}

func (h *GetTeamInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teamId := getTeamIdFromRequest(r)
	if teamId <= 0 {
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid team id")
		return
	}

	team, err := h.UseCase.Execute(r.Context(), teamId)
	if err != nil {
		if err == domain.ErrorNotFound {
			servehttp.ResponseErrorJSON(w, http.StatusNotFound, "Team not found")
			return
		}
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, team)
	return
}

func getTeamIdFromRequest(r *http.Request) (teamId int64) {
	vars := mux.Vars(r)
	teamId, _ = strconv.ParseInt(vars["team_id"], 10, 64)
	return teamId
}
