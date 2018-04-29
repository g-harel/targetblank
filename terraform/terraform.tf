terraform {
  backend "s3" {
    encrypt = true
    bucket  = "terraform-state-targetblank"
    region  = "us-east-2"
    key     = "terraform.tfstate"
  }
}

data "terraform_remote_state" "terraform_state_targetblank" {
  backend = "s3"

  config {
    bucket = "${aws_s3_bucket.terraform_state_targetblank.bucket}"
    region = "us-east-2"
    key    = "terraform.tfstate"
  }
}

resource "aws_s3_bucket" "terraform_state_targetblank" {
  bucket = "terraform-state-targetblank"
}
