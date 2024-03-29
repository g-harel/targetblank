data "archive_file" "file" {
  type        = "zip"
  source_file = "${path.root}/${var.bin}"
  output_path = "${var.bin}.zip"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = var.rest_api_id
  resource_id   = var.gateway_resource_id
  http_method   = var.http_method
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = var.rest_api_id
  resource_id             = var.gateway_resource_id
  http_method             = var.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.function.arn}/invocations"
}

resource "aws_lambda_function" "function" {
  function_name    = var.name
  filename         = data.archive_file.file.output_path
  source_code_hash = data.archive_file.file.output_base64sha256
  role             = var.role
  handler          = basename(var.bin)
  runtime          = "go1.x"
  tags             = var.tags
  memory_size      = var.memory
}

resource "aws_lambda_permission" "permission" {
  statement_id  = "AllowGatewayInvoke"
  function_name = aws_lambda_function.function.arn
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}
