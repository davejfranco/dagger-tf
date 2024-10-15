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

type Terraform struct{}

// TODO: Fix error: STS: Get CallerIdentity exceeded maximum number of attempts 9
//func (t *Terraform) Localstack(ctx context.Context, src *dagger.Directory) (string, error) {
//	init, err := t.init(ctx, src)
//	if err != nil {
//		return "", err
//	}
//
//	localstack := dag.Container().
//		From("localstack/localstack").
//		WithExposedPort(4566).
//		AsService()
//
//	lsSrv, err := localstack.Start(ctx)
//	if err != nil {
//		return "", err
//	}
//
//	defer localstack.Stop(ctx)
//
//	return init.
//		WithServiceBinding("localstack", lsSrv).
//		WithExec([]string{"terraform", "plan"}).
//		Stdout(ctx)
//}

func (t *Terraform) Apply(ctx context.Context,
	src *dagger.Directory,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
) (string, error) {
	init, err := t.init(
		ctx,
		src,
		awsAccessKey,
		awsSecretKey,
	)
	if err != nil {
		return "", err
	}

	return init.
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
		WithExec([]string{"terraform", "apply", "-auto-approve"}).
		Stdout(ctx)
}

func (t *Terraform) Plan(ctx context.Context,
	src *dagger.Directory,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	awsSessionToken *dagger.Secret,
) (string, error) {
	init, err := t.init(
		ctx,
		src,
		awsAccessKey,
		awsSecretKey,
		awsSessionToken,
	)
	if err != nil {
		return "", err
	}

	return init.
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
		WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken).
		WithExec([]string{"terraform", "plan"}).
		Stdout(ctx)
}

func (t *Terraform) Format(ctx context.Context, src *dagger.Directory) (string, error) {
	return t.BuildEnv(src).
		WithExec([]string{"terraform", "fmt", "-check"}).
		Stdout(ctx)
}

//func (t *Terraform) validate(ctx context.Context, src *dagger.Directory) (string, error) {
//	container, err := t.init(ctx, src)
//	if err != nil {
//		return "", err
//	}
//
//	return container.
//		WithExec([]string{"terraform", "validate"}).
//		Stdout(ctx)
//}

func (t *Terraform) init(ctx context.Context,
	src *dagger.Directory,
	awsAccessKey *dagger.Secret,
	awsSecretKey *dagger.Secret,
	awsSessionToken *dagger.Secret,
) (*dagger.Container, error) {
	container := t.BuildEnv(src).
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKey).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretKey).
		WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken).
		WithExec([]string{"terraform", "init", "-reconfigure"})

	_, err := container.Stdout(ctx)
	if err != nil {
		return nil, err
	}
	return container, nil
}

func (t *Terraform) BuildEnv(src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("hashicorp/terraform:latest").
		WithDirectory("/src", src).
		// Terminal(). This allows to debug step
		WithWorkdir("/src")
}
