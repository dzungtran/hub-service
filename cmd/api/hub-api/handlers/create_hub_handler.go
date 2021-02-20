package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"hub-service/pkg/core/servehttp"
	"hub-service/pkg/utils"
	"hub-service/usecase"
)

type CreateHubHandler struct {
	UseCase usecase.CreateHubUseCase
}

func (h *CreateHubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var input usecase.CreateHubInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Error while unmarshal json input")
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, "Invalid json input")
		return
	}

	// validate request
	if err := validateCreateHub(input); err != nil {
		log.Printf("Validation error, detail: %v", err.Error())
		servehttp.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	hubOutput, err := h.UseCase.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Error while create hub")
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	servehttp.ResponseSuccessJSON(w, hubOutput)
	return
}

func validateCreateHub(input usecase.CreateHubInput) error {
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
