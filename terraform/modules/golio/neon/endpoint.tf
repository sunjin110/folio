resource "neon_endpoint" "this" {
  project_id = neon_project.this.id
  branch_id  = neon_branch.this.id
  type       = "read_write"
  region_id  = "aws-ap-southeast-1"
}
