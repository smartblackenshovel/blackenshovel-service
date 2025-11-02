package routes

import (
	"net/http"

	"blackenshovel-service/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Organizations routes
	r.Get("/organizations", controllers.GetOrganizations)
	r.Post("/organizations", controllers.CreateOrganization)

	// Example: placeholder for future routes
	// r.Get("/organizations/{id}", controllers.GetOrganizationByID)
	// r.Put("/organizations/{id}", controllers.UpdateOrganization)
	// r.Delete("/organizations/{id}", controllers.DeleteOrganization)

	return r
}
