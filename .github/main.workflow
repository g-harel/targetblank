workflow "Deploy" {
  on = "push"
  resolves = ["terraform-apply"]
}

action "terraform-init" {
  uses = "hashicorp/terraform-github-actions/init@v0.1.2"
  secrets = ["GITHUB_TOKEN"]
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}

action "terraform-apply" {
  uses = "hharnisc/terraform-github-actions-apply@v0.0.3-beta-02"
  needs = ["terraform-init"]
  secrets = ["GITHUB_TOKEN"]  
  env = {
    TF_ACTION_WORKING_DIR = "."
  }
}
