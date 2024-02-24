resource "dockerhub_repository" "project" {
  name        = var.repository_name
  namespace   = var.dockerhub_username
  description = "Project description"
}
