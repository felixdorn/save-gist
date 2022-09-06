package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"os"
	"time"
)

var Client *github.Client

func main() {
	err := run(os.Args[1:])
	if err != nil {
		if usageError, ok := err.(UsageError); ok {
			fmt.Fprintln(os.Stderr, usageError)
		} else {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		}
		os.Exit(1)
	}
}

type UsageError struct {
	Err error
}

func (e UsageError) Error() string {
	return "usage: " + e.Err.Error()
}

func run(args []string) error {
	if len(args) == 0 {
		return UsageError{Err: errors.New("usage: sg [FILES...]")}
	}

	files := map[github.GistFilename]github.GistFile{}
	for _, file := range args {
		bytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		content := string(bytes)
		files[github.GistFilename(file)] = github.GistFile{
			Content: &content,
		}
	}

	description := "generated by sg (" + time.Now().Format(time.RFC3339) + ")"
	public := false
	gist, _, err := Client.Gists.Create(context.Background(), &github.Gist{
		Description: &description,
		Public:      &public,
		Files:       files,
	})
	if err != nil {
		return err
	}

	fmt.Println(*gist.HTMLURL)

	return nil
}

func init() {
	githubToken := os.Getenv("SG_GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Fprintf(os.Stderr, "error: missing SG_GITHUB_TOKEN\n")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	Client = github.NewClient(tc)
}
