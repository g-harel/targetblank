resource "aws_s3_bucket" "public_bucket" {
  bucket = "${var.bucket_name}"
  acl    = "public-read"

  website {
    index_document = "${var.root_file}"
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
      "Resource":["arn:aws:s3:::${var.bucket_name}/*"]
    }
  ]
}
EOF
}

resource "aws_s3_bucket_object" "root" {
  bucket       = "${aws_s3_bucket.public_bucket.bucket}"
  key          = "${var.root_file}"
  source       = "${var.source_dir}/${var.root_file}"
  etag         = "${md5(file("${var.source_dir}/${var.root_file}"))}"
}

resource "aws_s3_bucket_object" "files" {
  count        = "${length(var.source_files)}"
  bucket       = "${aws_s3_bucket.public_bucket.bucket}"
  key          = "${element(var.source_files, count.index)}"
  source       = "${var.source_dir}/${element(var.source_files, count.index)}"
  etag         = "${md5(file("${var.source_dir}/${element(var.source_files, count.index)}"))}"
}

resource "aws_cloudfront_distribution" "public_bucket" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2"]
    }

    domain_name = "${aws_s3_bucket.public_bucket.id}.s3-website.${aws_s3_bucket.public_bucket.region}.amazonaws.com"
    origin_id   = "${aws_s3_bucket.public_bucket.id}"
  }

  enabled             = true
  default_root_object = "${aws_s3_bucket_object.root.key}"
  aliases             = ["${var.alias_name}"]
  http_version        = "http2"

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/${aws_s3_bucket_object.root.key}"
  }

  default_cache_behavior {
    allowed_methods        = ["HEAD", "GET", "OPTIONS"]
    cached_methods         = ["HEAD", "GET", "OPTIONS"]
    target_origin_id       = "${aws_s3_bucket.public_bucket.id}"
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
    acm_certificate_arn = "${var.cert_arn}"
    ssl_support_method  = "sni-only"
  }
}

resource "aws_route53_record" "cloudfront_alias" {
  zone_id = "${var.primary_zone_id}"
  name    = "${var.alias_name}"
  type    = "A"

  alias {
    zone_id                = "${aws_cloudfront_distribution.public_bucket.hosted_zone_id}"
    name                   = "${aws_cloudfront_distribution.public_bucket.domain_name}"
    evaluate_target_health = false
  }
}
