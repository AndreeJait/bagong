package helper

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func CloneGit(repo, destination string) (repository *git.Repository, err error) {
	username, err := GoDotEnvVariable("GIT_USERNAME")
	if err != nil {
		return nil, err
	}
	password, err := GoDotEnvVariable("GIT_TOKEN")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s:%s@%s", username, password, repo)
	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	r, err := git.PlainClone(destination, false, &options)

	if err != nil {
		return nil, err
	}
	return r, err
}
