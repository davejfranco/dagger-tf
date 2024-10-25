// A generated module for Terraform functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v66/github"

	"dagger/terraform/internal/dagger"
)

func New(
	// +defaultPath="."
	source *dagger.Directory,
) *Terraform {
	return &Terraform{
		Src: source,
	}
}

type Terraform struct {
	Src *dagger.Directory
}

// Returns a build environment
func (t *Terraform) buildEnv() *dagger.Container {
	return dag.Container().
		From("hashicorp/terraform:latest").
		WithDirectory("/src", t.Src).
		WithWorkdir("/src")
}

// Checks if code is correctly formatted
func (t *Terraform) FmtCheck(ctx context.Context) (string, error) {
	return t.buildEnv().
		WithExec([]string{"terraform", "fmt", "-check"}).
		Stdout(ctx)
}

func (t *Terraform) init(
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	// +optional
	awsSessionToken *dagger.Secret,
) *dagger.Container {
	init := t.buildEnv().
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey)

	if awsSessionToken != nil {
		init = init.WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken)
	}

	return init.
		WithExec([]string{"terraform", "init", "-reconfigure"})
}

func prComment(ctx context.Context, token, owner, repo, pr, content string) error {
	body := fmt.Sprintf("```\n%s\n```", content)

	prInt, err := strconv.Atoi(pr)
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(token)
	_, _, err = client.Issues.CreateComment(ctx, owner, repo, prInt, &github.IssueComment{
		Body: &body,
	})

	return err
}

// Returns a Terraform plan
func (t *Terraform) Plan(ctx context.Context,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	// +optional
	awsSessionToken *dagger.Secret,
	// +optional
	// +default=""
	githubToken string,
	// +optional
	githubRepository string,
	// +optional
	githubRef string,
) (string, error) {
	container := t.init(
		awsAccessKey,
		awsSecretKey,
		awsSessionToken,
	)

	container = container.
		WithExec([]string{"terraform", "plan"})

	output, err := container.Stdout(ctx)
	if err != nil {
		return "", err
	}

	if githubToken != "" {
		repo := strings.Split(githubRepository, "/")
		pr := strings.Split(githubRef, "/")[2]
		err = prComment(ctx, githubToken, repo[0], repo[1], pr, output)
	}
	return output, err
}

// Executes Terraform Apply
func (t *Terraform) Apply(ctx context.Context,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	// +optional
	awsSessionToken *dagger.Secret,
) (string, error) {
	container := t.init(
		awsAccessKey,
		awsSecretKey,
		awsSessionToken,
	)

	container = container.
		WithExec([]string{"terraform", "apply", "-auto-approve"})

	return container.Stdout(ctx)
}
