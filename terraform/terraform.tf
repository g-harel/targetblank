provider "aws" {
  region     = "${var.region}"
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
}

terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-u2j91"
    key     = "targetblank"           # cannot use interpolation
    region  = "us-east-2"
  }
}

data "terraform_remote_state" "terraform_state" {
  backend = "s3"

  config {
    bucket = "terraform-state-u2j91"
    key    = "terraform.tfstate"
    region = "us-east-2"
  }
}
