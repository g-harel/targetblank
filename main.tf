provider "aws" {
  region = "us-east-1"
}

terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-u2j51"
    key     = "targetblank.tfstate"
    region  = "us-east-1"
  }
}

data "terraform_remote_state" "tfstate" {
  backend = "s3"

  config {
    bucket = "terraform-state-u2j51"
    key    = "targetblank.tfstate"
    region = "us-east-1"
  }
}

module "app" {
  source = "./app/terraform"
}

module "api" {
  source = "./api/terraform"
}
