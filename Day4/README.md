# Day 4

## Lab - Terraform modules
```
cd ~/terraform-may-2025
git pull
cd Day4/terraform/custom-terraform-module
terraform init
terraform apply --auto-approve
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/e9048a00-9068-439b-aaa4-9e1100f3d2d3)
![image](https://github.com/user-attachments/assets/78219f90-6d32-47e2-ba43-1e9bf434fdbb)
![image](https://github.com/user-attachments/assets/1723b206-6d69-4468-b1b5-f7ea09066c7f)
![image](https://github.com/user-attachments/assets/c7c52aa4-4434-4538-9b0f-0aaf322473ff)
![image](https://github.com/user-attachments/assets/b22e61b7-928d-407f-95f9-36bca1b3c377)
![image](https://github.com/user-attachments/assets/fb45d21b-314b-4861-a875-6300ce910425)
![image](https://github.com/user-attachments/assets/d33c3f20-4336-4d1b-b320-42948210f61c)
![image](https://github.com/user-attachments/assets/a4a56b7f-e77a-4358-95a0-e9383d7915fb)

Once you are done with this exercise, make sure you are disposing the resources
```
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/cc0206b7-886a-434f-aa19-ddcca6d02839)
![image](https://github.com/user-attachments/assets/67048628-2dc4-4c8d-8897-723a84232a2e)
![image](https://github.com/user-attachments/assets/409deae0-2e1a-4128-8e01-fc4757e406d2)
![image](https://github.com/user-attachments/assets/2c5b59c3-4833-4b1d-8b58-266d8138f487)


## Lab - Develop your first Custom Terraform Provider Plugin in Golang
You need to create a folder
```
mkdir -p /home/rps/go/bin
touch ~/.terraformrc
```
Paste the below code in your ~/.terraformrc file
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",    
  }
  direct {}
}
```

Then you may proceed with the below instructions
```
cd ~/terraform-may-2025
git pull
cd Day4/custom-terraform-providers/file
tree
go mod init github.com/tektutor/terraform-provider-file
go mod tidy
ls -l
go build
ls -l
go install
ls -l /home/rps/go/bin
```

Expected output
![image](https://github.com/user-attachments/assets/ce2331d5-cbeb-4164-bb8c-8fc627851059)
![image](https://github.com/user-attachments/assets/06772e84-0125-408d-9bed-7d7a6a389c7c)
![image](https://github.com/user-attachments/assets/a872714b-b2ba-47a7-96bd-fdd0746f7816)

## Lab - Using our custom terraform provider plugin in a terraform manifest script
```
cd ~/terraform-may-2025
git pull
cd Day4/custom-terraform-providers/test-file-custom-terraform-provider
cat main.tf
terraform plan
ls
terraform apply --auto-approve
ls
cat myfile.txt
cat terraform.tfstate
terraform show
```

Expected output
![image](https://github.com/user-attachments/assets/a917a884-ccb5-49d9-a5d9-ffeb5bba7608)
![image](https://github.com/user-attachments/assets/e9a0c240-079b-40a8-8297-ced508278d84)
![image](https://github.com/user-attachments/assets/5b202d14-230d-45df-80e7-b10fd0043fd6)
![image](https://github.com/user-attachments/assets/1b![image](https://github.com/user-attachments/assets/81e66e0d-ee62-416b-82a9-1a15750713b1)
5a2bd1-e981-4547-a81e-dcd69c25ac2f)
![image](https://github.com/user-attachments/assets/ba968ae7-a644-440f-aa11-af3a8311589c)

Try updating the main.tf file_content and try apply again
```
terraform apply --auto-approve
```

Once you are done with this lab exercise, you may dispose all the resources created by terraform
```
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/6959de02-da1c-4187-b623-8405a2c9df1a)

