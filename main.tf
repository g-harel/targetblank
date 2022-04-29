terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-u2j51"
    key     = "targetblank.tfstate"
    region  = "us-east-1"
  }
  required_providers {
    archive = {
      source  = "hashicorp/archive"
      version = "2.2.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "4.10.0"
    }
    external = {
      source  = "hashicorp/external"
      version = "2.2.2"
    }
    local = {
      source  = "hashicorp/local"
      version = "2.2.2"
    }
  }
}

provider "archive" {}

provider "aws" {
  region = "us-east-1"
}

module "targetblank" {
  source = "./terraform"
}
