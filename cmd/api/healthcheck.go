package main

import (
	"fmt"
	"net/http"
)

// Implemented as a method on the Application struct. Makes it easier to inject dependancies
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
