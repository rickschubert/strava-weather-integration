package handlers

import (
	"net/http"
	"net/url"
	"os"
	"strava-weather-integration/utils/response"
)

func OAuth2Authenticate(w http.ResponseWriter, r *http.Request) {
	redirectUrl, err := url.Parse("https://www.strava.com/oauth/authorize")
	if err != nil {
		response.ForwardError(err)
	}
	queries := redirectUrl.Query()
	queries.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	queries.Set("response_type", "code")
	// TODO: Port and base URL should be tied to actual port and base URL, not just hard-coded
	queries.Set("redirect_uri", "http://localhost:3000/auth/redirect")
	queries.Set("approval_prompt", "force")
	queries.Set("scope", "read,activity:read_all")
	redirectUrl.RawQuery = queries.Encode()
	http.Redirect(w, r, redirectUrl.String(), 302)
}
