package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"hub-service/domain"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
)

type GetHubInfoHandler struct {
	UseCase usecase.GetHubInfoUseCase
}

func (h *GetHubInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hubId := getHubIdFromRequest(r)
	if hubId <= 0 {
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid hub id")
		return
	}

	team, err := h.UseCase.Execute(r.Context(), hubId)
	if err != nil {
		if err == domain.ErrorNotFound {
			servehttp.ResponseErrorJSON(w, http.StatusNotFound, "Hub not found")
			return
		}
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, team)
	return
}

func getHubIdFromRequest(r *http.Request) (hubId int64) {
	vars := mux.Vars(r)
	hubId, _ = strconv.ParseInt(vars["hub_id"], 10, 64)
	return hubId
}
