provider "aws" {
  region = "${var.region}"

  # credentials should be included as environment variables
  #   export AWS_ACCESS_KEY_ID="access_key"
  #   export AWS_SECRET_ACCESS_KEY="secret_key"
  #   export AWS_DEFAULT_REGION="us-east-1"
  #   export AWS_REGION="us-east-1"
}

terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-u2j51" # cannot use interpolation
    key     = "targetblank.tfstate"   # cannot use interpolation
    region  = "us-east-1"             # cannot use interpolation
  }
}

data "terraform_remote_state" "tfstate" {
  backend = "s3"

  config {
    bucket = "terraform-state-u2j51"
    key    = "${var.name}.tfstate"
    region = "${var.region}"
  }
}

module "website" {
  source = "./website"

  name   = "${var.name}"
  domain = "${var.domain}"
  region = "${var.region}"
}

module "api" {
  source = "./api"

  name   = "${var.name}"
  domain = "${var.domain}"
  region = "${var.region}"
}
