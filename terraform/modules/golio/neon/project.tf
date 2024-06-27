resource "neon_project" "this" {
  name                      = "${var.prefix}-folio"
  history_retention_seconds = 86400
  region_id                 = "aws-ap-southeast-1"
}
