resource "aws_s3_bucket" "static_website" {
  bucket = "targetblank-static-website"
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
      "Resource":["arn:aws:s3:::targetblank-static-website/*"]
    }
  ]
}
EOF
}

resource "aws_s3_bucket_object" "root" {
  bucket       = "${aws_s3_bucket.static_website.bucket}"
  key          = "index.html"
  source       = ".build/index.html"
  content_type = "text/html"
  etag         = "${md5(file(".build/index.html"))}"
}