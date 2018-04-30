resource "aws_route53_zone" "primary" {
  name = "${var.name}.org"
}

resource "aws_route53_record" "record" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "${var.name}.org"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.static_distribution.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.static_distribution.domain_name}"
    evaluate_target_health = false
  }
}

resource "aws_s3_bucket" "static_files" {
  bucket = "${var.name}-static-files"

  website {
    index_document = "index.html"
  }
}

resource "aws_cloudfront_distribution" "static_distribution" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2"]
    }

    domain_name = "${aws_s3_bucket.static_files.id}.s3-website-${var.region}.amazonaws.com"
    origin_id   = "${aws_s3_bucket.static_files.id}"
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
    target_origin_id       = "${aws_s3_bucket.static_files.id}"
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
