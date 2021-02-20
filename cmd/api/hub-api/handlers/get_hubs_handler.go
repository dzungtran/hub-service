package handlers

import (
	"net/http"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
	"strings"
)

type GetHubsHandler struct {
	UseCase usecase.GetHubsUseCase
}

func (h *GetHubsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	input, err := parseRequestToGetHubsInput(r)
	if err != nil {
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.UseCase.Execute(r.Context(), input)
	if err != nil {
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	servehttp.ResponseSuccessJSON(w, users)
	return
}

func parseRequestToGetHubsInput(r *http.Request) (usecase.GetHubsInput, error) {
	req := usecase.GetHubsInput{
		Ids: make([]int64, 0),
	}
	params := r.URL.Query()

	if name, ok := params["name"]; ok {
		req.Name = name[0]
	}

	if idStr, ok := params["ids"]; ok {
		idsArr := strings.Split(idStr[0], ",")
		for _, i := range idsArr {
			id, _ := strconv.ParseInt(i, 10, 64)
			if id > 0 {
				req.Ids = append(req.Ids, id)
			}
		}
	}

	return req, nil
}
