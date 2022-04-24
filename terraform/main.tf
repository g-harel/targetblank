locals {
  domain_name = "targetblank.org"

  parcel_manifest = jsondecode(file(".website/parcel-manifest.json"))

  lambda_tags = {
    project = "targetblank"
  }
}

module "website" {
  source = "./modules/bucket-public"

  aliases     = ["${local.domain_name}"]
  cert_arn    = "${aws_acm_certificate_validation.cert.certificate_arn}"
  bucket_name = "targetblank-static-website"

  source_dir    = ".website"
  root_document = "index.html"

  files = [
    "${lookup(local.parcel_manifest, "index.tsx")}",
    "${lookup(local.parcel_manifest, "favicon.png")}",
  ]
}

module "authenticate" {
  source = "./modules/gateway-lambda"

  name = "authenticate"
  bin  = ".functions/authenticate"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "POST"
}

module "create" {
  source = "./modules/gateway-lambda"

  name   = "create"
  bin    = ".functions/create"
  role   = "${aws_iam_role.lambda.arn}"
  tags   = "${local.lambda_tags}"
  memory = 320

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page.id}"
  http_method         = "POST"
}

module "passwd" {
  source = "./modules/gateway-lambda"

  name = "passwd"
  bin  = ".functions/passwd"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "PUT"
}

module "read" {
  source = "./modules/gateway-lambda"

  name = "read"
  bin  = ".functions/read"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "GET"
}

module "reset" {
  source = "./modules/gateway-lambda"

  name = "reset"
  bin  = ".functions/reset"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
  http_method         = "DELETE"
}

module "update" {
  source = "./modules/gateway-lambda"

  name = "update"
  bin  = ".functions/update"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
  http_method         = "PUT"
}

module "validate" {
  source = "./modules/gateway-lambda"

  name = "validate"
  bin  = ".functions/validate"
  role = "${aws_iam_role.lambda.arn}"
  tags = "${local.lambda_tags}"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_validate.id}"
  http_method         = "POST"
}
