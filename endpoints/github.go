package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcusziade/github-api/models"
	"github.com/marcusziade/github-api/utils"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Endpoints struct {
	httpClient HttpClient
}

func NewEndpoints(client HttpClient) *Endpoints {
	return &Endpoints{httpClient: client}
}

// GetStarredRepos fetches the GitHub repositories starred by a given user.
// Parameters:
// - username: GitHub username as a string.
// - token: GitHub personal access token as a string.
// - pages: Number of pages to fetch from the API as an integer.
//
// Returns:
// - []Repository: A slice containing the repositories starred by the user.
// - error: An error object if an error occurs, otherwise nil.
func (e *Endpoints) GetStarredRepos(username, token string, pages int) ([]models.Repository, error) {
	var repos []models.Repository

	for i := 1; i <= pages; i++ {
		url := fmt.Sprintf("https://api.github.com/users/%s/starred?page=%d", username, i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		utils.SetHeaders(req, token)
		resp, err := e.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		var pageRepos []models.Repository
		if err := utils.HandleResponse(resp, &pageRepos); err != nil {
			return nil, err
		}

		repos = append(repos, pageRepos...)
	}

	return repos, nil
}

/*
If the authenticated user is authenticated with an OAuth token with the user scope, then the response lists public and private profile information. If the authenticated user is authenticated through OAuth without the user scope, then the response lists only public profile information.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
*/
func (e *Endpoints) GetAuthenticatedUser(token string) (*models.GitHubUser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	utils.SetHeaders(req, token)
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var user models.GitHubUser
	if err := utils.HandleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

/*
If your email is set to private and you send an email parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#update-the-authenticated-user
*/
func (e *Endpoints) UpdateAuthenticatedUser(token string, updatedFields map[string]interface{}) (*models.GitHubUser, error) {
	payload, err := json.Marshal(updatedFields)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", "https://api.github.com/user", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	utils.SetHeaders(req, token)
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var updatedUser models.GitHubUser
	if err := utils.HandleResponse(resp, &updatedUser); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

/*
Provides publicly available information about a GitHub account.
https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user
*/
func (e *Endpoints) GetUser(username string, token string) (*models.GitHubUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	if err != nil {
		return nil, err
	}

	utils.SetHeaders(req, token)

	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var user models.GitHubUser
	if err := utils.HandleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
