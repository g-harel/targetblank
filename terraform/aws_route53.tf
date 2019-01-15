resource "aws_acm_certificate" "ssl_cert" {
  domain_name               = "${local.domain_name}"
  validation_method         = "EMAIL"
  subject_alternative_names = ["*.${local.domain_name}"]
}

resource "aws_route53_zone" "primary" {
  name = "${local.domain_name}"
}

resource "aws_route53_record" "static_files" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "${local.domain_name}"
  type    = "A"

  alias {
    zone_id                = "${module.website.cloudfront_zone_id}"
    name                   = "${module.website.cloudfront_domain_name}"
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "api.${local.domain_name}"
  type    = "A"

  alias {
    zone_id                = "${module.api.cloudfront_zone_id}"
    name                   = "${module.api.cloudfront_domain_name}"
    evaluate_target_health = false
  }
}
