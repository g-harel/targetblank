resource "aws_lambda_function" "page_get" {
  function_name    = "page_get"
  filename         = ".build/page_get.zip"
  source_code_hash = "${base64sha256(file(".build/page_get.zip"))}"
  role             = "${aws_iam_role.lambda.arn}"
  handler          = "page_get"
  runtime          = "go1.x"
}
