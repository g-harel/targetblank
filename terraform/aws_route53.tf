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
    zone_id                = "${aws_api_gateway_domain_name.domain_name.cloudfront_zone_id}"
    name                   = "${aws_api_gateway_domain_name.domain_name.cloudfront_domain_name}"
    evaluate_target_health = false
  }
}

resource "aws_acm_certificate" "cert" {
  domain_name               = "${local.domain_name}"
  validation_method         = "DNS"
  subject_alternative_names = ["*.${local.domain_name}"]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "cert_validation" {
  name            = "${element(tolist(aws_acm_certificate.cert.domain_validation_options), 0).resource_record_name}"
  type            = "${element(tolist(aws_acm_certificate.cert.domain_validation_options), 0).resource_record_type}"
  zone_id         = "${aws_route53_zone.primary.id}"
  records         = ["${element(tolist(aws_acm_certificate.cert.domain_validation_options), 0).resource_record_value}"]
  ttl             = 60
  allow_overwrite = true
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = "${aws_acm_certificate.cert.arn}"
  validation_record_fqdns = ["${aws_route53_record.cert_validation.fqdn}"]
}
