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

resource "aws_lambda_function" "create" {
  function_name    = "create"
  filename         = ".build/create.zip"
  source_code_hash = "${base64sha256(file(".build/create.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "create"
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

resource "aws_lambda_function" "publish" {
  function_name    = "publish"
  filename         = ".build/publish.zip"
  source_code_hash = "${base64sha256(file(".build/publish.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "publish"
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

resource "aws_lambda_function" "delete" {
  function_name    = "delete"
  filename         = ".build/delete.zip"
  source_code_hash = "${base64sha256(file(".build/delete.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "delete"
  runtime          = "go1.x"
}
