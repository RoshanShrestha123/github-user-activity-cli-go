package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Action  string `json:"action"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
		PullRequest struct {
			Url string `json:"url"`
		} `json:"pull_request"`

		Issue struct {
			Url string `json:"url"`
		} `json:"issue"`
	} `json:"payload"`
}

func main() {

	args := os.Args
	githubUsername := args[1:]

	if len(githubUsername) == 0 {
		log.Fatal("Must enter the github username")
	}

	username := githubUsername[0]

	fmt.Printf("**Searching result for the %s on github**\n", username)

	res, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		log.Fatal("Something went wrong")
	}

	defer res.Body.Close()

	var activities []Event
	if err := json.NewDecoder(res.Body).Decode(&activities); err != nil {
		log.Fatal("Something went wrong")
	}

	for _, value := range activities {

		switch value.Type {
		case "IssuesEvent":

			fmt.Printf("- %s the new Issue in %s\n", value.Payload.Action, value.Repo.Name)

		case "PublicEvent":
			fmt.Printf("- %s made %s Public 🔓\n", username, value.Repo.Name)

		case "WatchEvent":
			fmt.Printf("- %s %s\n", value.Payload.Action, value.Repo.Name)

		case "CreateEvent":
			fmt.Printf("- Created %s on %s\n", value.Payload.RefType, value.Repo.Name)

		case "PushEvent":
			fmt.Printf("- %s %d commits to %s\n", value.Payload.Action, len(value.Payload.Commits), value.Repo.Name)

		}

	}

}
