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
  for_each = {
    for dvo in aws_acm_certificate.example.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  name    = each.value.name
  type    = each.value.type
  zone_id = "${aws_route53_zone.primary.id}"
  records = [each.value.record]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = "${aws_acm_certificate.cert.arn}"
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}
