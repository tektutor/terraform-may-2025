output "container_name1" {
  value = docker_container.my_ubuntu_container1.name
}

output "container_name2" {
  value = docker_container.my_rocky_container1.name
}
