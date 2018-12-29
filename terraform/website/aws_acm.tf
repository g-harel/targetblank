resource "aws_acm_certificate" "ssl_cert" {
  domain_name               = "targetblank.org"
  validation_method         = "EMAIL"
  subject_alternative_names = ["*.targetblank.org"]
}
