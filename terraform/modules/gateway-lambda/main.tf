resource "aws_api_gateway_method" "auth_addr_delete" {
  rest_api_id   = "${aws_api_gateway_rest_api.rest_api.id}"
  resource_id   = "${aws_api_gateway_resource.auth_addr.id}"
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_addr_delete" {
  rest_api_id             = "${aws_api_gateway_rest_api.rest_api.id}"
  resource_id             = "${aws_api_gateway_resource.auth_addr.id}"
  http_method             = "${aws_api_gateway_method.auth_addr_delete.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.password_reset.arn}/invocations"
}

resource "aws_lambda_function" "password_reset" {
  function_name    = "password_reset"
  filename         = ".build/reset.zip"
  source_code_hash = "${base64sha256(file(".build/reset.zip"))}"
  role             = "${var.role}"
  handler          = "reset"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "password_reset" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.password_reset.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}
