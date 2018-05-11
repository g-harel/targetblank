resource "aws_lambda_function" "page_fetch" {
  function_name    = "page_fetch"
  filename         = ".build/page_fetch.zip"
  source_code_hash = "${base64sha256(file(".build/page_fetch.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_fetch"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "page_validate" {
  function_name    = "page_validate"
  filename         = ".build/page_validate.zip"
  source_code_hash = "${base64sha256(file(".build/page_validate.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_validate"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "page_create" {
  function_name    = "page_create"
  filename         = ".build/page_create.zip"
  source_code_hash = "${base64sha256(file(".build/page_create.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_create"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "page_edit" {
  function_name    = "page_edit"
  filename         = ".build/page_edit.zip"
  source_code_hash = "${base64sha256(file(".build/page_edit.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_edit"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "page_publish" {
  function_name    = "page_publish"
  filename         = ".build/page_publish.zip"
  source_code_hash = "${base64sha256(file(".build/page_publish.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_publish"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "password_change" {
  function_name    = "password_change"
  filename         = ".build/password_change.zip"
  source_code_hash = "${base64sha256(file(".build/password_change.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "password_change"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "password_reset" {
  function_name    = "password_reset"
  filename         = ".build/password_reset.zip"
  source_code_hash = "${base64sha256(file(".build/password_reset.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "password_reset"
  runtime          = "go1.x"
}
