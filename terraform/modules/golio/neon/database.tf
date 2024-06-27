resource "neon_branch" "this" {
  project_id = neon_project.this.id
  name       = "${var.prefix}-golio"
}

resource "neon_database" "this" {
  project_id = neon_project.this.id
  branch_id  = neon_branch.this.id
  name       = "golio"
  owner_name = neon_role.this.name
}
