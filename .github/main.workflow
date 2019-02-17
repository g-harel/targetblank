workflow "Deploy" {
  on = "push"
  resolves = ["terraform apply"]
}

action "go mod download" {
  uses = "cedrickring/golang-action@1.1.0"
  args = "go mod download"
  env = {
    GO111MODULE = "on"
  }
}

action "go build authenticate" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/authenticate ./functions/authenticate"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build create" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/create ./functions/create"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build passwd" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/passwd ./functions/passwd"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build read" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/read ./functions/read"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build reset" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/reset ./functions/reset"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build update" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/update ./functions/update"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    GO111MODULE = "on"
  }
}

action "go build validate" {
  uses = "cedrickring/golang-action@1.1.0"
  needs = ["go mod download"]
  args = "go build -o .build/validate ./functions/validate"
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
    "go build create",
    "go build passwd",
    "go build read",
    "go build reset",
    "go build update",
    "go build validate",
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
