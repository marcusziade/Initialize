package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marcusziade/github-api/endpoints"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GitHub token must be set in environment variable GITHUB_TOKEN")
	}

	e := endpoints.NewEndpoints(&http.Client{})

	http.HandleFunc("/github_user/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Path[len("/github_user/"):]
		user, err := e.GetUser(username, token)
		handleRequest(w, user, err)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user, err := e.GetAuthenticatedUser(token)
		handleRequest(w, user, err)
	})

	http.HandleFunc("/github_starred/", func(w http.ResponseWriter, r *http.Request) {
		pagesStr := r.URL.Query().Get("pages")
		pages, err := strconv.Atoi(pagesStr)
		if err != nil || pages <= 0 {
			pages = 2
		}
		if pages > 5 {
			pages = 5
		}
		username := r.URL.Path[len("/github_starred/"):]
		starredRepos, err := e.GetStarredRepos(username, token, pages)
		handleRequest(w, starredRepos, err)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON)
}
