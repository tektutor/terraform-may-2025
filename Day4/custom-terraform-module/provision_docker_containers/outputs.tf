output "container_name" {
   value = docker_container.my_container[*].name 
}
