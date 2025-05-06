# Day 2

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
