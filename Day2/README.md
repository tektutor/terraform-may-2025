# Day 2

## Info - SOLID Design Principles
<pre>
S - Single Responsibility Principle (SRP)
O - Open Closed Principle (OCP)
    - Your design is good if it satisfies the following desing principles
    - A design/code is considered good if it is open of extension
    - Extending new functionality should be done without modifying existing code
    - To add new functionality you can write new code 
L - Liskov Substitution Principle (LSP)
I - Interface Seggregation
D - Dependency Injection or Dependency Inversion or Inversion of Control(IOC)
</pre>

## Lab - Using Refactored static inventory file
Make sure your containers are running
```
docker ps
```

In case they are not running, you need to start them
```
docker start ubuntu1 ubuntu1
```

Now proceed as shown below
```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/before-refactoring
cat hosts
ansible -i hosts all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/31c9d09c-6c55-4b32-8246-cc16ba91ac4a)

#### Note
<pre>
In the above output
- ansible_port is a host variable as it defined for every host and the value is different for every host
- ansible_user, ansible_host, ansible_private_key_file are called as group variables as they are common for all the machines under the "all" group defined in the inventory file
</pre>

## Lab - Building a custom rocky linux ansible node image
```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/CustomDockerImage/rocky
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/rocky-ansible-node:latest .
rm authorized_keys
docker images 
```

Expected output
![image](https://github.com/user-attachments/assets/05dbb747-4869-47cf-a360-a91accec0bed)
![image](https://github.com/user-attachments/assets/3643b48a-8ea6-4aa3-86d0-25576b9471ee)

## Lab - Let's create couple of rocky containers using our custom rocky ansible node docker image
```
docker images | grep tektutor
docker ps

docker run -d --name rocky1 --hostname rocky1 -p 2003:22 -p 8003:80 tektutor/rocky-ansible-node:latest
docker run -d --name rocky2 --hostname rocky2 -p 2004:22 -p 8004:80 tektutor/rocky-ansible-node:latest 

docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/7bcab2e7-0702-47b9-a2dc-716fc22f03b6)

