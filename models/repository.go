package models

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
