package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/marcusziade/github-api/endpoints"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GitHub token must be set in environment variable GITHUB_TOKEN")
	}

	r := gin.Default()

	r.GET("/github_user/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, err := endpoints.GetUser(username, token)
		handleUserRequest(c, user, err)
	})

	r.GET("/user", func(c *gin.Context) {
		user, err := endpoints.GetAuthenticatedUser(token)
		handleUserRequest(c, user, err)
	})

	r.PATCH("/user", func(c *gin.Context) {
		var updatedFields map[string]interface{}
		if err := c.ShouldBindJSON(&updatedFields); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := endpoints.UpdateAuthenticatedUser(token, updatedFields)
		handleUserRequest(c, user, err)
	})

	r.GET("/github_user/:username/starred", func(c *gin.Context) {
		username := c.Param("username")
		starredRepos, err := endpoints.GetStarredRepos(username, token)
		handleUserRequest(c, starredRepos, err)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %s", err.Error())
	}
}

func handleUserRequest(c *gin.Context, user interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Disabled during debugging.
	// c.JSON(http.StatusOK, user)

	// Pretty-printed JSON during debugging.
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, endpoints.PrettyPrintedJSON(user))
}
