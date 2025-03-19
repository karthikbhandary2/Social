package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSON(w, http.StatusConflict, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSON(w, http.StatusNotFound, "not found")
}
