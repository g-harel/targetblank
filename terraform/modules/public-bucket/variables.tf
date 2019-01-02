variable zone_id {
  description = "aws_route53_zone.<name>.zone_id"
}

variable alias_name {
  description = "cloudfront's alias"
}

variable cert_arn {
  description = "acl certificate arn"
}

variable "bucket_name" {
  description = "bucket name"
}

variable "source_dir" {
  description = "source directory"
}

variable "root_document" {
  description = "default html document served at the root and used as fallback"
}

variable "files" {
  description = "file paths (+ mime type) in the source directory to include in the bucket"
  type        = "map"
}
