resource "aws_api_gateway_rest_api" "rest_api" {
  name = "targetblank-api"
}

resource "aws_api_gateway_deployment" "deployment" {
  depends_on = [
    "module.authenticate",
    "module.create",
    "module.passwd",
    "module.read",
    "module.reset",
    "module.update",
    "module.validate",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  stage_name  = "prod"

  stage_description = "${md5(file("terraform/aws_api_gateway.tf"))}"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_domain_name" "domain_name" {
  domain_name     = "api.${local.domain_name}"
  certificate_arn = "${aws_acm_certificate.ssl_cert.arn}"
}

resource "aws_api_gateway_base_path_mapping" "base_path_mapping" {
  api_id      = "${aws_api_gateway_rest_api.rest_api.id}"
  stage_name  = "${aws_api_gateway_deployment.deployment.stage_name}"
  domain_name = "${aws_api_gateway_domain_name.domain_name.domain_name}"
}

resource "aws_api_gateway_resource" "auth" {
  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  parent_id   = "${aws_api_gateway_rest_api.rest_api.root_resource_id}"
  path_part   = "auth"
}

module "cors_auth" {
  source = "modules/gateway-cors"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth.id}"
}

resource "aws_api_gateway_resource" "page" {
  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  parent_id   = "${aws_api_gateway_rest_api.rest_api.root_resource_id}"
  path_part   = "page"
}

module "cors_page" {
  source = "modules/gateway-cors"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page.id}"
}

resource "aws_api_gateway_resource" "auth_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  parent_id   = "${aws_api_gateway_resource.auth.id}"
  path_part   = "{addr}"
}

module "cors_auth_addr" {
  source = "modules/gateway-cors"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.auth_addr.id}"
}

resource "aws_api_gateway_resource" "page_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  parent_id   = "${aws_api_gateway_resource.page.id}"
  path_part   = "{addr}"
}

module "cors_page_addr" {
  source = "modules/gateway-cors"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_addr.id}"
}

resource "aws_api_gateway_resource" "page_validate" {
  rest_api_id = "${aws_api_gateway_rest_api.rest_api.id}"
  parent_id   = "${aws_api_gateway_resource.page.id}"
  path_part   = "validate"
}

module "cors_page_validate" {
  source = "modules/gateway-cors"

  rest_api_id         = "${aws_api_gateway_rest_api.rest_api.id}"
  gateway_resource_id = "${aws_api_gateway_resource.page_validate.id}"
}
