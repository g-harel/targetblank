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

variable "root_file" {
    description = "default file served at the root and used as fallback"
}

variable "source_files" {
    description = "file paths in the source directory to include in the bucket"
    type = "list"
}
