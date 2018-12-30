resource "aws_route53_zone" "primary" {
  name = "targetblank.org"
}

module "website" {
  source = "./website"

  primary_zone_id = "${aws_route53_zone.primary.zone_id}"
}

module "api" {
  source = "./functions"

  primary_zone_id = "${aws_route53_zone.primary.zone_id}"
}
