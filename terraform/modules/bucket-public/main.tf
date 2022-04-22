resource "aws_s3_bucket" "public_bucket" {
  bucket = "${var.bucket_name}"
}

resource "aws_s3_bucket_acl" "bucket_acl" {
  bucket = aws_s3_bucket.public_bucket.id
  acl    = "public-read"
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.public_bucket.id
  policy = data.aws_iam_policy_document.policy_document.json
}

data "aws_iam_policy_document" "policy_document" {
  statement {
    sid = "PublicRead"
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.public_bucket.arn}/*"]
  }
}

resource "aws_s3_bucket_website_configuration" "website_configuration" {
  bucket = aws_s3_bucket.public_bucket.bucket

  index_document {
    suffix = "${var.root_document}"
  }
}

resource "aws_s3_object" "root" {
  bucket        = aws_s3_bucket.public_bucket.id
  key           = "${var.root_document}"
  source        = "${var.source_dir}/${var.root_document}"
  content_type  = "text/html"
  cache_control = "no-cache"
  etag          = "${md5(file("${var.source_dir}/${var.root_document}"))}"
}

resource "aws_s3_object" "files" {
  count        = "${length(var.files)}"
  bucket       = aws_s3_bucket.public_bucket.id
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
