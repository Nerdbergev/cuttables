package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/nerdbergev/cuttables/pkg/api"
	"github.com/nerdbergev/cuttables/pkg/service"
	"github.com/nerdbergev/cuttables/pkg/storage/memory"
)

func New() Server {
	return Server{}
}

type Server struct{}

func (s Server) ListenAndServe() error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		AllowedHeaders: []string{},
	}))

	userSvc := service.NewUserService()
	cuttableSvc := service.NewCuttableService(memory.NewCuttableStorage())

	router.Mount("/api", api.Handler(userSvc, cuttableSvc))

	return http.ListenAndServe(":8080", router)
}
