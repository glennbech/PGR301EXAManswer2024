terraform {
  required_version = ">= 1.6.4"

  required_providers {
    dockerhub = {
      source  = "BarnabyShearer/dockerhub"
      version = ">= 0.0.15"
    }
  }
}

provider "dockerhub" {
  username = var.dockerhub_username
  password = var.dockerhub_password
}