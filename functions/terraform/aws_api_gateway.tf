resource "aws_api_gateway_rest_api" "gateway" {
  name = "targetblank-api"
}

resource "aws_api_gateway_deployment" "deployment" {
  count = 0 # remove when ready to deploy

  depends_on = [
    "aws_api_gateway_method.api_v1_auth_addr_delete",
    "aws_api_gateway_method.api_v1_auth_addr_post",
    "aws_api_gateway_method.api_v1_auth_addr_put",
    "aws_api_gateway_method.api_v1_page_post",
    "aws_api_gateway_method.api_v1_page_addr_delete",
    "aws_api_gateway_method.api_v1_page_addr_get",
    "aws_api_gateway_method.api_v1_page_addr_patch",
    "aws_api_gateway_method.api_v1_page_addr_put",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  stage_name  = "prod"
}

resource "aws_api_gateway_resource" "api" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_rest_api.gateway.root_resource_id}"
  path_part   = "api"
}

resource "aws_api_gateway_resource" "api_v1" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.api.id}"
  path_part   = "v1"
}

resource "aws_api_gateway_resource" "api_v1_auth" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.api_v1.id}"
  path_part   = "auth"
}

resource "aws_api_gateway_resource" "api_v1_page" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.api_v1.id}"
  path_part   = "page"
}

resource "aws_api_gateway_resource" "api_v1_auth_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.api_v1_auth.id}"
  path_part   = "{addr}"
}

resource "aws_api_gateway_resource" "api_v1_page_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.api_v1_page.id}"
  path_part   = "{addr}"
}

resource "aws_api_gateway_method" "api_v1_auth_addr_delete" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_auth_addr.id}"
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_auth_addr_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_auth_addr.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_auth_addr_put" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_auth_addr.id}"
  http_method   = "PUT"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_page_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_page.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_page_addr_delete" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_page_addr.id}"
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_page_addr_get" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_page_addr.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_page_addr_patch" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_page_addr.id}"
  http_method   = "PATCH"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "api_v1_page_addr_put" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.api_v1_page_addr.id}"
  http_method   = "PUT"
  authorization = "NONE"
}
