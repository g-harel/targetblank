workflow "Deploy" {
  on = "push"
  resolves = ["terraform-apply"]
}

action "terraform-init" {
  uses = "hashicorp/terraform-github-actions/init@v0.1.2"
  secrets = [
    "GITHUB_TOKEN",
    "AWS_ACCESS_KEY_ID",
    "AWS_SECRET_ACCESS_KEY",
  ]
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}

action "terraform-apply" {
  uses = "g-harel/terraform-github-actions-apply@d49255c"
  needs = ["terraform-init"]
  secrets = [
    "GITHUB_TOKEN",
    "AWS_ACCESS_KEY_ID",
    "AWS_SECRET_ACCESS_KEY",
  ]
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}
