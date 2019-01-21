resource "aws_lambda_function" "authenticate" {
  function_name    = "authenticate"
  filename         = ".build/authenticate.zip"
  source_code_hash = "${base64sha256(file(".build/authenticate.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
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
  role             = "${aws_iam_role.lambda.arn}"
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
  role             = "${aws_iam_role.lambda.arn}"
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
  filename         = ".build/update.zip"
  source_code_hash = "${base64sha256(file(".build/update.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "update"
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
  filename         = ".build/read.zip"
  source_code_hash = "${base64sha256(file(".build/read.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "read"
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
  filename         = ".build/validate.zip"
  source_code_hash = "${base64sha256(file(".build/validate.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "validate"
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
  filename         = ".build/passwd.zip"
  source_code_hash = "${base64sha256(file(".build/passwd.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "passwd"
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
  filename         = ".build/reset.zip"
  source_code_hash = "${base64sha256(file(".build/reset.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "reset"
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
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "publish"
  runtime          = "go1.x"
}

resource "aws_lambda_permission" "publish" {
  statement_id  = "AllowGatewayInvoke"
  function_name = "${aws_lambda_function.publish.arn}"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
}
