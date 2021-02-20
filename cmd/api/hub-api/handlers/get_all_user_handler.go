package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"hub-service/domain"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
)

type GetAllUserHandler struct {
	UseCase usecase.GetAllUserUseCase
}

func (h *GetAllUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.UseCase.Execute(r.Context(), parseRequestToFindUsersRequest(r))
	if err != nil {
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	servehttp.ResponseSuccessJSON(w, users)
	return
}

func parseRequestToFindUsersRequest(r *http.Request) domain.FindUsersRequest {
	vars := mux.Vars(r)
	teamId, _ := strconv.ParseInt(vars["team_id"], 10, 64)
	req := domain.FindUsersRequest{}

	if teamId > 0 {
		req.TeamId = teamId
	}
	return req
}