## Info - Docker SDK API Reference
<pre>
https://pkg.go.dev/github.com/docker/docker/client  
</pre>
![image](https://github.com/user-attachments/assets/bcf0c7e6-c75e-4295-a386-74ca60266f1f)
![image](https://github.com/user-attachments/assets/88aa80f3-09d8-4d21-ba30-bffd27a03504)
![image](https://github.com/user-attachments/assets/e05dae92-6d9b-4b93-8381-95ee8bc6f925)
![image](https://github.com/user-attachments/assets/6fb4bbc3-b144-4595-823d-883fa1d7d168)
![image](https://github.com/user-attachments/assets/9a468414-2f34-49a8-8306-49040565c988)
![image](https://github.com/user-attachments/assets/060a9964-cb7d-4819-ad6a-407b7b80f364)


## Lab - Developing a custom docker terrform provider plugin using Golang
You need to update your ~/.terraformrc file as shown below to help terraform locate the our docker provider
<pre>
Paste the below code in your ~/.terraformrc file
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",    
    "registry.terraform.io/tektutor/docker" = "/home/rps/go/bin",      
  }
  direct {}
}  
</pre>

```
cd ~/terraform-may-2025
git pull
cd Day4/custom-terraform-providers/docker
tree
go mod init github.com/tektutor/terraform-provider-docker
go mod tidy
ls -l
go build
ls -l
go install
ls -l /home/rps/go/bin
```

Expected output
![image](https://github.com/user-attachments/assets/7cd56862-231e-45f4-b4d5-471ce5091230)
![image](https://github.com/user-attachments/assets/54177880-5fcd-4f78-b478-614ae1321db3)
![image](https://github.com/user-attachments/assets/dd0d7c58-3272-49a4-8ce7-7bf4b7cade76)

## Lab - Using our custom docker terraform provider plugin in Terraform automation project
```
cd ~/terraform-may-2025
git pull
cd Day4/custom-terraform-providers/test-docker-custom-terraform-provider
terraform plan
terraform apply
docker ps
terraform show
```

Expected output
![image](https://github.com/user-attachments/assets/f7efc07a-e50d-48a3-9dd2-71644238eced)
![image](https://github.com/user-attachments/assets/9f4ed051-59a9-41cf-93cf-30e45ce88dc7)
![image](https://github.com/user-attachments/assets/500222d5-001b-4bf5-8df7-ade3aed77ec8)
![image](https://github.com/user-attachments/assets/6dbd423b-ebd4-4f83-bf0f-ed72e0b2b547)

Let's try updating the name of the container. You need to edit the main.tf as shown below
```
terraform {
  required_providers {
    docker = {
       source = "tektutor/docker"
    }
  }
}

resource "docker_container" "container1" {
  container_name = "c1"
  host_name      = "container_1"
  image_name     = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "container2" {
  container_name = "c2"
  host_name      = "container_2"
  image_name     = "tektutor/ubuntu-ansible-node:latest"
}
```
Let's apply the changes
```
terraform plan
terraform apply --auto-approve
terraform show
terraform providers
```

Expected output
![image](https://github.com/user-attachments/assets/54737fbb-d53a-4905-ae6c-86868f0ba539)
![image](https://github.com/user-attachments/assets/334b8f6e-8290-4ac6-8f7f-eb99140b54ae)
![image](https://github.com/user-attachments/assets/384da50e-7ebe-4612-817e-de12be25e39c)
![image](https://github.com/user-attachments/assets/fc989ce9-a21d-431e-85dc-b53be1296beb)
![image](https://github.com/user-attachments/assets/e4efc53b-0571-434e-9690-9541a601cebc)
![image](https://github.com/user-attachments/assets/220d3da0-dcee-44bc-9bfa-d853319943bd)
![image](https://github.com/user-attachments/assets/890e47f3-bb3b-42a9-90a6-0ded79799e06)
![image](https://github.com/user-attachments/assets/867031e2-1744-4a2f-a43a-f05c624f20a5)
![image](https://github.com/user-attachments/assets/3401875f-0cf0-474b-9da7-4d354f016cd9)
![image](https://github.com/user-attachments/assets/123337a3-db3b-41a7-b0a0-8e79fd7575c7)


Clean up the resources provisioned by Terraform
```
docker images | grep bitnami
terraform destroy --auto-approve
docker images | grep bitnami
```

Expected output
![image](https://github.com/user-attachments/assets/26a2e88d-3d9a-42d0-998e-5d80b8d108fd)
![image](https://github.com/user-attachments/assets/e0bbf5bf-9cb8-48a1-be0e-2e5a90074c8c)

## Info - Ansible Automation Platform
<pre>
- was also called as Ansible Tower
- it is developed on top the AWX (open source)
- AWX is developed on top Ansible core (open source)
- unlike Ansible core, Ansible automation platform supports webconsole, user management, etc.,
- this is a Red Hat Enterprise product with world-wide support
- functionally AWX and Ansible Automation Platform are one and same
- it is not possible develop Ansible Playbook within AWX or Ansible Automation Platform
- hence, we still need Ansible core to develop playbooks and test them before we push it to GitHub or any version control
- the existing Ansible Playbooks we can execute via Ansible Automation Platform
- can be installed as a stand-alone application on any Linux Distributions or we can install inside Kubernetes or Openshift
</pre>

## Lab - Starting the minikube ( Launch Ansible Tower )
```
docker ps -a
minikube start
minikube status
kubectl get nodes
```

Expected output
![image](https://github.com/user-attachments/assets/43f30def-806c-4abe-9a9e-e31dec3a4a4d)
![image](https://github.com/user-attachments/assets/4bc8a187-be25-485a-9a6d-78a0028b35de)

Accessing Ansible Tower Dashboard from chrome web browser on your RPS Cloud Ubuntu lab machine
```
minikube service awx-demo-service --url -n ansible-awx
```
You can launch the AWX Webconsole using the url shown in the above command
```
http://192.168.49.2:31225
```
![image](https://github.com/user-attachments/assets/20b63425-5071-4a42-8be9-7a2ac340aeee)

Ansible Tower Login Credentials, save the login credentials in a text file to avoid typing complex command each time
<pre>
username - admin
password - 
</pre>

To get the password, you to type the below command
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Once you login, you will get the below page
![image](https://github.com/user-attachments/assets/6e5b9009-cbc6-4117-9d05-64a1ec45c85a)

## Lab - Creating a Project in Ansible AWX

Navigate to Ansible Automation Platform on your RPS Lab machine - chrome web browser
![image](https://github.com/user-attachments/assets/485a5eb5-82e1-48da-8bec-f24959d8bb30)

On the menu that appears on the left side, Navigate to Resources --> Projects
![image](https://github.com/user-attachments/assets/6788d7bf-2b92-4b39-a23c-49fabac26f4e)

Click "Add"
![image](https://github.com/user-attachments/assets/834056ce-1f09-4ac3-8d2b-73259fbdba95)
<pre>
Under the Name, type "TekTutor Training Repository"
Under the Source Code Type, select "Git"
Under the Source Control url, type "https://github.com/tektutor/terraform-may-2025.git"
Under the Source Control Branch/Tag/Commit, type "main"
Under the Options, enable the check box that says "Update Revision on Launch"
</pre>
![image](https://github.com/user-attachments/assets/049d38d8-58f6-4ec0-9ec5-0f3e663e546f)
Click "Save"
![image](https://github.com/user-attachments/assets/07eec3fb-9e2a-4618-ad67-90faaa64a811)
![image](https://github.com/user-attachments/assets/1268b2a1-e6dd-4855-8b6f-ba11ea944a94)
Click "Successful"
![image](https://github.com/user-attachments/assets/15c2a342-c054-4608-bf30-dcbc889d28b1)
![image](https://github.com/user-attachments/assets/2a722f64-09bd-4d14-8b1b-753ae5b80609)
![image](https://github.com/user-attachments/assets/07b48cd4-25fb-4238-b7c6-097fa7ba7276)

## Lab - Creating Inventory in Ansible Automation Platform(AWX)

Navigate to Ansible Automation Platform(AWX)
![image](https://github.com/user-attachments/assets/5b24830d-be9e-45d2-87dd-cce1164d6ed8)

Click Resources --> Inventories
![image](https://github.com/user-attachments/assets/56d74187-79ac-4db8-9994-fdfd7ee1f6b9)
Click Add
Select the first option "Add Inventory
![image](https://github.com/user-attachments/assets/d2bb9ea8-d2a6-4673-9441-2a9f84a6bbe1)
![image](https://github.com/user-attachments/assets/a4bf5394-64e6-43f9-95c7-205847ce8d4d)
Under the Name, type "Docker Inventory"
![image](https://github.com/user-attachments/assets/84564e2d-20a2-4456-866e-f637f55f7785)
Click "Save"
![image](https://github.com/user-attachments/assets/bca311c6-40c3-4948-af77-7bbd082f23a3)
Click the Tab named "Hosts" within the Inventory you saved just now
![image](https://github.com/user-attachments/assets/1aae0b65-d039-4c04-ab22-14b69ca20787)
Click "Add"
![image](https://github.com/user-attachments/assets/3bac4ab6-8d34-47e8-af65-0b486427a9df)
![image](https://github.com/user-attachments/assets/63488e47-0ac4-4044-915a-f1f7f34ba8ef)
![image](https://github.com/user-attachments/assets/d1014742-6a15-42a0-bde7-1a0670bb5dd8)
Click "Save"
![image](https://github.com/user-attachments/assets/e28e910a-2561-49aa-9146-4e23d1d2b8e6)
![image](https://github.com/user-attachments/assets/3613c7ca-d9fd-4680-b6df-25700f4d08d0)
click Add to create other ansible nodes on the similar fashion
![image](https://github.com/user-attachments/assets/96ac35a4-9c25-422f-b48b-2996b8c92eb7)
![image](https://github.com/user-attachments/assets/1f00bb0d-222b-452b-a429-a5d6769c413a)
Click "Save"
![image](https://github.com/user-attachments/assets/8b33b848-a681-4639-9cfb-a90a88db5f0c)
click Add to create other ansible nodes on the similar fashion
![image](https://github.com/user-attachments/assets/a90d7c1c-e96f-43c6-8993-e415e328e6b7)
Click "Add"
![image](https://github.com/user-attachments/assets/126fba92-08f4-4b34-b600-67ebfa73569d)
![image](https://github.com/user-attachments/assets/6cde45f6-b205-4c0b-a994-60f3b177ff56)
Click "Save"
![image](https://github.com/user-attachments/assets/cbbae3de-7dd8-466d-9be8-d03eeffe2875)

Repeat the procedure to add "Rocky2"
![image](https://github.com/user-attachments/assets/d28d5766-b9a3-4123-b0b0-54f5380102fa)
![image](https://github.com/user-attachments/assets/2ac1aa8c-6d43-4fba-ba24-df9d6c219a12)
![image](https://github.com/user-attachments/assets/277a7572-241a-4e61-ab49-649994ea8b38)

To verify if all the hosts(ansible nodes) added to the inventory are reachable to Ansible Tower, Click on your inventory and move to the Hosts tab
![image](https://github.com/user-attachments/assets/fa8711d3-fa9c-4d8b-a5a1-26135aa67581)
Click "Run command"
![image](https://github.com/user-attachments/assets/ae3f6fdd-ee6c-4db0-82d2-ffc8d06cdea2)
Under the Module, choose "ping"
![image](https://github.com/user-attachments/assets/12d94a11-883b-4ae7-b65b-f1fef3b01012)
![image](https://github.com/user-attachments/assets/e0426c71-e47f-4e20-90ec-2cc96213122e)
Click "Next"
![image](https://github.com/user-attachments/assets/aa126029-827c-4301-81e1-9f9ad8148738)
Click "Next"
![image](https://github.com/user-attachments/assets/6f5079b4-970d-4fed-a084-a33975b81a65)
Select "RPS Private Key" we saved
Click "Next"
![image](https://github.com/user-attachments/assets/ed473b7d-6b76-44c0-bb29-0e9b73346128)
Click "Launch"
![image](https://github.com/user-attachments/assets/c935c5e5-89c9-4753-96ca-856d6abfac55)
![image](https://github.com/user-attachments/assets/f2e374f1-fdec-4657-aef7-39b48f898aac)
![image](https://github.com/user-attachments/assets/ea5cd484-b463-4f11-8d3d-8a3772876d1f)


## Lab - Creating Credentials to store the Private key 
Navigate to Ansible Tower Dashboard
![image](https://github.com/user-attachments/assets/840ef26e-02fa-42c4-a85d-1c6f71b1438e)

Click Resources --> Credentials
![image](https://github.com/user-attachments/assets/443e9275-0979-4d2d-b1ac-6a70fd1320e6)
Click "Add"
![image](https://github.com/user-attachments/assets/c0177e6a-705a-4e38-be1e-f8a6c09c440b)
![image](https://github.com/user-attachments/assets/079b2d1e-c5b3-45e7-98d2-61ded1632742)
Select "Machine" Credential Type
![image](https://github.com/user-attachments/assets/9c098188-6564-4268-b4fc-bea03cd36288)
Open your RPS Cloud Machine Terminal, type "cat /home/rps/.ssh/id_ed25519"
![image](https://github.com/user-attachments/assets/748b606d-230c-41bc-b114-ebed912d3635)
Copy the private key including the Begin and End as shown below
![image](https://github.com/user-attachments/assets/d91fee7a-abf6-4225-adf5-4d60b114de25)
Paste the private key you copied under the "SSH Private Key" field (Remove extra space)
![image](https://github.com/user-attachments/assets/161f70d2-bb49-432e-9458-15dec389187f)
Scroll down to save
![image](https://github.com/user-attachments/assets/41afb417-6511-488c-bdae-ec409898d78b)
Click Save
![image](https://github.com/user-attachments/assets/6f2fc74b-c7df-41e2-ab9e-234dcbf82b3e)



