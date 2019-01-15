output "cloudfront_zone_id" {
    description = "The Cloudfront distribution's zone ID."
    value       = "${aws_api_gateway_domain_name.domain_name.cloudfront_zone_id}"
}

output "cloudfront_domain_name" {
    description = "The Cloudfront distribution's domain name."
    value       = "${aws_api_gateway_domain_name.domain_name.cloudfront_domain_name}"
}
