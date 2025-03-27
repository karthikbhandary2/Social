package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("conflict error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusConflict, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusNotFound, err.Error())
}
