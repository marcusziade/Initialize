// starred_repos.go
package endpoints

import (
	"fmt"
	"net/http"
)

type StarredRepo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GetStarredRepos returns a list of starred repos for the currently authenticated username.
func GetStarredRepos(username string, token string) ([]StarredRepo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/starred", username), nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var starredRepos []StarredRepo
	if err := handleResponse(resp, &starredRepos); err != nil {
		return nil, err
	}

	return starredRepos, nil
}
