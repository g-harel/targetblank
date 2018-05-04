resource "aws_dynamodb_table" "pages" {
  name           = "targetblank-pages"
  hash_key       = "addr"
  write_capacity = 1
  read_capacity  = 1

  attribute {
    name = "addr"
    type = "S"
  }
}
