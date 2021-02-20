package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"hub-service/pkg/core/servehttp"
	"hub-service/pkg/utils"
	"hub-service/usecase"
	"strconv"
)

type TeamAddUsersHandler struct {
	UseCase usecase.TeamAddUsersUseCase
}

func (h *TeamAddUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var input usecase.TeamAddUsersInput
	vars := mux.Vars(r)
	teamId, err := strconv.ParseInt(vars["team_id"], 10, 64)
	if err != nil {
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid team id")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Error while unmarshal json input")
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid json input")
		return
	}

	// validate request
	if err := validateTeamAddUsers(input); err != nil {
		log.Printf("Validation error, detail: %v", err.Error())
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	input.TeamId = teamId
	err = h.UseCase.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Error while create hub")
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
	})
	return
}

func validateTeamAddUsers(input usecase.TeamAddUsersInput) error {
	validator, _ := utils.NewGoPlayground()

	err := validator.Validate(input)
	if err != nil {
		return err
	}

	if len(validator.Messages()) > 0 {
		return errors.New(validator.Messages()[0])
	}

	return nil
}
