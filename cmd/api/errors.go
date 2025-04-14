package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Warnw("forbidden", "method", r.Method,"path", r.URL.Path,"error")
	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) unauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusUnauthorized, err.Error())
}

func (app *application) unauthorizedBasicError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized basic error", "method", r.Method,"path", r.URL.Path,"error",err)
	w.Header().Set("WWW-Authenticate", `BASIC realm="restricted", charset="UTF-8"`)
	writeJSON(w, http.StatusUnauthorized, err.Error())
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("conflict error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusConflict, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found error", "method", r.Method,"path", r.URL.Path,"error",err)
	writeJSON(w, http.StatusNotFound, err.Error())
}

func (app *application) rateLimitExceeded(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}