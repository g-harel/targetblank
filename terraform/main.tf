locals {
  domain_name = "targetblank.org"
}

resource "aws_route53_zone" "primary" {
  name = "${local.domain_name}"
}

resource "aws_acm_certificate" "ssl_cert" {
  domain_name               = "${local.domain_name}"
  validation_method         = "EMAIL"
  subject_alternative_names = ["*.${local.domain_name}"]
}


module "website" {
  source = "./modules/public-bucket"

  zone_id = "${aws_route53_zone.primary.zone_id}"
  alias_name = "${local.domain_name}"
  cert_arn = "${aws_acm_certificate.ssl_cert.arn}"
  bucket_name = "targetblank-static-website"

  source_dir = ".build"
  root_file = "index.html"
  files = ["website.f69400ca.css", "website.f69400ca.js"]
}

module "api" {
  source = "./functions"

  primary_zone_id = "${aws_route53_zone.primary.zone_id}"
  role = "${aws_iam_role.lambda.arn}"
}