## Lab - Checking if we are able to SSH into the rocky1 and rock2 ansible node containers
```
docker ps

ssh -p 2003 root@localhost
exit

ssh -p 2004 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/9b949f6c-3277-4c80-ac5a-198b59043769)

## Lab - Integrating the rocky containers in the static inventory and check if ansible can ping rocky nodes
```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/before-refactoring
cat hosts
ansible all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/d605d56a-8248-4425-a2c3-290e36de7d18)
![image](https://github.com/user-attachments/assets/47d1c797-c04a-4bb7-8bb9-672cc61c9feb)
![image](https://github.com/user-attachments/assets/c377f7cd-dbbf-4a31-b2ea-96161f91bca2)
![image](https://github.com/user-attachments/assets/c14e518d-9c9c-494b-84eb-3e4f9078e14b)

## Lab - Adding support for Rocky linux in our install nginx playbook
```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/before-refactoring
cat hosts
cat install-nginx-playbook.yml
ansible-playbook install-nginx-playbook.yml

curl http://localhost:8001
curl http://localhost:8002
curl http://localhost:8003
curl http://localhost:8004
```

Expected output
![image](https://github.com/user-attachments/assets/ba746c30-bafd-4bd8-8bb8-3f6dd5f70cbd)
![image](https://github.com/user-attachments/assets/9bd83ea6-5473-45ab-829d-d7a6d3c2993e)
![image](https://github.com/user-attachments/assets/abe31e2b-dda6-432d-9e6a-4d5a02ee8fda)
![image](https://github.com/user-attachments/assets/bdca46ad-1106-4e69-974d-b0fa1246bd80)
![image](https://github.com/user-attachments/assets/1f2d556a-f9fb-4617-939a-5f9c2ff5d8b2)

## Lab - Running the refactored install nginx playbook
```
cd terraform-may-2025
git pull
cd Day2/ansible/after-refactoring
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/ec11fae1-3d38-42bb-9d16-ee2c836ccced)
![image](https://github.com/user-attachments/assets/112da6fd-3ba5-409d-9f68-38107ecd3e12)
![image](https://github.com/user-attachments/assets/170aff77-3fb7-437a-bafb-9130758796b7)
![image](https://github.com/user-attachments/assets/31a7b7bb-93df-4e00-bb6e-318c62394fd0)
![image](https://github.com/user-attachments/assets/e54ec909-792a-4df3-aaca-96a7ad7d6202)
![image](https://github.com/user-attachments/assets/8d38eecc-1f08-467a-b5b2-6f29a76beefd)
![image](https://github.com/user-attachments/assets/87a7df0e-dba9-425c-a73c-3daae5e52e91)
![image](https://github.com/user-attachments/assets/47eac379-5599-4b3b-bf0a-f53df9d1bd06)
![image](https://github.com/user-attachments/assets/80e51fbd-c237-4701-996a-e01bd303c915)

## Info - Ansible Role
<pre>
- is a reusable code
- ansible role though it looks like a playbook it can't be executed directly like we run a playbook
- ansible roles can be invoked from one or more playbooks
- if multiple softwares has to be installed, for each software we need to write a separate role
  for example
  - to install Oracle DB Server in Ubuntu, Fedora, RHEL we can develop a single oracle ansible role
  - to install weblogic app server in Ubuntu, RHEL, Suse we can develop a single weblogic ansible role
  - these roles can then be invoked from a playbook
- roles follows all the ansible recommended best practices, directory structure etc, to help us reuse code
</pre>

## Lab - Creating a custom nginx ansible role using ansible galaxy tool

As the nginx role is already created, you can skip the "ansible-galaxy create nginx" command.

```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/nginx-role
ansible-galaxy create nginx
tree nginx
cat install-nginx-playbook
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/a0984bd4-e10f-43d4-a3cc-282fbb75f10a)
![image](https://github.com/user-attachments/assets/04b06758-e3dd-48bc-94d1-82c7df043cb5)
![image](https://github.com/user-attachments/assets/fac4bb7d-cdae-4df3-b379-1e0d8b58f08f)
![image](https://github.com/user-attachments/assets/4e8e9041-f108-4589-9761-d7cd64817838)
![image](https://github.com/user-attachments/assets/cba20082-722c-4dc6-9ade-7198faa58d27)
![image](https://github.com/user-attachments/assets/c8797179-ae83-445c-94be-658995350e9d)

## Info - Ansible vault
<pre>
- any sensitive information like login credentials, public key, private key, certs, etc can be secured saved in ansible vault file
- ansible vault file is encrypted in AES256 encryption algorithm
- to encrypt/decrypt vault file, we need to provide a password
- the password can be typed manually or we can store it somewhere and give the path in the ansible.cfg
</pre>

## Lab - Using ansible vault tool to secure your sensitive information
```
ansible-vault create mysql-login-credentials.yml
ansible-vault view mysql-login-credentials.yml
ansible-vault edit mysql-login-credentials.yml
ansible-vault decrypt mysql-login-credentials.yml
ansible-vault encrypt mysql-login-credentials.yml
```

Expected output
![image](https://github.com/user-attachments/assets/6f2c9415-31eb-4df4-881c-5ae2ab31e126)
![image](https://github.com/user-attachments/assets/fb7bc88c-f780-4381-b5fa-9c3ef42cf9ff)

## Lab - Accessing vault protected data from an ansible playbook
```
cd ~/terraform-may-2025
git pull
cd Day2/ansible/vault
cat mysql-login-credentials.yml
cat playbook.yml
ansible-playbook playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/744e7af1-d7d4-43f8-b6d9-657b51e6b745)
![image](https://github.com/user-attachments/assets/c617995a-ab4b-4424-a932-87bbd417a36e)

## Info - Golang overview
<pre>
- golang is developed by Google using C programming language
- golang is a compiled programming language
- golang syntax resembles very close to C programming language
- just like C/C++/C# main function is the entry-point function ( the very function that will invoked when you run a go application )
- golang supports pointers but memory is managed by garbage collector unlike C/C++
- golang doesn't support class
- golang only supports functions
- using golang one can develop a new compiler/interpreter, a game, console based application, graphical application that runs on your local machine, can develop mobile applications, can develop AI/ML applications, web applications, etc.,
- using golang one can develop REST API, Microservices, etc.,
- is case-sensitive
- statically typed programming language
- performance wise, it is faster than most compiled languages, definitely more faster than interpretted and scripting languages
- even compilation is done faster in go lang for bulky applications
- Some of the popular tools developed in Golang
  - Terraform
  - Docker
  - Kubernetes
  - Openshift
</pre>

## Lab - Install Microsoft visual studio code editor in Ubuntu
```
sudo snap install code --classic
```

## Lab - Write your first Hello world in golang
Create a file named hello.go with the below content
```
package main

import "fmt"

func main() {
  fmt.Println( "Hello World !" )
}
```

Run your application
```
go run ./hello.go
```

Expected output
![image](https://github.com/user-attachments/assets/f6c31537-03f1-41f5-b543-d3a2c6e9b128)

## Lab - golang program that demonstrates accepting user inputs

Create a file name userinputs.go with the below content
```
package main

import "fmt"

func main() {
  //The below line declares a variable named x and it is initializing it with a value 0
  x := 0
  y := 0

  fmt.Print("Enter your first integer :" )
  fmt.Scanf("%d", &x )

  fmt.Print("Enter your second integer :" )
  fmt.Scanf("%d", &y )

  fmt.Println("Value of x :", x )
  fmt.Println("Value of y :", y )

  //Declares a variable named tmp of type string 
  var temp string
  fmt.Println("Press any key to exit ...")
  fmt.Scanln(&temp)
}
```

Run your application
```
go run ./userinputs.go
```

Expected output
![image](https://github.com/user-attachments/assets/12454c5e-9d6b-4923-9262-5a661071acf2)

## Lab - Understanding golang if else statement
Create a file named if-else.go with the below code
```
package main

import "fmt"

func main() {
   x := 100

   if x%2 == 0 {
	   fmt.Println(x, " is an even number")
   } else {
	   fmt.Println(x, " is an odd number")
   }
}
```

Run the program
```
go run ./if-else.go
```

Expected output
![image](https://github.com/user-attachments/assets/4eae6057-679f-4235-a7f4-d61dfd7d0f41)
