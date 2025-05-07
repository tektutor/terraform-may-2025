variable "image_name" {
  description = "Docker image name"
  type = string
  default= "tektutor/ubuntu-ansible-node:latest"
}

variable "container_count" {
  description = "Number of container"
  type = number
}
