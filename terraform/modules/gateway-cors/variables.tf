variable "rest_api_id" {
  description = "Root API Gateway REST API's ID."
}

variable "gateway_resource_id" {
  description = "The ID for the API Gateway resource (path) which will allow CORS."
}

variable "allow_origin" {
  description = "Allowed origin."
  type        = "string"
  default     = "*"
}

variable "allow_methods" {
  description = "Allowed HTTP methods."
  type        = "list"
  default     = ["GET", "POST", "PUT", "DELETE", "PATCH"]
}

variable "allow_headers" {
  description = "Allowed HTTP headers."
  type        = "list"
  default     = ["Content-Type", "Authorization"]
}

variable "allow_credentials" {
  description = "Allow caller to view response after sending credentials."
  type        = "string"
  default     = "true"
}

