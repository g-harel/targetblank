resource "aws_iam_role" "lambda" {
  name = "lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "lambda_dynamodb" {
  name = "dynamo-policy"
  role = "${aws_iam_role.lambda.id}"

  policy = <<EOF
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Action": "dynamodb:*",
            "Resource": "${aws_dynamodb_table.pages.arn}",
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}
