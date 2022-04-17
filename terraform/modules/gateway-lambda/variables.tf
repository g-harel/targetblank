variable "rest_api_id" {
  description = "Root API Gateway REST API's ID."
}

variable "gateway_resource_id" {
  description = "The ID for the API Gateway resource (path) to which the function will be attached."
}

variable "http_method" {
  description = "HTTP method on the API Gateway resource (path) which the Lambda function is invoked on."
}

variable "bin" {
  description = "Local path to the compiled Lambda function binary."
}

variable "name" {
  description = "Unique name given to the Lambda function."
}

variable "role" {
  description = "IAM role attached to the Lambda function."
}

variable "tags" {
  description = "Tags to assign to the Lambda function."
  type        = map(string)
  default     = {}
}

variable "memory" {
  description = "Lambda function's memory."
  default     = 128
}
