package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Healthcheck
//	@Description	Returns the health status of the application
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Health status"
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
		"env": app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w,r, err)
	}
}