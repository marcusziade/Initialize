// Description: This file contains the functions that interact with the GitHub API to get user information.
package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GitHubUser struct {
	Login                   string `json:"login"`
	ID                      int    `json:"id"`
	NodeID                  string `json:"node_id"`
	AvatarURL               string `json:"avatar_url"`
	GravatarID              string `json:"gravatar_id"`
	URL                     string `json:"url"`
	HTMLURL                 string `json:"html_url"`
	FollowersURL            string `json:"followers_url"`
	FollowingURL            string `json:"following_url"`
	GistsURL                string `json:"gists_url"`
	StarredURL              string `json:"starred_url"`
	SubscriptionsURL        string `json:"subscriptions_url"`
	OrganizationsURL        string `json:"organizations_url"`
	ReposURL                string `json:"repos_url"`
	EventsURL               string `json:"events_url"`
	ReceivedEventsURL       string `json:"received_events_url"`
	Type                    string `json:"type"`
	SiteAdmin               bool   `json:"site_admin"`
	Name                    string `json:"name"`
	Company                 string `json:"company"`
	Blog                    string `json:"blog"`
	Location                string `json:"location"`
	Email                   string `json:"email"`
	Hireable                bool   `json:"hireable"`
	Bio                     string `json:"bio"`
	TwitterUsername         string `json:"twitter_username"`
	PublicRepos             int    `json:"public_repos"`
	PublicGists             int    `json:"public_gists"`
	Followers               int    `json:"followers"`
	Following               int    `json:"following"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	PrivateGists            int    `json:"private_gists"`
	TotalPrivateRepos       int    `json:"total_private_repos"`
	OwnedPrivateRepos       int    `json:"owned_private_repos"`
	DiskUsage               int    `json:"disk_usage"`
	Collaborators           int    `json:"collaborators"`
	TwoFactorAuthentication bool   `json:"two_factor_authentication"`
	Plan                    Plan   `json:"plan"`
}

type Plan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	PrivateRepos  int    `json:"private_repos"`
	Collaborators int    `json:"collaborators"`
}

var httpClient = &http.Client{}

/*
If the authenticated user is authenticated with an OAuth token with the user scope, then the response lists public and private profile information. If the authenticated user is authenticated through OAuth without the user scope, then the response lists only public profile information.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
*/
func GetAuthenticatedUser(token string) (*GitHubUser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	if err := handleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

/*
If your email is set to private and you send an email parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#update-the-authenticated-user
*/
func UpdateAuthenticatedUser(token string, updatedFields map[string]interface{}) (*GitHubUser, error) {
	req, err := http.NewRequest("PATCH", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, token)

	jsonData, err := json.Marshal(updatedFields)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	if err := handleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

/*
Provides publicly available information about someone with a GitHub account.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user
*/
func GetUser(username string, token string) (*GitHubUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	if err := handleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func PrettyPrintedJSON(data interface{}) string {
	indentedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate JSON: ", err)
	}
	return string(indentedJSON)
}

func setHeaders(req *http.Request, token string) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
}

func unmarshalBody(body io.ReadCloser, obj interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, obj)
}

func handleResponse(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	return unmarshalBody(resp.Body, obj)
}
