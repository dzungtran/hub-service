package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"hub-service/adapter/repository"
	"hub-service/cmd/api/hub-api/handlers"
	"hub-service/infra"
	"hub-service/pkg/core/servehttp"
	"hub-service/usecase"
	"time"
)

func main() {

	// Init App server
	appServer := &servehttp.AppServer{}
	appServer.Init()

	db, err := infra.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	for _, handler := range getListHandler(db) {
		appServer.RegisterHandler(handler.Method, handler.Route, handler.Handler)
	}

	localPort := os.Getenv("APP_PORT")
	srv := &http.Server{
		Handler:      appServer.GetRouter(),
		Addr:         fmt.Sprintf(":%v", localPort),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println(fmt.Sprintf("Starting API server with port :%v", localPort))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	servehttp.WaitForShutdown(srv)
}

// Defines list handler to serve requests
func getListHandler(db *sql.DB) []servehttp.AppHandler {

	userRepo := repository.NewUserPostgresRepository(db)
	hubRepo := repository.NewHubPostgresRepository(db)
	teamRepo := repository.NewTeamPostgresRepository(db)

	return []servehttp.AppHandler{
		{
			Route:   "/",
			Method:  http.MethodGet,
			Handler: &handlers.GetHelloHandler{},
		},

		// ======================================
		// User api endpoints
		// Get all user api
		{
			Route:  "/users",
			Method: http.MethodGet,
			Handler: &handlers.GetAllUserHandler{
				UseCase: usecase.NewGetAllUserUseCase(userRepo),
			},
		},
		// Create user api
		{
			Route:  "/users",
			Method: http.MethodPost,
			Handler: &handlers.CreateUserHandler{
				UseCase: usecase.NewCreateUserUseCase(userRepo),
			},
		},
		// Get user info
		{
			Route:  "/users/:user_id",
			Method: http.MethodGet,
			Handler: &handlers.GetUserInfoHandler{
				UseCase: usecase.NewGetUserInfoUseCase(userRepo),
			},
		},

		// ======================================
		// Hub API endpoints
		// Create a hub api
		{
			Route:  "/hubs",
			Method: http.MethodPost,
			Handler: &handlers.CreateHubHandler{
				UseCase: usecase.NewCreateHubUseCase(hubRepo),
			},
		},
		// Find hubs
		{
			Route:  "/hubs",
			Method: http.MethodGet,
			Handler: &handlers.GetHubsHandler{
				UseCase: usecase.NewGetHubsUseCase(hubRepo),
			},
		},
		// Get hub info by id
		{
			Route:  "/hubs/{hub_id:[0-9]+}",
			Method: http.MethodGet,
			Handler: &handlers.GetHubInfoHandler{
				UseCase: usecase.NewGetHubInfoUseCase(hubRepo),
			},
		},
		// Get all teams in a hub
		{
			Route:  "/hubs/{hub_id:[0-9]+}/teams",
			Method: http.MethodGet,
			Handler: &handlers.GetTeamsHandler{
				UseCase: usecase.NewGetTeamsUseCase(teamRepo),
			},
		},

		// ========================================
		// Team API endpoints
		// Find teams API
		{
			Route:  "/teams",
			Method: http.MethodGet,
			Handler: &handlers.GetTeamsHandler{
				UseCase: usecase.NewGetTeamsUseCase(teamRepo),
			},
		},
		// Create team API
		{
			Route:  "/teams",
			Method: http.MethodPost,
			Handler: &handlers.CreateTeamHandler{
				UseCase: usecase.NewCreateTeamUseCase(teamRepo, hubRepo),
			},
		},
		// Add Users to team
		{
			Route:  "/teams/{team_id:[0-9]+}/add-users",
			Method: http.MethodPut,
			Handler: &handlers.TeamAddUsersHandler{
				UseCase: usecase.NewTeamAddUsersUseCase(teamRepo, userRepo),
			},
		},
		// Get a team info
		{
			Route:  "/teams/{team_id:[0-9]+}",
			Method: http.MethodGet,
			Handler: &handlers.GetTeamInfoHandler{
				UseCase: usecase.NewGetTeamInfoUseCase(teamRepo),
			},
		},
		// Get all users in a team
		{
			Route:  "/teams/{team_id:[0-9]+}/users",
			Method: http.MethodGet,
			Handler: &handlers.GetAllUserHandler{
				UseCase: usecase.NewGetAllUserUseCase(userRepo),
			},
		},
	}
}
