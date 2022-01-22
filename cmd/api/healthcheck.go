package main

import (
	"net/http"
)

// Implemented as a method on the Application struct. Makes it easier to inject dependancies
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// Approach 1: Write the JSON as the HTTP response body.
	/*
		js := `{"status": "available", "environment": %q, "version": %q}`
		js = fmt.Sprintf(js, app.config.env, version)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(js))
	*/
	// Approach 2: Uses json.Marshal()
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered an unexpected error", http.StatusInternalServerError)
	}
}
