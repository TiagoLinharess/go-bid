package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	// csrfMiddleware := csrf.Protect(
	// 	[]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Path("/api/v1"),
	// 	csrf.Secure(true), // DEV ONLY
	// )
	// api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// r.Get("/csrf_token", api.HandleGetCSFRToken)
			r.Route("/users/", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLogin)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogout)
			})
		})
	})
}
