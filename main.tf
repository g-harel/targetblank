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

resource "aws_route53_zone" "primary" {
  name = "targetblank.org"
}

module "website" {
  source = "./terraform/website"

  primary_zone_id = "${aws_route53_zone.primary.zone_id}"
}

module "api" {
  source = "./terraform/functions"

  primary_zone_id = "${aws_route53_zone.primary.zone_id}"
}
