package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marcusziade/github-api/endpoints"
)

const GitHubTokenEnv = "GITHUB_TOKEN"
const DefaultPages = 2
const MaxPages = 5

func main() {
	token := os.Getenv(GitHubTokenEnv)
	if token == "" {
		log.Fatal("GitHub token missing.")
	}

	e := endpoints.NewEndpoints(&http.Client{})

	http.HandleFunc("/github_user/", getUserHandler(e, token))
	http.HandleFunc("/user", getAuthUserHandler(e, token))
	http.HandleFunc("/github_starred/", getStarredHandler(e, token))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserHandler(e *endpoints.Endpoints, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Path[len("/github_user/"):]
		user, err := e.GetUser(username, token)
		handleRequest(w, user, err)
	}
}

func getAuthUserHandler(e *endpoints.Endpoints, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := e.GetAuthenticatedUser(token)
		handleRequest(w, user, err)
	}
}

func getStarredHandler(e *endpoints.Endpoints, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages := getPageCount(r)
		username := r.URL.Path[len("/github_starred/"):]
		starredRepos, err := e.GetStarredRepos(username, token, pages)
		handleRequest(w, starredRepos, err)
	}
}

func getPageCount(r *http.Request) int {
	pagesStr := r.URL.Query().Get("pages")
	pages, err := strconv.Atoi(pagesStr)
	if err != nil || pages <= 0 {
		return DefaultPages
	}
	if pages > MaxPages {
		return MaxPages
	}
	return pages
}

func handleRequest(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON)
}
