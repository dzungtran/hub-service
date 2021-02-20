package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"hub-service/domain"
	"hub-service/pkg/core/servehttp"
	"hub-service/pkg/utils"
	"hub-service/usecase"
)

type CreateTeamHandler struct {
	UseCase usecase.CreateTeamUseCase
}

func (h *CreateTeamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input usecase.CreateTeamInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Error while unmarshal json input")
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid json input")
		return
	}

	// validate request
	if err := validateCreateTeam(input); err != nil {
		log.Printf("Validation error, detail: %v", err.Error())
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	teamOutput, err := h.UseCase.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Error while create user")
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, teamOutput)
	return
}

func validateCreateTeam(input usecase.CreateTeamInput) error {
	validator, _ := utils.NewGoPlayground()

	err := validator.Validate(input)
	if err != nil {
		return err
	}

	if len(validator.Messages()) > 0 {
		return errors.New(validator.Messages()[0])
	}

	// check type valid
	valid := false
	for _, r := range domain.AvailableTypes {
		if string(r) == input.Type {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New(fmt.Sprintf("Invalid type: %v", input.Type))
	}

	return nil
}
