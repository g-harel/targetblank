variable "aliases" {
  description = "The Cloudfront distribution's aliases."
  type        = list(string)
}

variable "cert_arn" {
  description = "The ARN of the certificate used by the Cloudfront distribution."
}

variable "bucket_name" {
  description = "Name of the public S3 bucket where files will be stored."
}

variable "source_dir" {
  description = "Source directory for all files being uploaded."
}

variable "root_document" {
  description = "Default html document served at the root and used as custom 404 fallback."
}

variable "files" {
  description = "List of files to be included in the bucket."
  type        = list(string)
}
