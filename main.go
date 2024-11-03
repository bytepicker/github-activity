package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Unmarshal event based on type and handle the specific payloads for each case.
func handleEvent(event BasicEventProperties, rawEvent json.RawMessage) {
	switch event.Type {
	case "PushEvent":
		var pushEvent PushEvent
		if err := json.Unmarshal(rawEvent, &pushEvent); err != nil {
			fmt.Println("Error unmarshaling PushEvent:", err)
			return
		}
		fmt.Printf("- PushEvent in repo %s by %s\n", pushEvent.Repo.Name, pushEvent.Actor.Login)
		for _, commit := range pushEvent.Payload.Commits {
			fmt.Printf("  Commit: %s by %s\n", commit.Message, commit.Author.Name)
		}

	case "WatchEvent":
		var watchEvent WatchEvent
		if err := json.Unmarshal(rawEvent, &watchEvent); err != nil {
			fmt.Println("Error unmarshaling WatchEvent:", err)
			return
		}
		fmt.Printf("- WatchEvent: %s on repo %s\n", watchEvent.Payload.Action, watchEvent.Repo.Name)

	case "ForkEvent":
		var forkEvent ForkEvent
		if err := json.Unmarshal(rawEvent, &forkEvent); err != nil {
			fmt.Println("Error unmarshaling ForkEvent:", err)
			return
		}
		fmt.Printf("- Forked repo %s to %s\n", forkEvent.Repo.Name, forkEvent.Payload.Forkee.FullName)

	case "IssueCommentEvent":
		var issueCommentEvent IssueCommentEvent
		if err := json.Unmarshal(rawEvent, &issueCommentEvent); err != nil {
			fmt.Println("Error unmarshaling IssueCommentEvent:", err)
			return
		}
		fmt.Printf("- IssueCommentEvent on issue #%d: %s\n", issueCommentEvent.Payload.Issue.Number, issueCommentEvent.Payload.Comment.Body)

	case "CreateEvent":
		var createEvent CreateEvent
		if err := json.Unmarshal(rawEvent, &createEvent); err != nil {
			fmt.Println("Error unmarshaling CreateEvent:", err)
			return
		}
		fmt.Printf("- CreateEvent %s on repo %s\n", createEvent.Payload.RefType, createEvent.BasicEventProperties.Repo.Name)

	case "PullRequestEvent":
		var pullRequestEvent PullRequestEvent
		if err := json.Unmarshal(rawEvent, &pullRequestEvent); err != nil {
			fmt.Println("Error unmarshaling PullRequestEvent:", err)
			return
		}
		fmt.Printf("- PullRequestEvent %s on repo %s\n", pullRequestEvent.Payload.Action, pullRequestEvent.BasicEventProperties.Repo.Name)

	case "DeleteEvent":
		var deleteEvent DeleteEvent
		if err := json.Unmarshal(rawEvent, &deleteEvent); err != nil {
			fmt.Println("Error unmarshaling DeleteEvent:", err)
			return
		}
		fmt.Printf("- DeleteEvent on repo %s\n", deleteEvent.Repo.Name)

	default:
		fmt.Printf("- Event type: %s (not handled yet)\n", event.Type)
	}
}

func fetchGithubApi(username string) {
	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var rawEvents []json.RawMessage
	if err := json.Unmarshal(body, &rawEvents); err != nil {
		fmt.Println("Cannot unmarshal JSON:", err)
		return
	}

	for _, rawEvent := range rawEvents {
		var event BasicEventProperties
		if err := json.Unmarshal(rawEvent, &event); err != nil {
			fmt.Println("Error unmarshaling BasicEventProperties:", err)
			continue
		}
		handleEvent(event, rawEvent)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: github-activity username")
		return
	}

	fetchGithubApi(os.Args[1])
}
