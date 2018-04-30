resource "aws_route53_zone" "primary" {
  name = "${var.domain}"
}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "${var.domain}"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.website.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.website.domain_name}"
    evaluate_target_health = false
  }
}

resource "aws_cloudfront_distribution" "website" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2"]
    }

    domain_name = "${aws_s3_bucket.website.id}.s3-website.${var.region}.amazonaws.com"
    origin_id   = "${aws_s3_bucket.website.id}"
  }

  enabled             = true
  default_root_object = "index.html"
  aliases             = ["${var.domain}"]
  http_version        = "http2"

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  default_cache_behavior {
    allowed_methods        = ["HEAD", "GET", "OPTIONS"]
    cached_methods         = ["HEAD", "GET", "OPTIONS"]
    target_origin_id       = "${aws_s3_bucket.website.id}"
    viewer_protocol_policy = "redirect-to-https"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true # TODO
  }
}

resource "aws_s3_bucket" "website" {
  bucket = "${var.name}-static-files"
  acl    = "public-read"

  website {
    index_document = "index.html"
  }

  policy = <<EOF
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"AddPerm",
      "Effect":"Allow",
      "Principal": "*",
      "Action":["s3:GetObject"],
      "Resource":["arn:aws:s3:::${var.name}-static-files/*"]
    }
  ]
}
EOF
}

resource "aws_s3_bucket_object" "home" {
  bucket       = "${aws_s3_bucket.website.bucket}"
  key          = "index.html"
  source       = "../website/index.html"
  content_type = "text/html"
}
