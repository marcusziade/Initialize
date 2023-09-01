package endpoints

import (
	"fmt"
	"net/http"
	"sync"
)

type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}

type Repository struct {
	ID                       int         `json:"id"`
	NodeID                   string      `json:"node_id"`
	Name                     string      `json:"name"`
	FullName                 string      `json:"full_name"`
	Private                  bool        `json:"private"`
	Owner                    Owner       `json:"owner"`
	HTMLURL                  string      `json:"html_url"`
	Description              string      `json:"description"`
	Fork                     bool        `json:"fork"`
	URL                      string      `json:"url"`
	ForksURL                 string      `json:"forks_url"`
	KeysURL                  string      `json:"keys_url"`
	CollaboratorsURL         string      `json:"collaborators_url"`
	TeamsURL                 string      `json:"teams_url"`
	HooksURL                 string      `json:"hooks_url"`
	IssueEventsURL           string      `json:"issue_events_url"`
	EventsURL                string      `json:"events_url"`
	License                  License     `json:"license"`
	AllowForking             bool        `json:"allow_forking"`
	IsTemplate               bool        `json:"is_template"`
	WebCommitSignoffRequired bool        `json:"web_commit_signoff_required"`
	Topics                   []string    `json:"topics"`
	Visibility               string      `json:"visibility"`
	Forks                    int         `json:"forks"`
	OpenIssues               int         `json:"open_issues"`
	Watchers                 int         `json:"watchers"`
	DefaultBranch            string      `json:"default_branch"`
	Permissions              Permissions `json:"permissions"`
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
func (e *Endpoints) GetStarredRepos(username string, token string, pages int) ([]Repository, error) {
	var wg sync.WaitGroup
	ch := make(chan []Repository, pages)

	for i := 1; i <= pages; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			url := fmt.Sprintf("https://api.github.com/users/%s/starred?page=%d", username, page)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			setHeaders(req, token)
			resp, err := e.httpClient.Do(req)
			if err != nil {
				return
			}
			var starredRepos []Repository
			if err := unmarshalBody(resp.Body, &starredRepos); err != nil {
				return
			}
			ch <- starredRepos
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var allStarredRepos []Repository
	for repos := range ch {
		if repos != nil {
			allStarredRepos = append(allStarredRepos, repos...)
		}
	}

	return allStarredRepos, nil
}
