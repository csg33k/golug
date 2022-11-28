package main

import (
	"context"
	"github.com/safaci2000/golug/services"
	"net/http"

	"github.com/safaci2000/golug/config"
	"github.com/safaci2000/golug/www"
	log "github.com/sirupsen/logrus"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// @title Golug Demo Code
// @version 1.0
// @BasePath /
//
//go:generate rm -fr dbmodels
//go:generate sqlc generate
//go:generate swag init --parseDependency -o ./api -g main.go

var (
	appConfig *config.ServerConfig
)

func main() {
	wwwSrv := services.InitializeServices(appConfig.Server.DatabaseUri)
	ctx := context.Background()
	context.WithValue(ctx, "db", wwwSrv)
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//User Route
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/list", www.ListUsers)
		r.Get("/{id}", www.GetUser)
		r.Delete("/{id}", www.DeleteUser)
		r.Post("/", www.CreateUser)
		r.Put("/{id}", www.UpdateUser)
	})
	r.Get("/", www.Redirect("/swaggerui/"))

	//Static Mount generated files
	fs := http.FileServer(http.Dir("api"))
	r.Handle("/swagger/*", http.StripPrefix("/swagger/", fs))

	//Loads Swagger UI
	r.Get("/swaggerui/*", httpSwagger.Handler(
		//httpSwagger.URL(fmt.Sprintf("/swagger/swagger.json", appConfig.Server.Address)), //The url pointing to API definition
		httpSwagger.URL("/swagger/swagger.json"), //The url pointing to API definition
	))

	log.Infof("Starting server on: '%s'", appConfig.Server.Address)
	http.ListenAndServe(appConfig.Server.Address, r)
}

func init() {
	config.NewConfig()
	appConfig = config.LoadConfig("golug")
}
