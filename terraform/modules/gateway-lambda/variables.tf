variable "rest_api_id" {
  description = "Root API Gateway REST API's ID."
}

variable "gateway_resource_id" {
  description = "The ID for the API Gateway resource (path) to which the function will be attached."
}

variable "http_method" {
  description = "HTTP method on the API Gateway resource (path) which the Lambda function is invoked on."
}

variable "file" {
  description = "Local path to the Lambda function's deployment package."
}

variable "name" {
  description = "Unique name given to the Lambda function."
}

variable "handler_name" {
  description = "Lambda funtion's entrypoint"
  default     = "handler"
}

variable "runtime" {
  description = "An identifier for the Lmabda function's runtime."
  default     = "go1.x"
}

variable "role" {
  description = "IAM role attached to the Lambda function."
}

variable "tags" {
  description = "Tags to assign to the Lambda function."
  type        = "map"
  default     = {}
}

variable "memory" {
  description = "Lambda function's memory."
  default     = 128
}

