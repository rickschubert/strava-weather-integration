package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strava-weather-integration/utils/response"
	"time"
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

type AthleteActivity struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string      `json:"name"`
	Distance           float64     `json:"distance"`
	MovingTime         int         `json:"moving_time"`
	ElapsedTime        int         `json:"elapsed_time"`
	TotalElevationGain float64     `json:"total_elevation_gain"`
	Type               string      `json:"type"`
	WorkoutType        int         `json:"workout_type"`
	ID                 int64       `json:"id"`
	ExternalID         string      `json:"external_id"`
	UploadID           int64       `json:"upload_id"`
	StartDate          time.Time   `json:"start_date"`
	StartDateLocal     time.Time   `json:"start_date_local"`
	Timezone           string      `json:"timezone"`
	UtcOffset          float64     `json:"utc_offset"`
	StartLatlng        []float64   `json:"start_latlng"`
	EndLatlng          []float64   `json:"end_latlng"`
	LocationCity       interface{} `json:"location_city"`
	LocationState      interface{} `json:"location_state"`
	LocationCountry    interface{} `json:"location_country"`
	StartLatitude      float64     `json:"start_latitude"`
	StartLongitude     float64     `json:"start_longitude"`
	AchievementCount   int         `json:"achievement_count"`
	KudosCount         int         `json:"kudos_count"`
	CommentCount       int         `json:"comment_count"`
	AthleteCount       int         `json:"athlete_count"`
	PhotoCount         int         `json:"photo_count"`
	Map                struct {
		ID              string `json:"id"`
		SummaryPolyline string `json:"summary_polyline"`
		ResourceState   int    `json:"resource_state"`
	} `json:"map"`
	Trainer                    bool        `json:"trainer"`
	Commute                    bool        `json:"commute"`
	Manual                     bool        `json:"manual"`
	Private                    bool        `json:"private"`
	Visibility                 string      `json:"visibility"`
	Flagged                    bool        `json:"flagged"`
	GearID                     interface{} `json:"gear_id"`
	FromAcceptedTag            bool        `json:"from_accepted_tag"`
	UploadIDStr                string      `json:"upload_id_str"`
	AverageSpeed               float64     `json:"average_speed"`
	MaxSpeed                   float64     `json:"max_speed"`
	HasHeartrate               bool        `json:"has_heartrate"`
	HeartrateOptOut            bool        `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool        `json:"display_hide_heartrate_option"`
	ElevHigh                   float64     `json:"elev_high"`
	ElevLow                    float64     `json:"elev_low"`
	PrCount                    int         `json:"pr_count"`
	TotalPhotoCount            int         `json:"total_photo_count"`
	HasKudoed                  bool        `json:"has_kudoed"`
}

// Returns the access token retrieved from the strava API after having authenticated
// successfully. Should the authentication have been unsuccessful, there is also
// a network response returned which can be displayed to the user - in that case
// the authentication token would be empty.
func retrieveAccessToken(authorizationCode string) (string, response.Response) {
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
		return "", response.CustomResponse(err.Error(), resp.StatusCode)
	}
	var accessTokenResponse = new(StravaOauthTokenResponse)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", response.ForwardError(err)
	}
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", response.ForwardError(err)
	}
	return accessTokenResponse.AccessToken, response.Response{}
}

func retrieveRuns(accessToken string) response.Response {
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities", nil)
	if err != nil {
		return response.ForwardError(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return response.ForwardError(reqErr)
	}

	var activities []AthleteActivity
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.ForwardError(err)
	}
	err = json.Unmarshal(body, &activities)
	if err != nil {
		return response.ForwardError(err)
	}
	return response.SuccessResponse(activities[0].Name)
}

func OAuth2Redirect(w http.ResponseWriter, r *http.Request) {
	receivedQueryParams := r.URL.Query()
	authCode := receivedQueryParams.Get("code")
	if authCode == "" {
		response.WriteResponse(w, response.InternalServerError("The API didn't return an authorization code, sorry"))
	} else {
		accessToken, errorResponse := retrieveAccessToken(authCode)
		if accessToken == "" {
			response.WriteResponse(w, errorResponse)
		} else {
			runsResponse := retrieveRuns(accessToken)
			response.WriteResponse(w, runsResponse)
		}
	}
}
