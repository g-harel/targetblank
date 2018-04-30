provider "aws" {
  region     = "${var.region}"
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
}

terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-u2j91" # cannot use interpolation
    key     = "targetblank"           # cannot use interpolation
    region  = "us-east-2"             # cannot use interpolation
  }
}

data "terraform_remote_state" "tfstate" {
  backend = "s3"

  config {
    bucket = "terraform-state-u2j91"
    key    = "${var.name}"
    region = "${var.region}"
  }
}
