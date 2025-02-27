# GitHub Events Fetcher

A simple Go program that fetches recent public events for a given GitHub user and displays them in a readable format.

## Roadmap Project Github Activity Assignment
https://roadmap.sh/projects/github-user-activity

## Features
- Fetches GitHub events using the GitHub API.
- Displays event types such as Issues, Repository Visibility, Watches, Creations, and Pushes.
- Parses JSON response and structures output for easy readability.

## Installation
Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

Clone the repository:
```sh
git clone https://github.com/RoshanShrestha123/github-user-activity-cli-go
cd github-user-activity-cli-go
```

## Usage
Run the program with a GitHub username as an argument:
```sh
go run main.go <github-username>
```
Example:
```sh
go run main.go torvalds
```

## Output Example
```
**Searching result for torvalds on github**
- Created repository on linux
- torvalds made linux Public 🔓
- Pushed 3 commits to linux
- opened a new Issue in linux
```

## Dependencies
- `encoding/json`
- `fmt`
- `log`
- `net/http`
- `os`

## Notes
- The script fetches data from `https://api.github.com/users/{username}/events`.
- Make sure you have an active internet connection.
- GitHub API rate limits might apply for excessive requests.


