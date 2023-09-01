package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jdkato/prose/v2"
)

func main() {
	title := os.Args[1]

	splitTitle := strings.SplitN(title, ": ", 2)

	if len(splitTitle) < 2 {
		fmt.Println("Incomplete PR title. Ensure you have both a prefix and a description.")
		os.Exit(1)
	}

	prefix, description := splitTitle[0], splitTitle[1]

	allowedPrefixes := []string{"refactor", "chore", "feat", "fix", "regfix"}
	if !contains(allowedPrefixes, prefix) {
		fmt.Printf("Invalid prefix. Allowed prefixes are: %v\n", allowedPrefixes)
		os.Exit(1)
	}

	doc, _ := prose.NewDocument(description)
	tokens := doc.Tokens()

	if len(tokens) > 0 && tokens[0].Tag != "VB" {
		fmt.Println("The PR title description should be in the imperative mood (e.g., 'Add feature', 'Fix bug').")
		os.Exit(1)
	}
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
