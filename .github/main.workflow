workflow "Deploy" {
  on = "push"
  resolves = ["hharnisc/terraform-github-actions-apply"]
}

action "hharnisc/terraform-github-actions-apply" {
  uses = "hharnisc/terraform-github-actions-apply@fef8a3"
  secrets = ["GITHUB_TOKEN"]
}
