package main

import (
	"fmt"
	"net/http"
)

// logError is a generic logger
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

//errorResponse sends json formatted error responses
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request,
	status int, message interface{}) {

	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

//serverErrorResponse used for unexpected server errors
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request,
	err error) {

	app.logError(r, err)
	message := "server encountered a problem"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {

	message := "resource cannot be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {

	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
