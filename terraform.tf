terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.67.0"
    }
    #helm = {
    #  source  = "hashicorp/helm"
    #  version = "2.15.0"
    #}
    #kubectl = {
    #  source  = "alekc/kubectl"
    #  version = "~> 2.0"
    #}
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

#provider "helm" {
#  kubernetes {
#    host                   = module.eks.cluster_endpoint
#    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
#    token                  = data.aws_eks_cluster_auth.cluster.token
#
#    exec {
#      api_version = "client.authentication.k8s.io/v1beta1"
#      args        = ["eks", "get-token", "--cluster-name", module.eks.cluster_name]
#      command     = "aws"
#    }
#  }
#}
#
#
#data "aws_eks_cluster_auth" "cluster" {
#  name = "${local.stack_name}-cluster"
#}
#
#provider "kubectl" {
#  host                   = module.eks.cluster_endpoint
#  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
#  token                  = data.aws_eks_cluster_auth.cluster.token
#}
