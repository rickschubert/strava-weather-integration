package handlers

import (
	"fmt"
	"net/http"
)

func RouteNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route not found")
}
