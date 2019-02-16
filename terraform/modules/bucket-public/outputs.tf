output "cloudfront_zone_id" {
  description = "The Cloudfront distribution's zone ID."
  value       = "${aws_cloudfront_distribution.public_bucket.hosted_zone_id}"
}

output "cloudfront_domain_name" {
  description = "The Cloudfront distribution's domain name."
  value       = "${aws_cloudfront_distribution.public_bucket.domain_name}"
}
