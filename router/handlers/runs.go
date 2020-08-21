package handlers

import (
	"fmt"
	"net/http"
)

func GetRuns(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here will come all the runs you did")
}
