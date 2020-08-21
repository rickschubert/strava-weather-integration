package router

import (
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/rick", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Guten Tag")
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My second route")
	})

}
