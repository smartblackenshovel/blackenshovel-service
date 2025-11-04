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
	r.Get("/organizations/{id}", controllers.GetOrganizationByID)
	r.Patch("/organizations/{id}", controllers.UpdateOrganization)
	r.Delete("/organizations/{id}", controllers.DeleteOrganization)

	// Shovel routes
	r.Get("/shovels", controllers.GetShovels)
	r.Post("/shovels", controllers.CreateShovel)
	r.Get("/shovels/{id}", controllers.GetShovelByID)
	r.Patch("/shovels/{id}", controllers.UpdateShovel)
	r.Delete("/shovels/{id}", controllers.DeleteShovel)

	// ShovelLog routes
	r.Get("/shovel_logs", controllers.ListShovelLogs)
	r.Post("/shovel_logs", controllers.CreateShovelLog)
	r.Get("/shovel_logs/{id}", controllers.GetShovelLogByID)
	r.Delete("/shovel_logs/{id}", controllers.DeleteShovelLog)

	return r
}
