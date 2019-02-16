workflow "Deploy" {
  on = "push"
  resolves = ["terraform-apply"]
}

action "npm-install" {
  uses = "actions/npm@59b64a598378f31e49cb76f27d6f3312b582f680"
  args = "install"
}

action "npm-build" {
  uses = "actions/npm@59b64a598378f31e49cb76f27d6f3312b582f680"
  needs = ["npm-install"]
  args = "run build"
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
  needs = ["npm-build", "terraform-init"]
  secrets = [
    "GITHUB_TOKEN",
    "AWS_ACCESS_KEY_ID",
    "AWS_SECRET_ACCESS_KEY",
  ]
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}
