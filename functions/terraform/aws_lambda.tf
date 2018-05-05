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
