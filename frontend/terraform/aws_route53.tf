resource "aws_route53_zone" "primary" {
  name = "targetblank.org"
}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "targetblank.org"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.static_website.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.static_website.domain_name}"
    evaluate_target_health = false
  }
}
