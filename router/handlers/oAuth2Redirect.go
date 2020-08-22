package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strava-weather-integration/utils/response"
)

type StravaAthlete struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Put in the rest of the response here, not needed for now I find
}

type StravaOauthTokenResponse struct {
	TokenType    string        `json:"token_type"`
	ExpiresAt    int           `json:"expires_at"`
	ExpiresIn    int           `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	AccessToken  string        `json:"access_token"`
	Athlete      StravaAthlete `json:"athlete"`
}

func retrieveAccessToken(authorizationCode string) response.Response {
	stravaTokenUrl := "https://www.strava.com/oauth/token"
	authenticationDetails := url.Values{
		"code":          []string{authorizationCode},
		"client_secret": []string{os.Getenv("STRAVA_CLIENT_SECRET")},
		"client_id":     []string{os.Getenv("STRAVA_CLIENT_ID")},
		"grant_type":    []string{"authorization_code"},
	}
	resp, err := http.PostForm(
		stravaTokenUrl,
		authenticationDetails,
	)
	if err != nil {
		return response.CustomResponse(err.Error(), resp.StatusCode)
	}
	var accessTokenResponse = new(StravaOauthTokenResponse)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.ForwardError(err)
	}
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return response.ForwardError(err)
	}
	return response.SuccessResponse(fmt.Sprintf("Here goes your access token: %v", accessTokenResponse))
}

func OAuth2Redirect(w http.ResponseWriter, r *http.Request) {
	receivedQueryParams := r.URL.Query()
	authCode := receivedQueryParams.Get("code")
	var resp response.Response
	if authCode == "" {
		resp = response.InternalServerError("The API didn't return an authorization code, sorry")
	} else {
		resp = retrieveAccessToken(authCode)
	}
	response.WriteResponse(w, resp)
}
