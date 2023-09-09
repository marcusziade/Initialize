package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	http.HandleFunc("/github_readme_images/", getReadmeImagesHandler(e, token))

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

func getReadmeImagesHandler(e *endpoints.Endpoints, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		if len(paths) < 4 {
			http.Error(w, "Invalid URL format. It should be /github_readme_images/{owner}/{repo}", http.StatusBadRequest)
			return
		}
		owner, repo := paths[2], paths[3]

		imageURLs, err := e.GetReadmeImages(owner, repo, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(imageURLs) == 0 {
			http.Error(w, "No images found in README", http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(imageURLs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
