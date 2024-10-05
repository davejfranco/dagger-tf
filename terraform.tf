terraform {
  backend "s3" {
    bucket = "444106639146-tf-state"
    key    = "dagger"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.67.0"
    }
  }
}

provider "aws" {
  alias                       = "locastack"
  access_key                  = "test"
  secret_key                  = "test"
  region                      = "us-east-1"
  s3_use_path_style           = false
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    ec2         = "http://localhost:4566"
    elasticache = "http://localhost:4566"
    iam         = "http://localhost:4566"
    redshift    = "http://localhost:4566"
    route53     = "http://localhost:4566"
    sts         = "http://localhost:4566"
  }
}

provider "aws" {
  region = "us-east-1"

  default_tags {
    tags = {
      Terraform   = "true"
      Environment = "development"
    }
  }
}

