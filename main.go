package main

import (
	"fmt"
	"log"
	"net/http"
	"strava-weather-integration/router"
	"strava-weather-integration/utils"
)

func main() {
	port := utils.GetPortToUse()

	router.SetupRoutes()
	fmt.Println(fmt.Sprintf("Starting server at port %d", port))
	serverAddress := fmt.Sprintf("localhost:%d", port)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
