data "docker_image" "tektutor_image" {
   name = var.docker_image_name
}

resource "docker_container" "my_container" {
   count = var.container_count

   image = data.docker_image.tektutor_image.name
   name  = "container-${count.index+1}"
   hostname  = "container-${count.index+1}"

   ports {
     internal = "22"
     external = "200${count.index+1}"
   }
   ports {
     internal = "80"
     external = "800${count.index+1}"
   }

}
