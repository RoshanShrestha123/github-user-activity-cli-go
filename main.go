package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

func main() {
	fmt.Println("Enter the Username of the target github:")
	var username string
	fmt.Scan(&username)

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

	fmt.Println(activities)

}
