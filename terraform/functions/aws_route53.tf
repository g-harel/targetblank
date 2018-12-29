variable primary_zone_id {}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${var.primary_zone_id}"
  name    = "api.targetblank.org"
  type    = "A"

  alias {
    name                   = "${aws_api_gateway_domain_name.domain_name.cloudfront_domain_name}"
    zone_id                = "${aws_api_gateway_domain_name.domain_name.cloudfront_zone_id}"
    evaluate_target_health = false
  }
}
