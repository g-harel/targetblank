resource "aws_s3_bucket" "public_bucket" {
  bucket = "${var.bucket_name}"
  acl    = "public-read"

  website {
    index_document = "${var.root_document}"
  }

  policy = <<EOF
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"PublicRead",
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
  key          = "${var.root_document}"
  source       = "${var.source_dir}/${var.root_document}"
  content_type = "text/html"
  etag         = "${md5(file("${var.source_dir}/${var.root_document}"))}"
}

resource "aws_s3_bucket_object" "files" {
  count        = "${length(var.files)}"
  bucket       = "${aws_s3_bucket.public_bucket.bucket}"
  key          = "${element(var.files, count.index)}"
  source       = "${var.source_dir}/${element(var.files, count.index)}"
  content_type = "${lookup(local.mime, replace(element(var.files, count.index), "/^.*\\.(\\w+)$/", "$1"), "text/plain")}"
  etag         = "${md5(file("${var.source_dir}/${element(var.files, count.index)}"))}"
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
  aliases             = "${var.aliases}"
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
    compress               = true

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
