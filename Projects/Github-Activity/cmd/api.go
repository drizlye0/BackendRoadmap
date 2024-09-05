package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"unicode"
	"unicode/utf8"

	"github.com/drizlye0/GithubAPI-CLI/types"
)

func RequestAPI(endpoint string) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept":               {"application/vnd.github+json"},
		"X-GitHub-Api-Version": {"2022-11-28"},
		"X-Poll-Interval":      {"30"},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("User not found")
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	repo := []types.Repository{}
	err = json.Unmarshal(body, &repo)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range repo {
		message, err := getStringEvent(v)
		if err != nil {
			return err
		}
		fmt.Println(message)
	}
	return nil
}

func getStringEvent(repo types.Repository) (message string, err error) {
	EventTypes := map[string]string{
		"CommitCommentEvent":            fmt.Sprintf("%s commit comment in %s", repo.Payload.Action, repo.Repo.Name),
		"CreateEvent":                   fmt.Sprintf("created a git %s in %s", repo.Payload.Ref_Type, repo.Repo.Name),
		"DeleteEvent":                   fmt.Sprintf("deleted a git %s in %s", repo.Payload.Ref_Type, repo.Repo.Name),
		"ForkEvent":                     fmt.Sprintf("created a for of %s", repo.Repo.Name),
		"GollumEvent":                   fmt.Sprintf("create or update a wiki page in %s", repo.Repo.Name),
		"IssueCommentEvent":             fmt.Sprintf("%s issue comment in %s", repo.Payload.Action, repo.Repo.Name),
		"IssuesEvent":                   fmt.Sprintf("%s a new issue in %s", repo.Payload.Action, repo.Repo.Name),
		"MemberEvent":                   fmt.Sprintf("%s a colaborator in %s", repo.Payload.Action, repo.Repo.Name),
		"PublicEvent":                   fmt.Sprintf("he made public a private repository"),
		"PullRequestEvent":              fmt.Sprintf("%s a pull request in %s", repo.Payload.Action, repo.Repo.Name),
		"PullRequestReviewEvent":        fmt.Sprintf("%s a pull request review in %s", repo.Payload.Action, repo.Repo.Name),
		"PullRequestReviewCommentEvent": fmt.Sprintf("%s comment in pull request review in %s", repo.Payload.Action, repo.Repo.Name),
		"PullRequestReviewThreadEvent":  fmt.Sprintf("%s pull request in %s", repo.Payload.Action, repo.Repo.Name),
		"PushEvent":                     fmt.Sprintf("pushed %d commits to %s", repo.Payload.Size, repo.Repo.Name),
		"ReleaseEvent":                  fmt.Sprintf("%s release in %s", repo.Payload.Action, repo.Repo.Name),
		"SponsorshipEvent":              fmt.Sprintf("%s sponsor on list in %s", repo.Payload.Action, repo.Repo.Name),
		"WatchEvent":                    fmt.Sprintf("starred %s", repo.Repo.Name),
	}
	s := EventTypes[repo.Type]
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return "", errors.New("decode error")
	}
	message = string(unicode.ToUpper(r)) + s[size:]
	return message, nil
}
