package main

import (
	"fmt"
	"net/http"
	"strava-weather-integration/router"
	"strava-weather-integration/utils"
)

func main() {
	port := utils.GetPortToUse()
	r := router.GetRouter()
	fmt.Println(fmt.Sprintf("Starting server at port %d", port))
	serverAddress := fmt.Sprintf("localhost:%d", port)
	http.ListenAndServe(serverAddress, r)
}
