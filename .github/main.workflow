workflow "Deploy" {
  on = "push"
  resolves = ["terraform apply"]
}

action "go build authenticate" {
  uses = "cedrickring/golang-action@1.1.0"
  args = "go build -o .build/authenticate ./functions/authenticate"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "npm install" {
  uses = "actions/npm@59b64a5"
  args = "install"
}

action "npm build" {
  uses = "actions/npm@59b64a5"
  needs = ["npm install"]
  args = "run build"
}

action "terraform init" {
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

action "terraform apply" {
  uses = "g-harel/terraform-github-actions-apply@d49255c"
  needs = [
    "go build authenticate",
    "npm build",
    "terraform init",
  ]
  secrets = [
    "GITHUB_TOKEN",
    "AWS_ACCESS_KEY_ID",
    "AWS_SECRET_ACCESS_KEY",
  ]
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}
