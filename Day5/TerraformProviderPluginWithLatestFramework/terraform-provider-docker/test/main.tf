terraform {
  required_providers {
    docker = {
       source = "tektutor/docker"
    }
  }
}

resource "docker_container" "container1" {
  container_name = "c1"
  container_hostname      = "container_1"
  container_image     = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "container2" {
  container_name = "c2"
  container_hostname      = "container_2"
  container_image     = "tektutor/ubuntu-ansible-node:latest"
}
