resource "aws_api_gateway_rest_api" "gateway" {
  name = "targetblank-api"
}

resource "aws_api_gateway_deployment" "deployment" {
  depends_on = [
    "aws_api_gateway_method.auth_addr_delete",
    "aws_api_gateway_method.auth_addr_post",
    "aws_api_gateway_method.auth_addr_put",
    "aws_api_gateway_method.page_addr_delete",
    "aws_api_gateway_method.page_addr_get",
    "aws_api_gateway_method.page_addr_patch",
    "aws_api_gateway_method.page_addr_put",
    "aws_api_gateway_method.page_post",
    "aws_api_gateway_method.page_validate_post",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  stage_name  = "prod"

  stage_description = "${md5(file("api/terraform/aws_api_gateway.tf"))}"
}

#
#

resource "aws_api_gateway_resource" "auth" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_rest_api.gateway.root_resource_id}"
  path_part   = "auth"
}

resource "aws_api_gateway_resource" "page" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_rest_api.gateway.root_resource_id}"
  path_part   = "page"
}

#

resource "aws_api_gateway_resource" "auth_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.auth.id}"
  path_part   = "{addr}"
}

resource "aws_api_gateway_resource" "page_addr" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.page.id}"
  path_part   = "{addr}"
}

resource "aws_api_gateway_resource" "page_validate" {
  rest_api_id = "${aws_api_gateway_rest_api.gateway.id}"
  parent_id   = "${aws_api_gateway_resource.page.id}"
  path_part   = "validate"
}

#
#

resource "aws_api_gateway_method" "auth_addr_delete" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.auth_addr.id}"
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_addr_delete" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.auth_addr.id}"
  http_method             = "${aws_api_gateway_method.auth_addr_delete.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.password_reset.arn}/invocations"
}

#

resource "aws_api_gateway_method" "auth_addr_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.auth_addr.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_addr_post" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.auth_addr.id}"
  http_method             = "${aws_api_gateway_method.auth_addr_post.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.authenticate.arn}/invocations"
}

#

resource "aws_api_gateway_method" "auth_addr_put" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.auth_addr.id}"
  http_method   = "PUT"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_addr_put" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.auth_addr.id}"
  http_method             = "${aws_api_gateway_method.auth_addr_put.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.password_change.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_validate_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page_validate.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_validate_post" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page_validate.id}"
  http_method             = "${aws_api_gateway_method.page_validate_post.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.page_validate.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_post" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page.id}"
  http_method             = "${aws_api_gateway_method.page_post.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.create.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_addr_delete" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page_addr.id}"
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_addr_delete" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page_addr.id}"
  http_method             = "${aws_api_gateway_method.page_addr_delete.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.delete.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_addr_get" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page_addr.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_addr_get" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page_addr.id}"
  http_method             = "${aws_api_gateway_method.page_addr_get.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.page_fetch.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_addr_patch" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page_addr.id}"
  http_method   = "PATCH"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_addr_patch" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page_addr.id}"
  http_method             = "${aws_api_gateway_method.page_addr_patch.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.publish.arn}/invocations"
}

#

resource "aws_api_gateway_method" "page_addr_put" {
  rest_api_id   = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id   = "${aws_api_gateway_resource.page_addr.id}"
  http_method   = "PUT"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "page_addr_put" {
  rest_api_id             = "${aws_api_gateway_rest_api.gateway.id}"
  resource_id             = "${aws_api_gateway_resource.page_addr.id}"
  http_method             = "${aws_api_gateway_method.page_addr_put.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.page_edit.arn}/invocations"
}
