resource "aws_dynamodb_table" "pages" {
  name           = "targetblank-pages"
  hash_key       = "addr"
  write_capacity = 1
  read_capacity  = 1

  point_in_time_recovery {
    enabled = true
  }

  attribute {
    name = "addr"
    type = "S"
  }
}
