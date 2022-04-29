resource "aws_api_gateway_method" "cors" {
  rest_api_id   = var.rest_api_id
  resource_id   = var.gateway_resource_id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "cors" {
  rest_api_id = var.rest_api_id
  resource_id = var.gateway_resource_id
  http_method = aws_api_gateway_method.cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers"     = true
    "method.response.header.Access-Control-Allow-Origin"      = true
    "method.response.header.Access-Control-Allow-Methods"     = true
    "method.response.header.Access-Control-Allow-Credentials" = true
  }

  depends_on = [aws_api_gateway_method.cors]
}

resource "aws_api_gateway_integration" "cors" {
  rest_api_id = var.rest_api_id
  resource_id = var.gateway_resource_id
  http_method = aws_api_gateway_method.cors.http_method
  type        = "MOCK"

  passthrough_behavior = "WHEN_NO_MATCH"

  request_templates = {
    "application/json" = "{'statusCode': 200}"
  }

  depends_on = [aws_api_gateway_method.cors]
}

resource "aws_api_gateway_integration_response" "cors" {
  rest_api_id = var.rest_api_id
  resource_id = var.gateway_resource_id
  http_method = aws_api_gateway_method.cors.http_method
  status_code = aws_api_gateway_method_response.cors.status_code

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin"      = "'${var.allow_origin}'"
    "method.response.header.Access-Control-Allow-Methods"     = "'${join(",", var.allow_methods)}'"
    "method.response.header.Access-Control-Allow-Headers"     = "'${join(",", var.allow_headers)}'"
    "method.response.header.Access-Control-Allow-Credentials" = "'${var.allow_credentials}'"
  }

  depends_on = [aws_api_gateway_method_response.cors]
}
