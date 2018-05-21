variable primary_zone_id {}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${var.primary_zone_id}"
  name    = "targetblank.org"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.static_website.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.static_website.domain_name}"
    evaluate_target_health = false
  }
}
