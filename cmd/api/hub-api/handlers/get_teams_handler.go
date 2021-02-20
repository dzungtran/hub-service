package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"strconv"
	"strings"
)

type GetTeamsHandler struct {
	UseCase usecase.GetTeamsUseCase
}

func (h *GetTeamsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	input, err := parseRequestToGetTeamsInput(r)
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

func parseRequestToGetTeamsInput(r *http.Request) (usecase.GetTeamsInput, error) {
	result := usecase.GetTeamsInput{
		Ids:   make([]int64, 0),
		Types: make([]string, 0),
	}
	params := r.URL.Query()
	vars := mux.Vars(r)
	hubId, _ := strconv.ParseInt(vars["hub_id"], 10, 64)

	if name, ok := params["name"]; ok {
		result.Name = name[0]
	}

	if idStr, ok := params["ids"]; ok {
		idsArr := strings.Split(idStr[0], ",")
		for _, i := range idsArr {
			id, _ := strconv.ParseInt(i, 10, 64)
			if id > 0 {
				result.Ids = append(result.Ids, id)
			}
		}
	}

	if typeStr, ok := params["types"]; ok {
		typeArr := strings.Split(typeStr[0], ",")
		for _, i := range typeArr {
			if len(i) > 0 {
				result.Types = append(result.Types, i)
			}
		}
	}

	if hubId > 0 {
		result.HubId = hubId
	}

	return result, nil
}
