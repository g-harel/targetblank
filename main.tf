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

module "targetblank" {
  source = "./terraform"
}
