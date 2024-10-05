module "gha-access" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-github-oidc-role"
  version = "5.44.1"

  name = "gha-iam-role"

  // Limit access to the role from actions deployed on the ono-platform only
  subjects = [
    "repo:davejfranco/dagger-tf:pull_request",
    "repo:davejfranco/dagger-tf:ref:refs/heads/main"

  ]

  policies = {
    account-full-access = "arn:aws:iam::aws:policy/AdministratorAccess"
  }

}
