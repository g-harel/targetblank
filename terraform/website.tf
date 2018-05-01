resource "aws_s3_bucket" "static_website" {
  bucket = "${var.name}-static-website"
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
      "Resource":["arn:aws:s3:::${var.name}-static-website/*"]
    }
  ]
}
EOF
}

resource "aws_acm_certificate" "ssl_cert" {
  domain_name               = "${var.domain}"
  validation_method         = "EMAIL"
  subject_alternative_names = ["*.${var.domain}"]
}

resource "aws_cloudfront_distribution" "static_website" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2"]
    }

    domain_name = "${aws_s3_bucket.static_website.id}.s3-website.${var.region}.amazonaws.com"
    origin_id   = "${aws_s3_bucket.static_website.id}"
  }

  enabled             = true
  default_root_object = "${aws_s3_bucket_object.root.key}"
  aliases             = ["${var.domain}"]
  http_version        = "http2"

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/${aws_s3_bucket_object.root.key}"
  }

  default_cache_behavior {
    allowed_methods        = ["HEAD", "GET", "OPTIONS"]
    cached_methods         = ["HEAD", "GET", "OPTIONS"]
    target_origin_id       = "${aws_s3_bucket.static_website.id}"
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
    acm_certificate_arn = "${aws_acm_certificate.ssl_cert.arn}"
    ssl_support_method  = "sni-only"
  }
}

resource "aws_route53_zone" "primary" {
  name = "${var.domain}"
}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${aws_route53_zone.primary.zone_id}"
  name    = "${var.domain}"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.static_website.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.static_website.domain_name}"
    evaluate_target_health = false
  }
}

resource "aws_s3_bucket_object" "root" {
  bucket       = "${aws_s3_bucket.static_website.bucket}"
  key          = "index.html"
  source       = "../website/index.html"
  content_type = "text/html"
  etag         = "${md5(file("../website/index.html"))}"
}
