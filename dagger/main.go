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

func (t *Terraform) BuildEnv() *dagger.Container {
	return dag.Container().
		From("hashicorp/terraform:latest").
		WithDirectory("/src", t.Src).
		WithWorkdir("/src")
}

func (t *Terraform) FmtCheck(ctx context.Context) (string, error) {
	return t.BuildEnv().
		WithExec([]string{"terraform", "fmt", "-check"}).
		Stdout(ctx)
}

func (t *Terraform) init(
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	// +optional
	awsSessionToken *dagger.Secret,
) *dagger.Container {
	init := t.BuildEnv().
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey)

	if awsSessionToken != nil {
		init = init.WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken)
	}

	return init.
		WithExec([]string{"terraform", "init", "-reconfigure"})
}

func (t *Terraform) Plan(ctx context.Context,
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
		WithExec([]string{"terraform", "plan"})

	return container.Stdout(ctx)
}

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
