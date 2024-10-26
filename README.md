# Dagger Demo

This is a demonstration of how to use Dagger for deploying Terraform and integrating it with GitHub actions.

## Pre-requisites

- go > 1.22
- dagger >= 0.13.5

## Useful dagger commands

### List functions

```shell
dagger functions
```

### Call functions

Verify that the Terraform call is correctly formatted.

```shell
dagger call fmt-check
```

Terraform plan

```shell
dagger call plan --aws-access-key=env:AWS_ACCESS_KEY_ID \
    --aws-secret-key=env:AWS_SECRET_ACCESS_KEY
```

Terraform apply

```shell
dagger call apply --aws-access-key=env:AWS_ACCESS_KEY_ID \
    --aws-secret-key=env:AWS_SECRET_ACCESS_KEY
```

## Github actions

A key benefit of using GitHub Actions is that it removes the need for provisioning CI/CD.
However, the development process can be challenging.

Since we're using Terraform with AWS notice, that the functions are requiring
aws credentials

```Go
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
```

To make it more secure we have a IAM Role with OIDC that trust our github repo

```hcl
module "gha-access" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-github-oidc-role"
  version = "5.44.1"

  name = "gha-iam-role"

  // Limit access to the role from actions deployed on the ono-platform only
  subjects = [
    "repo:davejfranco/dagger-tf:pull_request",
    "repo:davejfranco/dagger-tf:ref:refs/heads/*"

  ]

  policies = {
    account-full-access = "arn:aws:iam::aws:policy/AdministratorAccess"
  }

}
```

and finally our gha workflow get credentials and we use these to pass it to the function

```yaml
- name: AWS Login
  id: creds
  uses: aws-actions/configure-aws-credentials@v4
  with:
    aws-region: us-east-1
    role-to-assume: arn:aws:iam::444106639146:role/gha-iam-role
    role-session-name: ghaSession
    output-credentials: true

- name: tf plan
  uses: dagger/dagger-for-github@v6
  with:
    verb: call
    args: plan
      --aws-access-key=AWS_ACCESS_KEY_ID
      --aws-secret-key=AWS_SECRET_ACCESS_KEY
      --aws-session-token=AWS_SESSION_TOKEN
      --github-ref=$GITHUB_REF
      --github-token=${{ secrets.GITHUB_TOKEN }}
      --github-repository=$GITHUB_REPOSITORY
    cloud-token: ${{ secrets.DAGGER_CLOUD_TOKEN }}
  env:
    AWS_ACCESS_KEY_ID: ${{ steps.creds.outputs.aws-access-key-id }}
    AWS_SECRET_ACCESS_KEY: ${{ steps.creds.outputs.aws-secret-access-key }}
    AWS_SESSION_TOKEN: ${{ steps.creds.outputs.aws-session-token }}
    GITHUB_REF: ${{ env.GITHUB_REF }}
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    GITHUB_REPOSITORY: ${{ env.GITHUB_REPOSITORY }}
```
