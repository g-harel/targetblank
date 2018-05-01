resource "aws_acm_certificate" "ssl_cert" {
  domain_name               = "${var.domain}"
  validation_method         = "EMAIL"
  subject_alternative_names = ["*.${var.domain}"]
}
