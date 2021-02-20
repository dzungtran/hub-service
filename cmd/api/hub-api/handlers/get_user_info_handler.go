package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"hub-service/domain"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
)

type GetUserInfoHandler struct {
	UseCase usecase.GetUserInfoUseCase
}

func (h *GetUserInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromRequest(r)
	if userId <= 0 {
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	user, err := h.UseCase.Execute(r.Context(), userId)
	if err != nil {
		if err == domain.ErrorNotFound {
			servehttp.ResponseErrorJSON(w, http.StatusNotFound, "User not found")
			return
		}
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, user)
	return
}

func getUserIdFromRequest(r *http.Request) (userId int64) {
	vars := mux.Vars(r)
	userId, _ = strconv.ParseInt(vars["user_id"], 10, 64)
	return userId
}
