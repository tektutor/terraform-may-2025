variable "docker_image_name" {
   description = "Name of the docker image"
   type        = string
   default     = "tektutor/ubuntu-ansible-node:latest"
}

variable "container_count" {
  description = "Number of containers you wish to create"
  type        = number
  
  validation {
     condition = var.container_count > 1 && var.container_count < 10
     error_message = "The container count should be in the range 2 to 9"
  }
}
