variable "role" {
  description = "role for all lambda funcs"
}

resource "aws_lambda_function" "authenticate" {
  function_name    = "authenticate"
  filename         = ".build/authenticate.zip"
  source_code_hash = "${base64sha256(file(".build/authenticate.zip"))}"
  role             = "${var.role}"
  handler          = "authenticate"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "authenticate" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.authenticate.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "create" {
  function_name    = "create"
  filename         = ".build/create.zip"
  source_code_hash = "${base64sha256(file(".build/create.zip"))}"
  role             = "${var.role}"
  handler          = "create"
  runtime          = "go1.x"
  memory_size      = 320
}

resource "aws_lambda_permission" "create" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.create.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "delete" {
  function_name    = "delete"
  filename         = ".build/delete.zip"
  source_code_hash = "${base64sha256(file(".build/delete.zip"))}"
  role             = "${var.role}"
  handler          = "delete"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "delete" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.delete.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "page_edit" {
  function_name    = "page_edit"
  filename         = ".build/page_edit.zip"
  source_code_hash = "${base64sha256(file(".build/page_edit.zip"))}"
  role             = "${var.role}"
  handler          = "page_edit"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "page_edit" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.page_edit.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "page_fetch" {
  function_name    = "page_fetch"
  filename         = ".build/page_fetch.zip"
  source_code_hash = "${base64sha256(file(".build/page_fetch.zip"))}"
  role             = "${var.role}"
  handler          = "page_fetch"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "page_fetch" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.page_fetch.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "page_validate" {
  function_name    = "page_validate"
  filename         = ".build/page_validate.zip"
  source_code_hash = "${base64sha256(file(".build/page_validate.zip"))}"
  role             = "${var.role}"
  handler          = "page_validate"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "page_validate" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.page_validate.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "password_change" {
  function_name    = "password_change"
  filename         = ".build/password_change.zip"
  source_code_hash = "${base64sha256(file(".build/password_change.zip"))}"
  role             = "${var.role}"
  handler          = "password_change"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "password_change" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.password_change.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "password_reset" {
  function_name    = "password_reset"
  filename         = ".build/password_reset.zip"
  source_code_hash = "${base64sha256(file(".build/password_reset.zip"))}"
  role             = "${var.role}"
  handler          = "password_reset"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "password_reset" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.password_reset.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}

#

resource "aws_lambda_function" "publish" {
  function_name    = "publish"
  filename         = ".build/publish.zip"
  source_code_hash = "${base64sha256(file(".build/publish.zip"))}"
  role             = "${var.role}"
  handler          = "publish"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "publish" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.publish.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}
