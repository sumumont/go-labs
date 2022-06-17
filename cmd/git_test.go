package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	directory := "code"
	gitUrl := "git@gitlab.apulis.com.cn:app/iqi-adapter.git"
	privateKey := `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`
	publicKeys, err := ssh.NewPublicKeys("git", []byte(privateKey), "")
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth:     publicKeys,
		URL:      gitUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
}
