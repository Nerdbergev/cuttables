package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/nerdbergev/cuttables/pkg/service"
)

func Handler(userSvc service.UserService, cuttableSvc service.CuttableService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", usersRouter(userSvc))
		r.Route("/cuttables", cuttablesRouter(cuttableSvc))
	})
	return r
}

func usersRouter(_ service.UserService) func(chi.Router) {
	return func(r chi.Router) {
	}
}

func cuttablesRouter(svc service.CuttableService) func(chi.Router) {
	h := cuttableHandler{svc: svc}
	return func(r chi.Router) {
		r.Get("/", h.handleGetAll)
	}
}

type cuttableHandler struct {
	svc service.CuttableService
}

func (h cuttableHandler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	cuttables, _ := h.svc.GetAll()
	render.JSON(w, r, cuttables)
}
