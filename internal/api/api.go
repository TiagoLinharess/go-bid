package api

import (
	"gobid/internal/services"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router          *chi.Mux
	UserService     services.UserService
	ProductsService services.ProductsService
	Sessions        *scs.SessionManager
}
