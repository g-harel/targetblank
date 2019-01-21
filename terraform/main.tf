locals {
  domain_name = "targetblank.org"
}

module "website" {
  source = "./modules/bucket-public"

  aliases  = ["${local.domain_name}"]
  cert_arn    = "${aws_acm_certificate.ssl_cert.arn}"
  bucket_name = "targetblank-static-website"

  source_dir    = ".build"
  root_document = "index.html"

  files {
    "website.f69400ca.css" = "text/css"
    "website.f69400ca.js"  = "application/javascript"
  }
}
