terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.5.0"
    }
    local = {
      source = "hashicorp/local"
      version = "2.5.2"
    }
  }
}

provider "docker" {
  # Configuration options
}
