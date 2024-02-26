variable "dockerhub_username" {
  type        = string
  description = "Docker Hub username"
  sensitive   = true
}
variable "dockerhub_password" {
  type        = string
  description = "Docker Hub password"
  sensitive   = true
}
variable "repository_name" {
  type        = string
  description = "Docker Hub repository name"
  default     = "nbx"
}