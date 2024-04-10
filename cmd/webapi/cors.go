package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"x-example-header",
			"Content-Type",
			"Authorization",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(h)
}
