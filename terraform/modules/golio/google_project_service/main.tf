# 手動でないとダメだった
# resource "google_project_service" "cloudresourcemanager" {
#   service = "cloudresourcemanager.googleapis.com"
#   project = var.gcp.project_id
# }


resource "google_project_service" "iam_google_api" {
  service = "iam.googleapis.com"
  project = var.gcp.project_id
#   depends_on = [ google_project_service.cloudresourcemanager ]
}



resource "google_project_service" "custom_search_api" {
  service = "customsearch.googleapis.com"
  project = var.gcp.project_id
  depends_on = [
    google_project_service.iam_google_api
  ]
}
