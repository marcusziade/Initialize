package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PullRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}

func main() {
	http.HandleFunc("/pullrequests", pullRequestHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchPullRequests(token string) ([]PullRequest, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user/pulls", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pullRequests []PullRequest
	if err := json.Unmarshal(body, &pullRequests); err != nil {
		return nil, err
	}

	return pullRequests, nil
}

func pullRequestHandler(w http.ResponseWriter, r *http.Request) {
	token := "GITHUB_ACCESS_TOKEN"
	pullRequests, err := fetchPullRequests(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pullRequests)
}
