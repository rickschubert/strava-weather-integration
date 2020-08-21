run:
	@go build
	@./strava-weather-integration

dev:
	gin run main.go
