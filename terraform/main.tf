locals {
  domain_name = "targetblank.org"
  lambda_tags = {
    project = "targetblank"
  }
}

module "website" {
  source = "./modules/bucket-public"

  aliases  = ["${local.domain_name}"]
  cert_arn    = "${aws_acm_certificate.ssl_cert.arn}"
  bucket_name = "targetblank-static-website"

  source_dir    = ".build"
  root_document = "index.html"

  files = [
    "website.f69400ca.css",
    "website.f69400ca.js",
  ]
}

module "authenticate" {
  source = "./modules/gateway-lambda"

  name = "authenticate"
  file = ".build/authenticate.zip"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "POST"
}

module "create" {
  source = "./modules/gateway-lambda"

  name   = "create"
  file   = ".build/create.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"
  memory = 320

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page.id}"
  http_method         = "POST"
}

module "delete" {
  source = "./modules/gateway-lambda"

  name   = "delete"
  file   = ".build/delete.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "DELETE"
}

module "passwd" {
  source = "./modules/gateway-lambda"

  name   = "passwd"
  file   = ".build/passwd.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "PUT"
}

module "publish" {
  source = "./modules/gateway-lambda"

  name   = "publish"
  file   = ".build/publish.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "PATCH"
}

module "read" {
  source = "./modules/gateway-lambda"

  name   = "read"
  file   = ".build/read.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "GET"
}

module "reset" {
  source = "./modules/gateway-lambda"

  name   = "reset"
  file   = ".build/reset.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "DELETE"
}

module "update" {
  source = "./modules/gateway-lambda"

  name   = "update"
  file   = ".build/update.zip"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "PUT"
}

module "validate" {
  source = "./modules/gateway-lambda"

  name = "validate"
  file = ".build/validate.zip"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_validate.id}"
  http_method         = "POST"
}
