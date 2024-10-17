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

//func New(
//	// +optional
//	source *dagger.Directory,
//	// +optional
//	awsAccessKey *dagger.Secret,
//	// +optional
//	awsSecretKey *dagger.Secret,
//	// +optional
//	awsSessionToken *dagger.Secret,
//) *Terraform {
//	return &Terraform{
//		Src: source,
//	}
//}

type Terraform struct {
	// Src *dagger.Directory
	// AwsAccessKey, AwsSecretKey, AwsSessionToken *dagger.Secret
}

func (t *Terraform) BuildEnv(src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("hashicorp/terraform:latest").
		WithDirectory("/src", src).
		// Terminal(). This allows to debug step
		WithWorkdir("/src")
}

func (t *Terraform) Format(ctx context.Context, src *dagger.Directory) (string, error) {
	return t.BuildEnv(src).
		WithExec([]string{"terraform", "fmt", "-check"}).
		Stdout(ctx)
}

func runCommand(container *dagger.Container, command []string) *dagger.Container {
	return container.WithExec(command)
}

func (t *Terraform) Plan(ctx context.Context,
	src *dagger.Directory,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	awsSessionToken *dagger.Secret,
) (string, error) {
	init := t.BuildEnv(src).
		WithExec([]string{"terraform", "init", "-reconfigure"})

	return init.
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
		WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken).
		WithExec([]string{"terraform", "plan"}).
		Stdout(ctx)
}

//func (t *Terraform) Validate(
//  src *dagger.Directory,
//  awsAccessKey *dagger.Secret,
//  awsSecretKey *dagger.Secret,
//  awsSessionToken *dagger.Secret,
//) error {
//
//  ctx := context.Background()
//  container := t.BuildEnv(src)
//
//
//  if awsSessionToken != dagger.Secret{} {
//    return container.
//      WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
//      WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
//      WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken).
//      WithExec([]string{"terraform", "validate"}).
//      Stdout(ctx)
//    }
//
//  return container.
//    WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
//    WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
//    WithExec([]string{"terraform", "validate"}).
//    Stdout(ctx)
//}
