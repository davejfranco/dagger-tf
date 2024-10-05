locals {
  stack_name = "gha-stack"
  network = {
    cidr = "10.0.0.0/16"
  }

}
### Datasources
data "aws_caller_identity" "current" {}

data "aws_availability_zones" "this" {
  state = "available"
}

#module "vpc" {
#  source  = "terraform-aws-modules/vpc/aws"
#  version = "5.13.0"
#
#  name = "${local.stack_name}-vpc" // local.stack_name
#  cidr = local.network.cidr
#
#  azs             = slice(data.aws_availability_zones.this.names, 0, 2) // data.aws_availability_zones.this.names 
#  private_subnets = [for i in range(3) : cidrsubnet(local.network.cidr, 8, i)]
#  public_subnets  = [for i in range(3, 6) : cidrsubnet(local.network.cidr, 8, i)]
#
#  enable_nat_gateway = true
#  single_nat_gateway = true
#  enable_vpn_gateway = false
#
#  tags = {
#    Name = "${local.stack_name}-vpc"
#  }
#}





















